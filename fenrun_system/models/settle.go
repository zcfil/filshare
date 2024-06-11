package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"xAdmin/common"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/utils"
)

type Settle struct {
	fromWallet string
}

func NewSettle() *Settle {
	s := new(Settle)

	return s
}

// 未结算按周显示

func (this *Settle) GetWeekList(param map[string]string) (ret interface{}, err error) {

	sql := ` select DATE_FORMAT(time, '%u') weeks,  sum(to_customer_balance + customer_lock_release)     amount,
       CONCAT(DATE_FORMAT(subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 1),  '%Y-%m-%d'), '--',
         DATE_FORMAT(subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 7), '%Y-%m-%d'))                      date,
       (select name from tables_name where kid = active) as status from (select s.* from settle_log s  left join customer c on c.id = s.customer_id) a
             where active <> 1 AND enabled <> 1 GROUP BY weeks order by date asc `
	param["total"] = GetTotalCount(sql)

	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	sql += ` limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]

	type find struct {
		Week   int32   `gorm:"column:week" json:"week"`
		Amount float64 `gorm:"column:amount" json:"amount"`
		Date   string  `gorm:"column:date" json:"date"`
	}

	finds := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}
	ret = finds
	return
}

// 已经结算周

func (this *Settle) GetWeekListSettled(param map[string]string) (ret interface{}, err error) {
	sqlx := `select DATE_FORMAT(time, '%u') weeks,
		sum(to_customer_balance + customer_lock_release) amount,  CONCAT( DATE_FORMAT(
		subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 1),  '%Y-%m-%d'), '--',
		DATE_FORMAT( subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 7),  '%Y-%m-%d'))  date
		from (select s.* from settle_log s  inner join customer c on c.id = s.customer_id) a
		where is_transfer = 1    AND weekes = 1  GROUP BY weeks order by  date asc `

	param["total"] = GetTotalCount(sqlx)

	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	sqlx += ` limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]

	type find struct {
		Week   int32   `gorm:"column:week" json:"week"`
		Amount float64 `gorm:"column:amount" json:"amount"`
		Date   string  `gorm:"column:date" json:"date"`
	}

	finds := make([]find, 0)
	if err = orm.Eloquent.Raw(sqlx).Scan(&finds).Error; err != nil {
		return
	}
	ret = finds
	return
}

func (this *Settle) GetWeekCustomerList(param map[string]string) (ret interface{}, err error) {
	startDate, endDate, err := this.getStartEndTime(param["date"])
	if err != nil {
		return
	}
	sql := `select c.name, c.phone, a.* from (select s.customer_id, sum(s.to_customer_balance) as to_balance, sum(s.customer_lock_release) as lock_release,
						sum(s.to_customer_balance+s.customer_lock_release) as amount
						from settle_log s where UNIX_TIMESTAMP(time) >= UNIX_TIMESTAMP('%s') and
						UNIX_TIMESTAMP(time) <= UNIX_TIMESTAMP('%s') and
						is_transfer<>1 group by s.customer_id ) a
						inner join customer c on a.customer_id=c.id `

	sql = fmt.Sprintf(sql, startDate, endDate)

	param["total"] = GetTotalCount(sql)
	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	sql += ` limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]

	type find struct {
		CustomerID  string  `gorm:"column:customer_id" json:"customer_id"`
		Amount      float64 `gorm:"column:amount" json:"amount"`             // 总收益
		LockRelease float64 `gorm:"column:lock_release" json:"lock_release"` // 线性释放收益
		ToBalance   float64 `gorm:"column:to_balance" json:"to_balance"`     // 直接释放收益
		Name        string  `gorm:"column:name" json:"name"`
		Phone       float64 `gorm:"column:phone" json:"phone"`
	}

	findList := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	ret = findList
	return
}

func (this *Settle) GetWeekCustomerInvestmentList(param map[string]string) (ret interface{}, err error) {
	startDate, endDate, err := this.getStartEndTime(param["date"])
	if err != nil {
		return
	}
	sql := `select s.id, (s.to_customer_balance + s.customer_lock_release)  as amount,  s.time   settle_time,  c.name,  s.investment_id,  c.phone,  s.time                                              AS create_time,
      (select name from tables_name where kid = types_of) as types_of from settle_log s
        left join customer c on s.customer_id = c.id  left join investment i on s.investment_id = i.id
       where UNIX_TIMESTAMP(s.time) >= UNIX_TIMESTAMP('%s')  and UNIX_TIMESTAMP(s.time) <= UNIX_TIMESTAMP('%s') and s.customer_id=%s`
	sql = fmt.Sprintf(sql, startDate, endDate, param["customerID"])

	param["total"] = GetTotalCount(sql)
	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	sql += ` limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]
	log.Info(sql)
	type find struct {
		ID            string  `gorm:"column:id" json:"id"`
		Amount        float64 `gorm:"column:amount" json:"amount"`
		Name          string  `gorm:"column:name" json:"name"`
		Investment_id string  `gorm:"column:investment_id" json:"investment_id"`
		Phone         string  `gorm:"column:phone" json:"phone"`
		SettleTime    string  `gorm:"column:settle_time" json:"settle_time"`
		CreateTime    string  `gorm:"column:create_time" json:"create_time"`
		Types_of      string  `gorm:"column:types_of" json:"types_of"`
	}

	findList := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	ret = findList
	return
}

func (this *Settle) TransferWeek(date string) (err error) {
	if !common.CheckRequest(date) {
		return errors.New("已发送过请求")
	}
	startDate, endDate, err := this.getStartEndTime(date)
	if err != nil {
		return
	}
	sql := `select c.wallet, a.*
       from (select sum(to_customer_balance + customer_lock_release) amount, customer_id   from settle_log  where is_transfer <> 1  AND enabled<>1
        and UNIX_TIMESTAMP(time) >= UNIX_TIMESTAMP('%s') and UNIX_TIMESTAMP(time) <= UNIX_TIMESTAMP('%s')
         group by (customer_id)) a  inner join customer c on a.customer_id = c.id `

	sql = fmt.Sprintf(sql, startDate, endDate)
	type find struct {
		Wallet     string  `gorm:"column:wallet" json:"wallet"`
		Amount     float64 `gorm:"column:amount" json:"amount"`
		CustomerID string  `gorm:"column:customer_id" json:"customer_id"`
	}

	findList := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	if len(findList) <= 0 {
		return
	}

	var finance Default_Finance
	if err = finance.GetFromWallet(); err != nil {
		return
	}
	dateTimes := utils.TimeHMS()
	this.fromWallet = finance.ConfigValue
	log.Info(finance.ConfigValue)
	for _, msg := range findList {
		if err = this.transferCustomer(msg.CustomerID, msg.Wallet, msg.Amount, startDate, endDate, dateTimes); err != nil {
			log.Error("用户ID：", msg.CustomerID, " 转账金额:", msg.Amount, " 转账钱包:", msg.Wallet, " err:", err.Error())
			continue
		}
		log.Info("已转账用户ID：", msg.CustomerID, " 转账金额:", msg.Amount, " 转账钱包:", msg.Wallet)
	}
	return
}

func (this *Settle) TransferWeekCustomer(param map[string]string) (err error) {
	requestStr, _ := json.Marshal(param)
	// 给单独给客户转账也要验证一下日期
	if !common.CheckRequest(string(requestStr)) || !common.CheckRequest(param["date"]) {
		return errors.New("已发送过请求")
	}
	startDate, endDate, err := this.getStartEndTime(param["date"])
	if err != nil {
		return
	}
	datetimes := utils.TimeHMS()
	var finance Default_Finance
	if err = finance.GetFromWallet(); err != nil {
		return
	}
	this.fromWallet = finance.ConfigValue

	customerID := param["customerID"]

	sql := `select sum(to_customer_balance+customer_lock_release) amount, c.wallet from settle_log s
		left join customer c on s.customer_id = c.id
		where UNIX_TIMESTAMP(s.time) >= UNIX_TIMESTAMP('%s') and
		UNIX_TIMESTAMP(s.time) <= UNIX_TIMESTAMP('%s') and s.is_transfer<>1 and s.customer_id=%s`
	sql = fmt.Sprintf(sql, startDate, endDate, customerID)
	type find struct {
		Wallet string  `gorm:"column:wallet" json:"wallet"`
		Amount float64 `gorm:"column:amount" json:"amount"`
	}
	f := &find{}
	if err = orm.Eloquent.Raw(sql).Scan(f).Error; err != nil {
		return
	}

	if f.Amount <= 0 {
		return
	}

	if err = this.transferCustomer(customerID, f.Wallet, f.Amount, startDate, endDate, datetimes); err != nil {
		err = errors.New("转账失败!!!")
		log.Error("转账失败， 转账信息:", customerID, f.Wallet, f.Amount, startDate, endDate)
		return
	}
	return
}

func (this *Settle) transferCustomer(customerID, wallet string, amount float64, startDate, endDate string, date string) (err error) {
	session := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
	}()
	//先设置状态再转账
	sql := `update settle_log set is_transfer=1 ,enabled=1 ,active=1, update_time='` + date + `' where UNIX_TIMESTAMP(time)>=UNIX_TIMESTAMP('` + startDate + `') and 
			UNIX_TIMESTAMP(time)<=UNIX_TIMESTAMP('` + endDate + `') and is_transfer<>1 and customer_id=%s
`
	sql = fmt.Sprintf(sql, customerID)
	if err = session.Exec(sql).Error; err != nil {
		return
	}
	var transfer Transfer
	if err = transfer.Send(context.Background(), customerID, this.fromWallet, wallet, amount); err != nil {
		return
	}
	return
}

func (this *Settle) getStartEndTime(date string) (start, end string, err error) {
	list := strings.Split(date, "--")
	if len(list) != 2 {
		err = errors.New("日期格式不正确")
		log.Warning("date:", date)
		return
	}

	start = list[0] + " 00:00:00"
	end = list[1] + " 23:59:59"
	return
}

func (this *Settle) SettleList(param map[string]string) (ret interface{}, err error) {
	type settleLog struct {
		InvestmentID   string  `gorm:"column:investment_id" json:"investment_id"`     // 投资ID
		CustomerID     string  `gorm:"column:customer_id" json:"customer_id"`         // 投资者ID
		CustomerName   string  `gorm:"column:customer_name" json:"customer_name"`     // 投资者名字
		TotalIncome    float64 `gorm:"column:total_income" json:"total_income"`       // 总收益
		CustomerIncome float64 `gorm:"column:customer_income" json:"customer_income"` // 用户收益
		//CompanyIncome       float64   `gorm:"column:company_income" json:"company_income"`               // 公司收益
		ToCustomerBalance   float64   `gorm:"column:to_customer_balance" json:"to_customer_balance"`     // 到用户余额
		ToCustomerLock      float64   `gorm:"column:to_customer_lock" json:"to_customer_lock"`           // 到用户锁仓
		CustomerLockRelease float64   `gorm:"column:customer_lock_release" json:"customer_lock_release"` // 用户锁仓释放
		SettleTime          time.Time `gorm:"column:time" json:"time"`                                   // 结算时间
		IsTransfer          int       `gorm:"column:is_transfer" json:"is_transfer"`                     // 是否已转账 1已转账， 0未转账
		Income              float64   `gorm:"column:income" json:"income"`                               // 单T收益
		//is_kid              string    `gorm:"column:is_kid" json:"is_kid"`
	}

	sql := `select c.name customer_name, s.*  from settle_log s
				left join customer c on s.customer_id=c.id   WHERE  types_of = 2  AND type_kid=1 order by time desc`

	param["total"] = GetTotalCount(sql)
	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	sql += ` limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]

	finds := make([]settleLog, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}
	ret = finds
	return
}
