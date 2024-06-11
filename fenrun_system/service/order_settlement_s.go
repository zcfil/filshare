package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"xAdmin/common"
	orm "xAdmin/database"
	"xAdmin/define"
	log "xAdmin/logrus"
	"xAdmin/models"
	"xAdmin/utils"

	"github.com/jinzhu/gorm"
)

func (this *OrderSettlement) getratio() (float64, float64, error) {
	var GetOrder models.ConfigSet
	var FilFox1 float64

	damaged, err1 := GetOrder.Damaged()
	if err1 != nil {
		return 0, 0, nil
	}

	ret, err := GetOrder.Filfox()
	if err != nil && gorm.ErrRecordNotFound != err {
		return 0, 0, nil
	}

	switch ret[0].Locked {
	case 0:
		FilFox := ret[0].Value
		FilFox1 = utils.StringToFloat64(FilFox)
	case 1:
		data1, err2 := this.getAverageFil()
		if err2 != nil && gorm.ErrRecordNotFound != err2 {
			return 0, 0, nil
		}
		FilFox1 = data1.AverageFil

	}

	Damaged := utils.StringToFloat64(damaged) // 折损 0.98%
	return FilFox1, Damaged, nil
}

// 开始订单结算

func (this *OrderSettlement) settlement() (err error) {
	// 检查是否结果今天

	// 先取出结算折损 	 获取客户收益浮动比例
	filFox, damaged, err := this.getratio()
	if err != nil {
		return
	}

	sql := `insert into settle_date  (average_fil, date) values ('%s','%s')`
	sql = fmt.Sprintf(sql, strconv.FormatFloat(filFox, 'f', 10, 64), utils.TimeHMS())
	var idk []int64
	if err = orm.Eloquent.Exec(sql).Raw("select LAST_INSERT_ID() as id").Pluck("id", &idk).Error; err != nil {
		log.Error("插入结算记录数据错误， err:", err.Error())
		return err
	}

	lastID := idk[0]
	settleTime := time.Now()

	settleOrderList, err := this.RetrieveValidOrders() // 取出所有要结算的用户订单
	if err != nil {
		return
	}

	for _, msg := range settleOrderList { // 分别计算客户	FIL nanoFIL 10^9 attoFIL 10^18
		if err = this.settleOrder(&msg, filFox, damaged, settleTime, lastID); err != nil {
			return
		}
	}

	// 迁移的需求暂时不需要
	//MigrateData, err := this.GetMigrates(utils.TimeHMSStr(settleTime.Unix()))
	//if err != nil {
	//	log.Error("查询不到订单  : ", err.Error())
	//	return
	//}
	//
	//for _, msg := range MigrateData {
	//	if err = this.Migrates(&msg, settleTime); err != nil {
	//		log.Error("结算失败")
	//		return
	//	}
	//}

	// 到期订单状态更新
	data, err := this.getOderId(utils.TimeHMSStr(settleTime.Unix()))

	if err != nil && gorm.ErrRecordNotFound != err {
		log.Error(err.Error())
		return err
	}

	if len(data) != 0 {
		for _, msg := range data {
			if err = this.updateOderId(&msg); err != nil {
				return
			}
		}
	}
	return
}

func (this *OrderSettlement) getAverageFil() (*AverageFilConfig, error) {
	find := &AverageFilConfig{}
	sql := `select * from filearnings where id=(select max(id) from filearnings)`

	if err := orm.Eloquent.Raw(sql).Scan(find).Error; err != nil {
		log.Error("查找分润比例错误, err:", err.Error())
		return nil, err
	}
	return find, nil
}

func (this *OrderSettlement) settleOrder(order *OrderValid, filfox float64, damaged float64, times time.Time, lastID int64) (err error) {

	session := orm.Eloquent.Begin() // 保存到数据库  开始数据sql事务操作
	defer func() {
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
	}()

	StorageTotal := order.Storage                             // 购买存储单量
	TotalBalanceK := StorageTotal * filfox                    // 此单总收益  // 18.6
	TotalBalance := TotalBalanceK * damaged                   // 用户总收益  // 18.228
	toCustomerBalance := TotalBalance * define.INSTEAN_REWARD // 即时存储奖励25% 隔日结算发放  第1天 //4.557
	toCustomerLock := TotalBalance * define.LINEAR_REWARD     // 线性存储奖励75%   // 13.671
	CompanyTotal := TotalBalanceK - TotalBalance              // 公司收益 //0.372

	// 获取旧的用户锁定金额信息
	oldLockBalance, lockRelease, record, isExist, err1 := this.getLockBalanceMsg(order.Id, session, times)
	if err1 != nil {
		err = err1
		log.Error("获取锁仓数据错误， err:", err.Error())
		return
	}
	////23年3/6/
	//type Release struct {
	//	// 代码
	//	ID    int64  `gorm:"column:id" json:"id" `
	//	Name  string `gorm:"column:name" json:"name" `
	//	Days string `gorm:"column:days" json:"days" `
	//	IsDel int    `gorm:"column:is_del" json:"is_del"` //是否删除
	//}
	//sql := `select * from investment WHERE is_del=0`
	//var user Release
	//orm.Eloquent.Raw(sql).Scan(&user)

	var addLock float64
	day := len(record.Arr)
	//原来的判断
	//if day >= order.Day {
	//	addLock = lockRelease / float64(day) // 锁仓每天收益
	//	// 先减掉到期的
	//	record.Arr = record.Arr[1:]
	//} else if day <= order.Day {
	//	addLock = lockRelease / float64(order.Day)
	//}

	//新修改 23/3/6
	//dd, _ := utils.StringToInt64(order.Days)
	if day >= order.Days {
		addLock = lockRelease / float64(day) // 锁仓每天收益
		// 先减掉到期的
		record.Arr = record.Arr[1:]
	} else if day <= order.Days {
		addLock = lockRelease / float64(order.Days)
	}

	balanceAdd := addLock + toCustomerBalance // 用户余额

	newLockBalance := oldLockBalance + toCustomerLock - addLock // 锁仓记录 减去锁仓释放n分之一

	// 添加新的锁定部分
	toCustomerLock, _ = strconv.ParseFloat(fmt.Sprintf("%.6f", toCustomerLock), 64) // 保留6位小数
	this.addRecordNewLockBalance(record, toCustomerLock, times)
	if err = this.saveNewLockBalanceMsg(order.Id, session, newLockBalance, record, isExist); err != nil {
		log.Error("保存新的锁仓数据错误, err：", err.Error())
		return
	}

	log.Info("last id:", lastID)

	_, WalletBalance1, err := this.getBabalce(order.CustomerID, session) // 获取用户钱包余额
	if err != nil {
		log.Error("查询用户钱包异常 : ", err.Error())
		return
	}

	wallet := utils.StringToFloat64(WalletBalance1) // 用户余额

	TotalWallet := wallet + balanceAdd // 即时奖励奖励

	wallets := utils.Float64ToString(TotalWallet) // 及时奖励

	if err = this.balanceUpdate(order.CustomerID, utils.Float64ToString(newLockBalance), wallets, session, times); err != nil {
		log.Error("用户余额 跟更新 失败    ： ", err.Error())
		return
	}
	// 插入新的结算数据
	// 订单结算 日志
	if err = this.settleLog(session, order.Id, order.CustomerID, TotalBalanceK, TotalBalance, CompanyTotal, toCustomerBalance, addLock, newLockBalance, filfox, lastID, utils.TimeHMSStr(times.Unix())); err != nil {
		log.Error("记录失败     ： ", err.Error())
		return
	}

	return
}

func (this *OrderSettlement) settleLog(session *gorm.DB, orderIdk string, userid string, TotalBalanceK, TotalBalanceUser, CompanyTotal, Available, addLock, Locked, filfox float64, lastID int64, times string) (err error) {
	sqlx := `INSERT INTO settle_log
         (investment_id,customer_id,total_income,customer_income,company_income,to_customer_balance, 
           to_customer_lock,customer_lock_release,settle_date_id,time, is_transfer ,types_of,is_kid,active,enabled,income,type_kid)
      VALUES
        ( "%s","%s","%s","%v","%s","%s","%s", "%s","%v", "%s","%v",%d ,%v,%v,%v,"%s",%v)`

	TotalBalance := strconv.FormatFloat(TotalBalanceK, 'f', 18, 64)   // 此订单总收益
	UserBalance := strconv.FormatFloat(TotalBalanceUser, 'f', 18, 64) // 客户总收益
	TotalCompany := strconv.FormatFloat(CompanyTotal, 'f', -1, 64)    // 公司总收益
	UserAvailable := strconv.FormatFloat(Available, 'f', 10, 64)
	UserLocked := strconv.FormatFloat(Locked, 'f', 10, 64)
	AddLock := strconv.FormatFloat(addLock, 'f', 10, 64)
	FilfoxStr := utils.Float64ToString(filfox)
	sqlx = fmt.Sprintf(sqlx, orderIdk, userid, TotalBalance, UserBalance, TotalCompany, UserAvailable, UserLocked, AddLock, lastID, times, 0, 2, 2, 0, 0, FilfoxStr, 1)
	if err = session.Debug().Exec(sqlx).Error; err != nil {
		log.Error("插入失败 settle_log   : ", err.Error())
		return
	}
	return
}

func (this *OrderSettlement) getBabalce(userid string, session *gorm.DB) (string, string, error) { // 获取用户余额  锁仓余额

	var ret []CustomerBalance
	sqlx := `SELECT  locked_balance,wallet_balance  FROM customer where id = '%s'`
	fmt.Println(userid)
	sqlx = fmt.Sprintf(sqlx, userid)
	if err := session.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		log.Error("查询钱包错误 : ", err.Error())
		return "0", "0", err
	}
	if len(ret) > 0 {
		locked := ret[0].LockedBalance
		wallet := ret[0].WalletBalance
		return locked, wallet, nil
	}

	return "0", "0", nil

}

func (this *OrderSettlement) balanceUpdate(userid string, Locked string, Available string, session *gorm.DB, times time.Time) (err error) {
	// 更新用户钱包 余额 锁仓余额
	sqlx := `UPDATE  customer SET locked_balance="%s",wallet_balance="%s" , update_time="%s" WHERE id="%s"`
	sqlx = fmt.Sprintf(sqlx, Locked, Available, utils.TimeHMSStr(times.Unix()), userid)
	if err = session.Debug().Exec(sqlx).Error; err != nil {
		log.Error("更新用余额 锁仓余额  失败 : ", err.Error())
		return
	}
	log.Info("用户余额更新成功  successfully")
	return
}

func (this *OrderSettlement) RetrieveValidOrders() (ret []OrderValid, err error) { // 获取订单记录 为接下来开支准备进行
	sqlx := `SELECT  id AS order_id,storage ,total_day ,days,customer_id from investment where date_add(start_time, interval 1 day) <  now()  AND is_del<>1 AND status<>1 AND  end_time > now() ` //新增获取锁仓/释放周期
	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return ret, nil
}

func (this *OrderSettlement) GetOrderUserId(orderId string, session *gorm.DB) (userid string, err error) { // 通过订单号查询订单归属于那个用户
	type UserId struct {
		CustomerId string `gorm:"column:customer_id" json:"customer_id"`
	}
	sqlx := `SELECT  customer_id  FROM investment WHERE id = "%s"`
	sqlx = fmt.Sprintf(sqlx, orderId)
	ret := make([]UserId, 0)

	if err = session.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	if len(ret) > 0 {
		userid = ret[0].CustomerId
		return userid, err
	}

	fmt.Println(userid)
	return "", nil
}

func (this *OrderSettlement) getOderId(times string) (ret []OrderId, err error) {
	sqlx := `select  id,end_time,customer_id  from  investment where  datediff( '%s',end_time) = 0 `
	sqlx = fmt.Sprintf(sqlx, times)
	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}

func (this *OrderSettlement) updateOderId(msg *OrderId) (err error) {
	sqlx := `UPDATE  investment SET  status=2 WHERE  is_del=0 AND id='` + msg.Id + `'`
	if err = orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return
	}
	return
}

// getLockBalanceMsg 取出分批释放的数据
func (this *OrderSettlement) getLockBalanceMsg(customerID string, session *gorm.DB, settleTime time.Time) (lockBalance float64, lockRelease float64, record *LockBalanceRecord, isExist bool, err error) {
	isExist = true
	type findLockBalance struct {
		LockBalance float64 `gorm:"column:lock_balance"`
		Record      string  `gorm:"column:record"`
	}
	flb := &findLockBalance{}
	record = &LockBalanceRecord{}
	sql := `select * from lock_balance where customer_id = '` + customerID + `'`
	err = session.Raw(sql).Scan(flb).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		isExist = false
		record.Arr = make([]float64, 0)
		record.UpdateDay = common.TimeToDay(time.Now())
		return
	}
	if err != nil {
		log.Error("查询数据错误:", err.Error(), "顾客ID:", customerID)
		return
	}

	lockBalance = flb.LockBalance
	lockRelease = lockBalance
	if len(flb.Record) <= 0 {
		return
	}
	err = json.Unmarshal([]byte(flb.Record), record)
	if err != nil {
		return
	}
	// 防止一天多笔投资的情况出出现，上一单的锁定记录会成为下一单的释放
	if common.TimeToDay(settleTime) == record.UpdateDay {
		lockRelease -= record.Arr[len(record.Arr)-1]
	}
	return
}

func (this *OrderSettlement) addRecordNewLockBalance(record *LockBalanceRecord, toCustomerLock float64, settleTime time.Time) {
	day := common.TimeToDay(settleTime)
	if len(record.Arr) <= 0 {
		record.UpdateDay = day
		record.Arr = append(record.Arr, toCustomerLock)
		return
	}

	// 如果当天结算把锁仓金额累加到一起
	if record.UpdateDay == day {
		record.Arr[len(record.Arr)-1] += toCustomerLock
		record.UpdateDay = day
		return
	}

	record.UpdateDay = day
	record.Arr = append(record.Arr, toCustomerLock)
}

func (this *OrderSettlement) saveNewLockBalanceMsg(customerID string, session *gorm.DB, lockBalance float64, record *LockBalanceRecord, isExist bool) (err error) {
	bin, err1 := json.Marshal(record)
	if err1 != nil {
		err = err1
		return
	}

	recordStr := string(bin)
	lockBalanceStr := strconv.FormatFloat(lockBalance, 'f', 10, 64)
	if isExist {
		sql := `update lock_balance set lock_balance = '%s', record = '%s' where customer_id = %s`
		sql = fmt.Sprintf(sql, lockBalanceStr, recordStr, customerID)
		err = session.Exec(sql).Error
		if err != nil {
			return
		}
		return
	}

	sql := `insert into lock_balance (customer_id, lock_balance, record) values ('%s', '%s', '%s')`
	sql = fmt.Sprintf(sql, customerID, lockBalanceStr, recordStr)
	err = session.Exec(sql).Error
	if err != nil {
		return
	}
	return
}
