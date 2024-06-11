package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/utils"
)

func (this *OrderSettlement) GetMigrates(times string) (ret []Migrate, err error) {
	sqlx := `select *  from migrate_order  Where   active_status<>1  AND  start_time <= '` + times + `'  AND   end_time >'` + times + `' AND is_del<>1 AND status<>1`
	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		log.Error("err Error ", err.Error())
		return
	}
	return
}

func (this *OrderSettlement) Migrates(orderID *Migrate, times time.Time) (err error) {
	var day int
	session := orm.Eloquent.Begin() // 保存到数据库  开始数据sql事务操作
	defer func() {
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
	}()

	LockedBalance1, WalletBalance1, err := this.getBabalce(orderID.UserId, session) //钱包 余额  锁仓
	if err != nil {
		log.Error("查询用户钱包异常 : ", err.Error())
		return
	}

	orderKid := orderID.OrderMid                                     // 订单内存
	userid := orderID.UserId                                         // 用户id
	wallet := utils.StringToFloat64(WalletBalance1)                  // 用户余额
	TotalWallet := wallet + utils.StringToFloat64(orderID.OneReward) // 即时奖励奖励
	wallets := utils.Float64ToString(TotalWallet)                    // 余额
	if orderID.TotalDays == 0 {
		sqlx1 := `update  migrate_order  set   active_status=1, status=1 ,time=now()   where order_mid='` + orderID.OrderMid + `'`
		sqlx1 = fmt.Sprintf(sqlx1, day)
		if err = session.Debug().Exec(sqlx1).Error; err != nil {
			log.Error("插入失败 settle_log   : ", err.Error())
			return
		}
	} else {
		day = orderID.TotalDays - 1
	}

	sqlx1 := `update  migrate_order  set   total_days='%d',time=now()  where order_mid='` + orderID.OrderMid + `'`
	sqlx1 = fmt.Sprintf(sqlx1, day)
	if err = session.Debug().Exec(sqlx1).Error; err != nil {
		log.Error("插入失败 settle_log   : ", err.Error())
		return
	}

	if err = this.balanceUpdate(userid, LockedBalance1, wallets, session, times); err != nil {
		log.Error("用户余额 跟更新 失败    ： ", err.Error())
		return
	}

	if err = this.SettlementLog1(session, orderKid, orderID.OneReward, userid, utils.TimeHMSStr(times.Unix())); err != nil {
		return
	}

	return
}

func (this *OrderSettlement) SettlementLog1(session *gorm.DB, orderKid string, OneIncome string, userid string, times string) (err error) {
	log.Info(OneIncome)
	sqlx := `INSERT INTO settle_log
         (investment_id,customer_id,total_income,customer_income,company_income,to_customer_balance, 
           to_customer_lock,customer_lock_release,settle_date_id,time, is_transfer ,types_of,update_time,active,enabled,type_kid)
      VALUES
        ( "%s","%s",0,0,0,0,0,"%s", 0,"%s",0,7,"%s",0,0,2)`
	sqlx = fmt.Sprintf(sqlx, orderKid, userid, OneIncome, times, times)
	log.Info("sqlx", sqlx)
	if err = session.Debug().Exec(sqlx).Error; err != nil {
		log.Error("插入失败 settle_log   : ", err.Error())
		return
	}
	return
}
