package models

import (
	"errors"
	"strings"
	"sync"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

//type CustomerLog struct {
//	ID           string     `gorm:"column:id" json:"id"`
//	Salesmanid    	 string    `gorm:"column:salesmanid" json:"salesmanid"`               //
//	Name        string    `gorm:"column:name" json:"name"`            //
//	Namenew        string    `gorm:"column:namenew" json:"namenew"`
//	Phone        string    `gorm:"column:phone" json:"phone"`
//	Phonenew        string    `gorm:"column:phonenew" json:"phonenew"`
//	Investmentid    	 string    `gorm:"column:investmentid" json:"investmentid"`
//	Amount    	 string    `gorm:"column:amount" json:"amount"`
//	Amountnew    	 string    `gorm:"column:amountnew" json:"amountnew"`
//	IsDel		int			`gorm:"column:is_del" json:"is_del"`	//
//	CreateTime	time.Time			`gorm:"column:create_time" json:"create_time"`	//
//	UpdateTime	time.Time			`gorm:"column:update_time" json:"update_time"`	//创建时间
//	Bank        string    `gorm:"column:bank" json:"bank"`
//	Banknew        string    `gorm:"column:banknew" json:"banknew"`
//	Banknum        string    `gorm:"column:banknum" json:"banknum"`
//	Banknumnew        string    `gorm:"column:banknumnew" json:"banknumnew"`
//	Edittype		int			`gorm:"column:edittype" json:"edittype"`	//
//	Customerid    	 string    `gorm:"column:customerid" json:"customerid"`
//	Iremark    	 string    `gorm:"column:iremark" json:"iremark"`
//	Iremarknew    	 string    `gorm:"column:iremarknew" json:"iremarknew"`
//	Status		int			`gorm:"column:status" json:"status"`	//
//	Identity    	 string    `gorm:"column:identity" json:"identity"`
//	Identitynew    	 string    `gorm:"column:identitynew" json:"identitynew"`
//	Sex		string			`gorm:"column:sex" json:"sex"`	//
//	Sexnew		string			`gorm:"column:sexnew" json:"sexnew"`	//
//}
type CustomerLog struct {
	ID           string    `gorm:"column:id" json:"id"`
	Salesmanid   string    `gorm:"column:salesmanid" json:"salesmanid"` //
	Name         string    `gorm:"column:name" json:"name"`             //
	Phone        string    `gorm:"column:phone" json:"phone"`
	Investmentid string    `gorm:"column:investmentid" json:"investmentid"`
	Amount       string    `gorm:"column:amount" json:"amount"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"` //
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"` //创建时间
	Bank         string    `gorm:"column:bank" json:"bank"`
	Banknum      string    `gorm:"column:banknum" json:"banknum"`
	Edittype     int       `gorm:"column:edittype" json:"edittype"` //
	Customerid   string    `gorm:"column:customerid" json:"customerid"`
	Iremark      string    `gorm:"column:iremark" json:"remark"`
	Status       int       `gorm:"column:status" json:"status"` //
	Identity     string    `gorm:"column:identity" json:"identity"`
	Sex          string    `gorm:"column:sex" json:"sex"` //
	Profit       string    `gorm:"column:profit" json:"profit"`
	Flag         string    `json:"flag"`
}
type CustomerLogNew struct {
	Namenew      string    `gorm:"column:namenew" json:"namenew"`
	Phonenew     string    `gorm:"column:phonenew" json:"phonenew"`
	Amountnew    string    `gorm:"column:amountnew" json:"amountnew"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"` //
	Banknew      string    `gorm:"column:banknew" json:"banknew"`
	Banknumnew   string    `gorm:"column:banknumnew" json:"banknumnew"`
	Edittype     string    `gorm:"column:edittype" json:"edittype"` //
	Customerid   string    `gorm:"column:customerid" json:"customerid"`
	Investmentid string    `gorm:"column:investmentid" json:"investmentid"`
	Iremarknew   string    `gorm:"column:iremarknew" json:"iremarknew"`
	Status       int       `gorm:"column:status" json:"status"` //
	Identitynew  string    `gorm:"column:identitynew" json:"identitynew"`
	Sexnew       string    `gorm:"column:sexnew" json:"sexnew"` //
	Salesmanid   string    `gorm:"column:salesmanid" json:"salesmanid"`
}
type CustomerLogView struct {
	CustomerLog
	NickName string `json:"nick_name"`
}

func NewCustomerLogNew(param map[string]string) error {
	sql := `select * from customer_log where id = '` + param["id"] + `'`
	var c CustomerLogNew
	if err := orm.Eloquent.Raw(sql).Scan(&c).Error; err != nil {
		return err
	}
	param["namenew"] = c.Namenew
	param["phonenew"] = c.Phonenew
	param["amountnew"] = c.Amountnew
	param["banknew"] = c.Banknew
	param["banknumnew"] = c.Banknumnew
	param["customerid"] = c.Customerid
	param["investmentid"] = c.Investmentid
	param["iremarknew"] = c.Iremarknew
	param["identitynew"] = c.Identitynew
	param["sexnew"] = c.Sexnew
	param["edittype"] = c.Edittype
	param["userid"] = c.Salesmanid

	return nil
}
func (e *CustomerLog) CustomerAuditList(param map[string]string) (result interface{}, err error) {
	//拼凑筛选条件sql
	con := ``
	if !AUDIT[param["role"]] {
		con = ` and c.salesmanid = ` + param["userid"]
	}
	sql := ` select if(name=if(ifnull(namenew,'')='',name,namenew),name,CONCAT(name,' 改为 ',namenew))name,
		if(c.phone=if(ifnull(phonenew,'')='',c.phone,phonenew),c.phone,CONCAT(c.phone,' 改为 ',phonenew))phone,
		if(bank=if(ifnull(banknew,'')='',bank,banknew),bank,CONCAT(bank,' 改为 ',banknew))bank, 
		if(banknum=if(ifnull(banknumnew,'')='',banknum,banknumnew),banknum,CONCAT(banknum,' 改为 ',banknumnew))banknum, 
		if(iremark=if(ifnull(iremarknew,'')='',iremark,iremarknew),iremark,CONCAT(iremark,' 改为 ',iremarknew))iremark, 
		if(identity=if(ifnull(identitynew,'')='',identity,identitynew),identity,CONCAT(identity,' 改为 ',identitynew))identity,
		if(c.sex=if(ifnull(sexnew,'')='',c.sex,sexnew),c.sex,CONCAT(c.sex,' 改为 ',sexnew))sex,edittype,
		investmentid,
		if(amount=if(ifnull(amountnew,'')='',amount,amountnew),amount,CONCAT(amount,' 改为 ',amountnew))amount,
		c.create_time,c.update_time,c.status,u.nick_name,c.id
	from customer_log c
	left JOIN sys_user u on c.salesmanid = u.user_id
	where 1=1 ` + con
	//状态
	status := param["status"]
	if status != "" {
		sql += ` and c.status in ('` + strings.Replace(status, ",", "','", -1) + `')`
	}
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.phone like '%` + keyword + `%' or c.name like '%` + keyword + `%') `
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]CustomerLog, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}
func (e *CustomerLog) CustomerLogList(param map[string]string) (result interface{}, err error) {
	//拼凑筛选条件sql
	con := ``
	if !AUDIT[param["role"]] {
		con = ` and c.salesmanid = ` + param["userid"]
	}
	sql := ` select if(name=ifnull(namenew,name),name,CONCAT(name,' 改为 ',namenew))name,
		if(c.phone=ifnull(phonenew,c.phone),c.phone,CONCAT(c.phone,' 改为 ',phonenew))phone,
		if(bank=ifnull(banknew,bank),bank,CONCAT(bank,' 改为 ',banknew))bank, 
		if(banknum=ifnull(banknumnew,banknum),banknum,CONCAT(banknum,' 改为 ',banknumnew))banknum, 
		if(iremark=ifnull(iremarknew,iremark),iremark,CONCAT(iremark,' 改为 ',iremarknew))iremark, 
		if(identity=ifnull(identitynew,identity),identity,CONCAT(identity,' 改为 ',identitynew))identity,
		if(c.sex=ifnull(sexnew,c.sex),c.sex,CONCAT(c.sex,' 改为 ',sexnew))sex,edittype,
		investmentid,
		if(amount=ifnull(amountnew,amount),amount,CONCAT(amount,' 改为 ',amountnew))amount,
		c.create_time,c.update_time,c.status,u.nick_name,c.id
	from customer_log c
	left JOIN sys_user u on c.salesmanid = u.user_id
	where 1=1  ` + con
	//状态
	status := param["status"]
	if status != "" {
		sql += ` and c.status <> '` + status + `'`
	}
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.phone like '%` + keyword + `%' or c.name like '%` + keyword + `%') `
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]CustomerLog, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}
func (e *CustomerLog) GetCustomerLog(id, role string) (CustomerLogView, error) {
	//拼凑筛选条件sql
	con := ",0 flag"
	if AUDIT[role] {
		con = ",1 flag"
	}
	sql := `select if(name=ifnull(namenew,name),name,CONCAT(name,' 改为 ',namenew))name,
		if(c.phone=ifnull(phonenew,c.phone),c.phone,CONCAT(c.phone,' 改为 ',phonenew))phone,
		if(bank=ifnull(banknew,bank),bank,CONCAT(bank,' 改为 ',banknew))bank, 
		if(banknum=ifnull(banknumnew,banknum),banknum,CONCAT(banknum,' 改为 ',banknumnew))banknum, 
		if(iremark=ifnull(iremarknew,iremark),iremark,CONCAT(iremark,' 改为 ',iremarknew))iremark, 
		if(identity=ifnull(identitynew,identity),identity,CONCAT(identity,' 改为 ',identitynew))identity,
		if(c.sex=ifnull(sexnew,c.sex),c.sex,CONCAT(c.sex,' 改为 ',sexnew))sex,edittype,
		investmentid,
		if(amount=ifnull(amountnew,amount),amount,CONCAT(amount,' 改为 ',amountnew))amount,
		c.create_time,c.update_time,c.status,u.nick_name,c.id` + con + `
	from customer_log c
	left JOIN sys_user u on c.salesmanid = u.user_id
	where id = '` + id + `'`
	var c CustomerLogView
	err := orm.Eloquent.Raw(sql).Scan(&c).Error
	return c, err
}

var audit sync.Mutex

func (e *CustomerLog) CustomerLogAudit(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		audit.Unlock()
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	audit.Lock()
	sql := ` update customer_log set status = :status where id = :id and status <> :status`
	sql = utils.SqlReplaceParames(sql, param)
	r := orm1.Exec(sql)
	if r.RowsAffected == 0 {
		return errors.New("该记录不存在或已审核过")
	}
	//获取审核类型
	if err = NewCustomerLogNew(param); err != nil {
		return
	}
	sql1 := ""
	if param["status"] == "0" {
		//通过
		switch param["edittype"] {
		case "0":
			//新增客户审核
			sql1 = ` update customer set status = 0,is_del=0 where id = :customerid`

			//新增投资审核
			var i InvestmentShareProfit
			if err = i.AddInvestmentShareProfit(orm1, param); err != nil {
				return err
			}
			sql2 := ` update investment set status = 0,is_del=0,update_time=now() where id = :investmentid`
			sql2 = utils.SqlReplaceParames(sql2, param)
			if err = orm1.Exec(sql2).Error; err != nil {
				return err
			}
			//上级所有人都累加业绩
			in := GetInvestmentById(param["investmentid"])
			var ref Referrer
			ref = ref.GetReferrer(in.Userid)

			ref.Referrers += "," + in.Userid
			ref.Referrers = strings.Trim(ref.Referrers, ",")
			// sql3 := `update sys_user set accumulative = accumulative +` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) `
			// if err = orm1.Exec(sql3).Error; err != nil {
			// 	return
			// }
		case "1":
			//修改客户审核
			sql1 = `update customer set name = :namenew, phone = :phonenew,bank=:banknew,banknum=:banknumnew
				,identity=:identitynew,sex=:sexnew,update_time=now(),status=0
				where id =:customerid `
		case "2":
			//新增投资审核
			var i InvestmentShareProfit
			if err = i.AddInvestmentShareProfit(orm1, param); err != nil {
				return err
			}
			//sql1 = ` update investment set status = 0,is_del=0,create_time=now(),update_time=now(),monthly_time=now() where id = :investmentid`
			sql1 = ` update investment set status = 0,is_del=0,update_time=now() where id = :investmentid`
			//上级所有人都累加业绩
			in := GetInvestmentById(param["investmentid"])
			var ref Referrer
			ref = ref.GetReferrer(in.Userid)
			// ref.Referrers += "," + in.Userid
			ref.Referrers = strings.Trim(ref.Referrers, ",")
			// if strings.Trim(ref.Referrers, " ") != "" {
			// 	sql2 := `update sys_user set accumulative = accumulative + real_value +` + utils.Float64ToString(in.Amount) + `,real_value=0 where user_id in (` + ref.Referrers + `) and level_lock=0`
			// 	if err = orm1.Exec(sql2).Error; err != nil {
			// 		return
			// 	}
			// 	sql4 := `update sys_user set real_value = real_value +` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) and level_lock=1`
			// 	if err = orm1.Exec(sql4).Error; err != nil {
			// 		return
			// 	}
			// }

			// sql3 := `update sys_user set accumulative = accumulative + real_value +` + utils.Float64ToString(in.Amount) + `,real_value=0 where user_id in (` + in.Userid + `) `
			// if err = orm1.Exec(sql3).Error; err != nil {
			// 	return
			// }

		case "3":
			//删除分润
			sql2 := `delete from investmentprofit where investmentid = :investmentid `
			sql2 = utils.SqlReplaceParames(sql2, param)
			if err = orm1.Exec(sql2).Error; err != nil {
				return
			}
			in := GetInvestmentById(param["investmentid"])
			param["userid"] = in.Userid
			var ref Referrer
			ref = ref.GetReferrer(in.Userid)
			ref.Referrers += "," + in.Userid
			ref.Referrers = strings.Trim(ref.Referrers, ",")
			orm2 := orm.Eloquent.Begin()
			//删除累积业绩
			// sql3 := `update sys_user set accumulative = accumulative - ` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) `
			// if err = orm2.Exec(sql3).Error; err != nil {
			// 	orm2.Rollback()
			// 	return
			// }
			//修改投资审核
			sql1 = ` update investment set amount = :amountnew,remark=:iremarknew,update_time=now(),status=if(status=5||status=6,status,0) where id = :investmentid`
			sql1 = utils.SqlReplaceParames(sql1, param)
			if err = orm2.Exec(sql1).Error; err != nil {
				orm2.Rollback()
				return
			}
			orm2.Commit()
			//新增分润
			var i InvestmentShareProfit
			if err = i.AddInvestmentShareProfit(orm1, param); err != nil {
				return err
			}
			//所有人加上累积业绩
			sql4 := `update sys_user set accumulative = accumulative +` + param["amountnew"] + ` where user_id in (` + ref.Referrers + `) `
			if err = orm1.Exec(sql4).Error; err != nil {
				return
			}
			return
		case "4":
			//终止投资审核
			sql1 = ` update investment set status = 5,expiration_date = date(now()),manually_end = 1 where id = :investmentid`
			//2021年10月12日14:17:00 新增终止逻辑
			//减掉相关业务员业绩
			//删除分润
			sql2 := `delete from investmentprofit where investmentid = :investmentid `
			sql2 = utils.SqlReplaceParames(sql2, param)
			if err = orm1.Exec(sql2).Error; err != nil {
				return
			}
			in := GetInvestmentById(param["investmentid"])
			param["userid"] = in.Userid
			var ref Referrer
			ref = ref.GetReferrer(in.Userid)
			ref.Referrers += "," + in.Userid
			ref.Referrers = strings.Trim(ref.Referrers, ",")
			//删除累积业绩
			// sql3 := `update sys_user set accumulative = accumulative - ` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) `
			// if err = orm1.Exec(sql3).Error; err != nil {
			// 	orm1.Rollback()
			// 	return
			// }
		case "5":
			//已拒绝新增投资的修改
			in := GetInvestmentById(param["investmentid"])
			var ref Referrer
			ref = ref.GetReferrer(in.Userid)
			ref.Referrers += "," + in.Userid
			ref.Referrers = strings.Trim(ref.Referrers, ",")
			//修改投资审核
			sql1 = ` update investment set amount = :amountnew,remark=:iremarknew,create_time=now(),update_time=now(),status=if(status=5||status=6,status,0) where id = :investmentid`
			//新增分润
			var i InvestmentShareProfit
			if err = i.AddInvestmentShareProfit(orm1, param); err != nil {
				return err
			}
			//所有人加上累积业绩
			sql4 := `update sys_user set accumulative = accumulative +` + param["amountnew"] + ` where user_id in (` + ref.Referrers + `) `
			if err = orm1.Exec(sql4).Error; err != nil {
				return
			}
		}

	} else {
		//驳回
		switch param["edittype"] {
		case "0":
			sql1 = `update customer set status =:status where id = :customerid `
			sql5 := `update investment set status = 3 where id = :investmentid `
			sql5 = utils.SqlReplaceParames(sql5, param)
			if err = orm1.Exec(sql5).Error; err != nil {
				return err
			}
		case "1":
			sql1 = `update customer set status =:status where id = :customerid `
		case "2":
			//新增投资审核
			sql1 = `update investment set status = 3 where id = :investmentid `
		case "3":
			//修改投资审核
			sql1 = `update investment set status = 0 where id = :investmentid `
		case "4":
			//终止投资审核
			sql1 = `update investment set status = 0 where id = :investmentid `
		case "5":
			//已拒绝新增的修改投资审核
			sql1 = `update investment set status = 3 where id = :investmentid `
		}
	}
	sql1 = utils.SqlReplaceParames(sql1, param)

	return orm1.Exec(sql1).Error
}

//type editType struct {
//	Id string
//	Edittype int
//}
//func newEditType(id string)editType{
//	var e editType
//	sql := `select if(edittype<=1,customerid,investmentid)id,edittype from customer_log where id = '`+id+`'`
//	orm.Eloquent.Raw(sql).Scan(&e)
//	return e
//}
