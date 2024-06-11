package models

import (
	"fmt"
	"strconv"
	"time"
	orm "xAdmin/database"
)

func (this *Settle) GetWeekCustomerList1(param map[string]string) (ret interface{}, err error) {
	startDate, endDate, err := this.getStartEndTime(param["date"])
	if err != nil {
		return
	}

	sql := `select c.name, c.phone, a.* from (select s.customer_id,
           sum(s.to_customer_balance)    as to_balance,  sum(s.company_income)  AS company_income, sum(s.customer_lock_release)  as lock_release,  s.time   AS create_time,
           sum(s.to_customer_balance + s.customer_lock_release) as amount,  (select name from tables_name where kid = s.active) as status
           from settle_log s  where UNIX_TIMESTAMP(time) >= UNIX_TIMESTAMP('%s')  and UNIX_TIMESTAMP(s.time) <= UNIX_TIMESTAMP('%s')
          and is_transfer <> 1  AND enabled <>1   AND enabled <>1  group by s.customer_id) a  inner join customer c on a.customer_id = c.id`
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
		CustomerID     string  `gorm:"column:customer_id" json:"customer_id"`
		Company_income string  `gorm:"column:amount" json:"amount"`
		Amount         float64 `gorm:"column:company_income" json:"company_income"` // 总收益
		LockRelease    float64 `gorm:"column:lock_release" json:"lock_release"`     // 线性释放收益
		ToBalance      float64 `gorm:"column:to_balance" json:"to_balance"`         // 直接释放收益
		Name           string  `gorm:"column:name" json:"name"`
		Phone          float64 `gorm:"column:phone" json:"phone"`
		Create_time    string  `gorm:"column:create_time" json:"create_time"`
		Status         string  `gorm:"column:status" json:"status"`
	}

	findList := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	ret = findList
	return
}

func (this *Settle) GetWeekCustomerInvestmentList1(param map[string]string) (ret interface{}, err error) {
	startDate, endDate, err := this.getStartEndTime(param["date"])
	if err != nil {
		return
	}
	sql := `select s.id, (s.to_customer_balance+s.customer_lock_release) as amount, s.time settle_time, c.name, s.investment_id,c.phone, i.create_time ,       
          (select name from tables_name  where kid=types_of) as types_of  from settle_log s 
			left join customer c on s.customer_id=c.id
			left join investment i on s.investment_id = i.id 
			where UNIX_TIMESTAMP(s.time) >= UNIX_TIMESTAMP('%s') 
			and UNIX_TIMESTAMP(s.time) <= UNIX_TIMESTAMP('%s') and s.customer_id=%s and s.is_transfer <> 1`
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

	type find struct {
		ID            string    `gorm:"column:id" json:"id"`
		Amount        float64   `gorm:"column:amount" json:"amount"`
		Name          string    `gorm:"column:name" json:"name"`
		Investment_id string    `gorm:"column:investment_id" json:"investment_id"`
		Phone         string    `gorm:"column:phone" json:"phone"`
		SettleTime    time.Time `gorm:"column:settle_time" json:"settle_time"`
		CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
		Types_of      string    `gorm:"column:types_of" json:"types_of"`
	}

	findList := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	ret = findList
	return
}

func (this *Settle) SettleList1(param map[string]string) (ret interface{}, err error) {
	type settleLog struct {
		//InvestmentID        string    `gorm:"column:investment_id" json:"investment_id"`                 // 投资ID
		//CustomerID          string    `gorm:"column:customer_id" json:"customer_id"`                     // 投资者ID
		//CustomerName        string    `gorm:"column:customer_name" json:"customer_name"`                 // 投资者名字
		TotalIncome         float64 `gorm:"column:total_incomes" json:"total_incomes"`                   // 总收益
		CustomerIncome      float64 `gorm:"column:customer_incomes" json:"customer_incomes"`             // 用户收益
		CompanyIncome       float64 `gorm:"column:company_incomes" json:"company_incomes"`               // 公司收益
		ToCustomerBalance   float64 `gorm:"column:to_customer_balances" json:"to_customer_balances"`     // 到用户余额
		ToCustomerLock      float64 `gorm:"column:to_customer_locks" json:"to_customer_locks"`           // 到用户锁仓
		CustomerLockRelease float64 `gorm:"column:customer_lock_releases" json:"customer_lock_releases"` // 用户锁仓释放
		Date                string  `gorm:"column:date" json:"date"`                                     // 结算时间
		IsTransfer          string  `gorm:"column:is_transfer" json:"is_transfer"`                       // 是否已转账 1已转账， 0未转账
	}

	sql := `select  
       DATE_FORMAT(time, '%u')  weeks, sum(total_income)  total_incomes,   sum(customer_income) customer_incomes, sum(company_income) company_incomes,
       sum(to_customer_balance)   to_customer_balances, (sum(to_customer_lock) - sum(customer_lock_release)  ) to_customer_locks, sum(customer_lock_release)  customer_lock_releases,
       (select name from tables_name where kid = is_transfer) as is_transfer,
       CONCAT(DATE_FORMAT(  subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 1), '%Y-%m-%d'), '--',
              DATE_FORMAT(subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 7), '%Y-%m-%d'))                           date
        from settle_log s  left join customer c on s.customer_id = c.id GROUP BY weeks  desc`
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

// 导出周结算列表
func (this *Settle) ExportWeekCustomerList(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["title"] = "周结算报表" + param["date"]
	URL, err = GetWeekCustomerList(this.GetWeekCustomerList1, param)

	return
}
