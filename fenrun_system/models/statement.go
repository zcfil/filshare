package models

import (
	"errors"
	"strconv"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Statement struct {
	Percent      float64 `gorm:"column:percent" json:"percent"`
	Profits      float64 `gorm:"column:profits" json:"profits"`
	Amount       float64 `gorm:"column:amount" json:"amount"`
	Userid       string  `gorm:"column:userid" json:"userid"`
	Username     string  `gorm:"column:username" json:"username"`
	NickName     string  `gorm:"column:nick_name" json:"nick_name"`
	Name         string  `gorm:"column:customername" json:"customername"`
	Investmentid string  `gorm:"column:investmentid" json:"investmentid"`
}
type Total struct {
	Profits float64 `json:"profits"`
	Total   string  `json:"total"`
	Amount  float64 `json:"amount"`
}
type TotalList struct {
	Profits float64 `json:"profits"`
	Userid  string  `json:"userid"`
	Amount  float64 `json:"amount"`
}

//业务员报表
func (e *Statement) StatementSalesmanNew(param map[string]string) (result interface{}, total interface{}, err error) {
	//获取投资列表
	//sql := `select c.userid,u.nick_name,c.name customername,u.username,i.amount,ip.profits,ip.investmentid from investment i
	//	left join investmentprofit ip on i.id = ip.investmentid
	//	left join customer c on i.customerid = c.id
	//	left join sys_user u on ip.userid = u.user_id
	//	where i.status in (0,4,5) and i.is_del = 0 and i.create_time <= :end and
	//										:start <= i.create_time `
	//	//where i.status in (0,4,5) and i.is_del = 0 and i.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
	//	//									date_add(:date,interval -day(:date)+1 day) <= i.create_time `
	sql := `select ip.userid,u.nick_name,c.name customername,u.username,round(i.amount,2)amount,round(ip.profits,2)profits,ip.investmentid from investment i
		left join investmentprofit ip on i.id = ip.investmentid
		left join customer c on i.customerid = c.id
		left join sys_user u on ip.userid = u.user_id
		where i.status in (0,4,5,6) and ifnull(i.manually_end,0) = 0 and i.is_del = 0 and i.create_time <= :end and
											:start <= i.create_time `
	sql = utils.SqlReplaceParames(sql, param)
	var s []Statement
	var t Total
	if err = orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	//因为关联了子表，sum(amount)值不对
	var t1 Total
	sqlt1 := `select round(sum(i.amount),2)amount from investment i
		where i.status in (0,4,5,6) and i.is_del = 0 and i.create_time <= :end and
											:start <= i.create_time `
	sqlt1 = utils.SqlReplaceParames(sqlt1, param)
	if err = orm.Eloquent.Raw(sqlt1).Scan(&t1).Error; err != nil {
		return nil, nil, err
	}

	t.Amount = t1.Amount

	sqltotal := `select round(sum(i.amount),2) amount,round(sum(ip.profits),2) profits,ip.userid from investment i
		left join investmentprofit ip on i.id = ip.investmentid
		left join customer c on i.customerid = c.id
		left join sys_user u on ip.userid = u.user_id
		where i.status in (0,4,5,6) and i.is_del = 0 and i.create_time <= :end and
											:start <= i.create_time
		GROUP BY ip.userid`
	var tlist []TotalList
	sqltotal = utils.SqlReplaceParames(sqltotal, param)
	if err = orm.Eloquent.Raw(sqltotal).Scan(&tlist).Error; err != nil {
		return
	}
	param["total"] = t.Total
	param["sort"] = "ip.userid ,i.create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	if err = orm.Eloquent.Raw(sql).Scan(&s).Error; err != nil {
		return nil, nil, err
	}
	var pre Statement
	if len(s) > 0 {
		pre = s[0]
	}

	var st []Statement

	for i, v := range s {
		if v.Userid == pre.Userid {
			pre = v
			st = append(st, v)
			if i == len(s)-1 {
				for _, val := range tlist {
					if val.Userid == v.Userid {
						var total Statement
						total.NickName = "合计"
						total.Profits = val.Profits
						total.Amount = val.Amount
						st = append(st, total)
					}
				}
			}
			continue
		}
		for _, val := range tlist {
			if val.Userid == pre.Userid {
				//fmt.Println(val)
				var total Statement
				total.NickName = "合计"
				total.Profits = val.Profits
				total.Amount = val.Amount
				st = append(st, total)
			}
		}
		st = append(st, v)
		pre = v
	}

	return st, t, err
}
func (e *Statement) StatementSalesman(param map[string]string) (result interface{}, total interface{}, err error) {
	//sql := `select *,amount*percent profits from (
	//		select round(sum(percent),4)percent,amount,a.userid,u.username,u.nick_name,u.nick_name name from (
	//				select c.amount,ifnull(p.userid,0)pid,c.userid from (
	//						select sum(i1.amount)amount,c1.userid from customer c1
	//							left join investment i1 on i1.customerid = c1.id
	//							where c1.is_del = 0 and i1.status in (0,4,5) and i1.is_del = 0 and i1.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
	//										date_add(:date,interval -day(:date)+1 day) <= i1.create_time
	//							group by c1.userid
	//						)c
	//						left join (select userid from profitconfig WHERE profittype=1 GROUP BY userid) p on p.userid = c.userid
	//		)a
	//		left join profitconfig p1 on a.pid = p1.userid
	//		left join sys_user u on a.userid = u.user_id
	//		where p1.profitlevel < a.amount and p1.profittype = if(ifnull(p1.userid,0)=0,2,1)
	//		GROUP BY a.userid
	//		)b `
	sql := `select sum(amount) amount,sum(profits)profits,a.userid,u.username,u.nick_name,u.nick_name name from (
				select amount,a.userid,a.id,sum(p1.levelgain)+(amount-max(p1.profitlevel))*sum(percent) profits from (
					select c.amount,ifnull(p.userid,0)pid,c.userid,c.id from (
							select i1.amount,c1.userid,i1.id from customer c1
								left join investment i1 on i1.customerid = c1.id
								where c1.is_del = 0 and i1.status in (0,4,5) and i1.is_del = 0 and i1.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
											date_add(:date,interval -day(:date)+1 day) <= i1.create_time
							)c 
							left join (select userid from profitconfig WHERE profittype=1 GROUP BY userid) p on p.userid = c.userid

				)a 
				left join profitconfig p1 on a.pid = p1.userid
				where p1.profitlevel < a.amount and p1.profittype = if(ifnull(p1.userid,0)=0,2,1) 
				GROUP BY a.id
			)a
			left join sys_user u on a.userid = u.user_id
			GROUP BY a.userid `
	now := time.Now()
	nowdate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	if param["date"] != nowdate {
		sql = `select *,name nick_name from statement where DATE_FORMAT(create_time,'%Y-%m-01') =:date and stype = 1 `
	}
	sql = utils.SqlReplaceParames(sql, param)
	var s []Statement
	var t Total

	if err = orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	param["sort"] = "amount"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	err = orm.Eloquent.Raw(sql).Scan(&s).Error
	return s, t, err
}

type UserInvestment struct {
	UserLevel
	Investment
	Referrer
	//ShareProfit []InvestmentShareProfit
}

//顾客报表
func (e *Statement) StatementCustomerNew(param map[string]string) (result interface{}, total interface{}, err error) {

	sql := `select i.profit, 
			if(i.expiration_date>i.monthly_time,
						ROUND(i.profit*i.amount,2),
						ROUND(TIMESTAMPDIFF(day,date_sub(date(monthly_time) ,interval 1 month),expiration_date)/day(LAST_DAY(date_sub(monthly_time, interval 1 month)))*i.profit*i.amount,2)) profits,
			i.amount,c.id, c.name, c.phone, c.bank, c.banknum, c.userid, c.is_del, c.status, c.create_time, c.sex, c.identity, 
			u.nick_name, i.expiration_date,i.id investment_id, i.monthly_time
			from investment i
			LEFT JOIN customer c on i.customerid = c.id
			left join sys_user u on c.userid = u.user_id
			WHERE i.is_del = 0 and i.status in (0,4,5) and day(i.update_time) <= day(i.expiration_date) 
		`
	if param["day"] != "" {
		sql += ` and DATE_FORMAT(monthly_time,'%e') = :day `
	}
	var s []Customer
	sql = utils.SqlReplaceParames(sql, param)

	var t Total
	if err = orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	param["sort"] = "monthly_time"
	param["order"] = "asc"
	sql += utils.LimitAndOrderBy(param)
	err = orm.Eloquent.Raw(sql).Scan(&s).Error
	return s, t, err
}

//判断上月报表是否已生成
func (e *Statement) GetStatementBool(date string, stype int) (result bool) {
	sql := `select 1 from statement s where s.stype = ` + strconv.Itoa(stype) + ` and date_format(create_time,'%Y-%m') = date_format('` + date + `','%Y-%m')  `
	return GetTotalCount1(sql) > 0
}
func (e *Statement) GetSummaryBool(date string) (result bool) {
	sql := `select 1 from summary s where date_format(create_time,'%Y-%m') = date_format('` + date + `','%Y-%m')  `
	return GetTotalCount1(sql) > 0
}

//导出业务员报表
func (e *Statement) ExportExcelSalesman(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["filefield"] = "nick_name,investmentid,customername,amount,profits"
	param["filename"] = "业务员姓名,订单号,客户姓名,业绩,利润"
	param["title"] = "业务员报表" + param["date1"]
	URL, err = GetExcelTotal1(e.StatementSalesmanNew, param)
	return
}

//导出客户报表
func (e *Statement) ExportExcelCustomer(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["filefield"] = "name,phone,banknum,bank,sex,profit,amount,profits,nick_name,create_time,update_time,expiration_date"
	param["filename"] = "姓名,手机号,银行卡,开户行,性别,分润比例,总业绩,利润,业务员,投资时间,结算时间,合同到期时间"
	param["title"] = "客户报表" + param["date1"]
	URL, err = GetExcelTotal(e.StatementCustomerNew, param)
	return
}

func (e *Statement) CustomerSettle(investmentID string) (err error) {
	sql1 := `select c.name name, c.bank bank, c.banknum banknum, c.userid user_id, u.nick_name nick_name, c.id customer_id
		from investment i, customer c, sys_user u
		where i.id=` + investmentID + ` and c.id = i.customerid and c.userid = u.user_id`

	find := make([]SettleCustomerHistory, 0)
	if err = orm.Eloquent.Raw(sql1).Scan(&find).Error; err != nil {
		return
	}
	if len(find) != 1 {
		err = errors.New("查找数据错误")
		return
	}

	findMsg := find[0]
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()

	sql2 := `update investment SET monthly_time=if(expiration_date<monthly_time,monthly_time,adddate(monthly_time, interval 1 month)),status = if(status=5||monthly_time>expiration_date,6,status), update_time=now() WHERE id = ` + investmentID
	if err = orm1.Exec(sql2).Error; err != nil {
		return
	}

	sql3 := `insert into customer_settle (name, bank, banknum, user_id, nick_name, settle_time, invest_id) values ("` +
		findMsg.Name + `","` + findMsg.Bank + `","` + findMsg.BankNum + `",` + strconv.FormatInt(findMsg.UserID, 10) + `,"` +
		findMsg.NickName +
		`", now(), "` + investmentID +
		`")`
	if err = orm1.Exec(sql3).Error; err != nil {
		return
	}

	return
}
