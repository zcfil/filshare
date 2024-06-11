package models

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Investment struct {
	ID         string    `gorm:"column:id" json:"id"`
	Userid     string    `gorm:"column:user_id" json:"user_id"`
	Customerid string    `gorm:"column:customer_id" json:"customer_id"` //客户id
	Remark     string    `gorm:"column:remark" json:"remark"`           //备注
	Name       string    `gorm:"column:name" json:"name"`               //备注
	Phone      string    `gorm:"column:phone" json:"phone"`             //手机号
	Storage    string    `gorm:"column:storage" json:"storage"`         //备注
	IsDel      int       `gorm:"column:is_del" json:"is_del"`           //是否删除
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` //创建时间
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`   //创建时间
	EndTime    time.Time `gorm:"column:end_time" json:"end_time"`       //创建时间
	Status     string    `gorm:"column:status" json:"status"`
	TotalDay   string    `gorm:"column:totalDay" json:"totalDay"`
	Days       string    `gorm:"column:days" json:"days"` //新增释放周期字段
}
type InvestmentExport struct {
	ID             string `gorm:"column:id" json:"id"`
	Userid         string `gorm:"column:userid" json:"userid"`
	Amount         string `gorm:"column:amount" json:"amount"`         //金额
	Customerid     string `gorm:"column:customerid" json:"customerid"` //客户id
	Remark         string `gorm:"column:remark" json:"remark"`         //备注
	IsDel          string `gorm:"column:is_del" json:"is_del"`         //是否删除
	Status         string `gorm:"column:status" json:"status"`
	CreateTime     string `gorm:"column:create_time" json:"create_time"`         //创建时间
	UpdateTime     string `gorm:"column:update_time" json:"update_time"`         //修改时间
	InvestTime     string `gorm:"column:invest_time" json:"invest_time"`         //创建时间
	ExpirationDate string `gorm:"column:expiration_date" json:"expiration_date"` //创建时间
	MonthlyTime    string `gorm:"column:monthly_time" json:"monthly_time"`       //创建时间
	Profit         string `gorm:"column:profit" json:"profit"`
	Oncestatus     string `gorm:"column:oncestatus" json:"oncestatus"`
}
type InvestmentprofitExport struct {
	ID           string `gorm:"column:id" json:"id"`
	Investmentid string `gorm:"column:investmentid" json:"investmentid"`
	Userid       string `gorm:"column:userid" json:"userid"`   //金额
	Profits      string `gorm:"column:profits" json:"profits"` //客户id
}
type InvestmentView struct {
	Investment
	Summaryid string `json:"summaryid"`
}

func NewInvestment(param map[string]string) error {
	sql2 := `select c.*,i.amount,i.remark from customer c
						left join investment i on c.id = i.customerid and i.id = '` + param["investmentid"] + `'
						where c.id = '` + param["customerid"] + "'"
	var c Customer
	if err := orm.Eloquent.Raw(sql2).Scan(&c).Error; err != nil {
		fmt.Errorf("获取客户信息失败：", err)
		return err
	}
	param["name"] = c.Name
	param["phone"] = c.Phone
	// param["bank"] = c.Bank
	// param["banknum"] = c.Banknum
	param["identity"] = c.Identity
	// param["sex"] = c.Sex
	param["identity"] = c.Identity
	// param["remark"] = c.Remark
	// param["amount"] = utils.Float64ToString(c.Amount)
	return nil
}
func GetInvestmentById(id string) Investment {
	sql2 := ` select * from investment where id = '` + id + "'"
	var in Investment
	if err := orm.Eloquent.Raw(sql2).Scan(&in).Error; err != nil {
		return in
	}
	return in
}

func (e *Investment) InvestmentList(param map[string]string) (result interface{}, err error) {
	sql := `select id,customer_id, name,phone,storage,remark,user_id,status, create_time,update_time,start_time,end_time ,total_day AS  totalDay,days from investment where is_del <> 1` //23/3/30 新增释放周期字段
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (name like '%` + keyword + `%' or id like '%` + keyword + `%')`
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Investment, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Investment) InvestmentAdd(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	var c Customer
	c, err = c.NewCustomer(param["customer_id"])
	if err != nil {
		return err
	}
	param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(), 10)

	//type days struct {
	//	Name string `gorm:"column:name" json:"name"`
	//}
	//sqlx := `SELECT  value AS  name  from sys_config where  alies_name="day"`
	//days1 := make([]days, 0)
	//
	//if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&days1).Error; err != nil {
	//	return
	//}
	//dayk := days1[0].Name
	//= c.Request.FormValue("totalDay")
	dayk := param["dayk"]
	//sql := ` insert into investment(id,user_id,name,customer_id,remark,storage,status,start_time,end_time,active_mid)value(:id,:user_id,:name,:customer_id,:remark,:storage,0,now(),date_add(now(), interval 180 day),0)`
	sql := ` insert into investment(id,user_id,name,phone,customer_id,remark,storage,status,start_time,end_time,total_day,days)value(:id,:user_id,:name,:phone,:customer_id,:remark,:storage,0, now(),date_add('` + utils.TimeHMS() + `', interval ` + dayk + ` day),'` + dayk + `',:days)` //23/3/8添加days字段
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}
	return
}

func (e *Investment) InvestmentEdit(param map[string]string) (err error) {
	sql := ` update investment set storage="%s",remark="%s", update_time='` + utils.TimeHMS() + `' ,start_time='` + param["time"] + `',end_time=date_add('` + param["time"] + `', interval ` + param["dayk"] + ` day) ,total_day='` + param["dayk"] + `',days='` + param["days"] + `'  where id = %s` //23/3/13 新增释放周期字段
	sql = fmt.Sprintf(sql, param["storage"], param["remark"], param["id"])
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return
	}
	return
}

func (e *Investment) InvestmentDelete(param map[string]string) (err error) {
	sql := ` update investment set is_del=1 where id =` + param["id"]
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return
	}
	return
}
func (e *Investment) InvestmentBreak(param map[string]string) (err error) {
	sql := ` update investment set status=1 where id =` + param["id"]
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return
	}
	return
}

func (e *Investment) InvestmentRevoke(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()

	sql := ` DELETE FROM investment WHERE id = :investmentid`
	sql = utils.SqlReplaceParames(sql, param)
	err = orm1.Exec(sql).Error

	if err = NewInvestment(param); err != nil {
		return err
	}

	//上级所有人都减业绩
	in := GetInvestmentById(param["investmentid"])
	var ref Referrer
	ref = ref.GetReferrer(in.Userid)
	ref.Referrers = strings.Trim(ref.Referrers, ",")
	// if strings.Trim(ref.Referrers, " ") != "" {
	// 	sql2 := `update sys_user set accumulative = accumulative + real_value -` + utils.Float64ToString(in.Amount) + `,real_value=0 where user_id in (` + ref.Referrers + `) and level_lock=0`
	// 	if err = orm1.Exec(sql2).Error; err != nil {
	// 		return
	// 	}
	// 	sql4 := `update sys_user set real_value = real_value -` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) and level_lock=1`
	// 	if err = orm1.Exec(sql4).Error; err != nil {
	// 		return
	// 	}
	// }

	// sql3 := `update sys_user set accumulative = accumulative + real_value -` + utils.Float64ToString(in.Amount) + `,real_value=0 where user_id in (` + in.Userid + `) `
	// if err = orm1.Exec(sql3).Error; err != nil {
	// 	return
	// }

	sql1 := ` DELETE FROM customer_log WHERE investmentid = :investmentid`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error
	return
}

func (e *Investment) InvestmentById(id string) (res interface{}, err error) {
	sql := `select * from investment where id = ?`
	var in Investment
	err = orm.Eloquent.Raw(sql, id).Scan(&in).Error
	return in, err
}

func (e *Investment) ExportInvestmentList(param map[string]string) (interface{}, error) {
	sql := `select id,userid,amount,customerid,remark,is_del,status
			,DATE_FORMAT(create_time,'%Y-%m-%d %H:%i:%s') create_time
			,DATE_FORMAT(update_time,'%Y-%m-%d %H:%i:%s') update_time
			,DATE_FORMAT(invest_time,'%Y-%m-%d %H:%i:%s') invest_time
			,DATE_FORMAT(expiration_date,'%Y-%m-%d %H:%i:%s') expiration_date
			,DATE_FORMAT(monthly_time,'%Y-%m-%d %H:%i:%s') monthly_time,profit,oncestatus from investment order by invest_time asc `
	var ie []InvestmentExport
	err := orm.Eloquent.Raw(sql).Scan(&ie).Error

	return ie, err
}

//导出一次性分润报表
func (e *Investment) ExportInvestment(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["filefield"] = "id,amount,customerid,create_time,update_time,is_del,remark,status,invest_time,expiration_date,userid,monthly_time,profit,oncestatus"
	//param["filename"] = "ID,金额,客户ID,创建时间,修改时间,是否已删除,备注,状态,最初创建时间,合同到期时间,业务员ID,客户分润时间,客户利润点,一次性分润状态"
	param["filename"] = "id,amount,customerid,create_time,update_time,is_del,remark,status,invest_time,expiration_date,userid,monthly_time,profit,oncestatus"
	param["title"] = "投资列表导出" + param["date1"]
	URL, err = GetExcelURL(e.ExportInvestmentList, param)
	return
}

//删除订单
func (e *Investment) DeleteInvestment() (err error) {
	//先备份
	sees := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			sees.Rollback()
			return
		}
		sees.Commit()
	}()
	//sql := `select * from investment `
	//var ie []InvestmentExport
	//if err = orm.Eloquent.Raw(sql).Scan(&ie).Error;err!=nil{
	//	return err
	//}
	inv, err := e.ExportInvestmentList(nil)
	if err != nil {
		return nil
	}
	ie := inv.([]InvestmentExport)
	//获取备份批次
	sql0 := `select max(batch)+1 batch from investment_backup `
	type Batchs struct {
		Batch string
	}
	var b Batchs
	if err = orm.Eloquent.Raw(sql0).Scan(&b).Error; err != nil {
		return err
	}
	if b.Batch == "" {
		b.Batch = "0"
	}
	sql1 := `insert into investment_backup(investment_id,amount,customerid,create_time,update_time,is_del,remark,status
						,invest_time,expiration_date,userid,monthly_time,profit,oncestatus,batch)values`
	for k, v := range ie {
		sql1 += `('` + v.ID + `','` + v.Amount + `','` + v.Customerid + `','` + v.CreateTime + `','` + v.UpdateTime + `','` + v.IsDel + `','` + v.Remark + `','` + v.Status + `','` + v.InvestTime + `
					','` + v.ExpirationDate + `','` + v.Userid + `','` + v.MonthlyTime + `','` + v.Profit + `','` + v.Oncestatus + `','` + b.Batch + `')`
		if k < len(ie)-1 {
			sql1 += ","
		}
	}
	if len(ie) > 0 {
		if err = sees.Exec(sql1).Error; err != nil {
			return err
		}
	}
	//备份分润信息
	sql2 := `select * from investmentprofit  `
	var ipe []InvestmentprofitExport
	if err = orm.Eloquent.Raw(sql2).Scan(&ipe).Error; err != nil {
		return err
	}
	sql3 := `insert into investmentprofit_backup(investmentprofit_id,investmentid,userid,profits,batch)values`
	for k, v := range ipe {
		sql3 += `('` + v.ID + `','` + v.Investmentid + `','` + v.Userid + `','` + v.Profits + `','` + b.Batch + `')`
		if k < len(ipe)-1 {
			sql3 += ","
		}
	}
	if len(ie) > 0 {
		if err = sees.Exec(sql3).Error; err != nil {
			return err
		}
	}

	sql4 := `delete from investment `
	if err = sees.Exec(sql4).Error; err != nil {
		return err
	}

	sql5 := `delete from investmentprofit `
	if err = sees.Exec(sql5).Error; err != nil {
		return err
	}
	sql6 := `update sys_user set accumulative = 0 `
	if err = sees.Exec(sql6).Error; err != nil {
		return err
	}

	return
}

//导入订单
func (e *Investment) InvestmentImport(file multipart.File, Size int64) error {
	buf := make([]byte, Size)
	n, _ := file.Read(buf)
	//sees := orm.Eloquent.Begin()
	//var err error
	//defer func() {
	//	if err!=nil{
	//		sees.Rollback()
	//		return
	//	}
	//	sees.Commit()
	//}()

	xf, _ := xlsx.OpenBinary(buf[:n])
	for _, sheet := range xf.Sheets {
		if len(sheet.Rows) == 0 {
			continue
		}

		//sql1 = `insert into investment(id,amount,customerid,create_time,update_time,is_del,remark,status
		//				,invest_time,expiration_date,userid,monthly_time,profit,oncestatus)values`
		//记录字段位置
		filedvalue := make(map[string]int)
		sql1 := `insert into investment(`
		for j, row := range sheet.Rows {
			//保存参数
			sql2 := ""
			param := make(map[string]string)
			for i, cell := range row.Cells {
				if j == 0 {
					sql1 += cell.String()
					if i < len(row.Cells)-1 {
						sql1 += ","
					} else {
						sql1 += ")values"
					}
					filedvalue[cell.String()] = i
				} else {
					if i == 0 {
						//if j > 1{
						//	sql1 += ","
						//}
						sql2 += "("
					}
					sql2 += "'" + cell.String() + "'"
					if i < len(row.Cells)-1 {
						sql2 += ","
					} else {
						sql2 += ")"
					}
					//保存字段
					switch i {
					case filedvalue["amount"]:
						param["amount"] = cell.String()
					case filedvalue["userid"]:
						param["userid"] = cell.String()
					case filedvalue["id"]:
						param["investmentid"] = cell.String()
					case filedvalue["status"]:
						param["status"] = cell.String()
					}
				}

			}
			if j > 0 {
				var i InvestmentShareProfit
				//err = i.AddInvestmentShareProfitImport(sees,param)
				if param["status"] != "1" {
					err := i.AddInvestmentShareProfitImport(param)
					if err != nil {
						return err
					}
					if err = orm.Eloquent.Exec(sql1 + sql2).Error; err != nil {
						return err
					}
				}

				time.Sleep(time.Millisecond * 100)

			}

		}
		//if err = sees.Exec(sql1).Error;err!=nil{
		//if err = orm.Eloquent.Exec(sql1).Error;err!=nil{
		//	return err
		//}
	}
	return nil
}
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func (e *Investment) GetAccumulative(id string) (string, error) {
	var ref Referrer
	ref = ref.GetReferrer(id)
	als := strings.Trim(ref.Referrals, ",")
	sql := `select sum(amount)accumulative from investment where userid in(` + als + `) `
	type Acc struct {
		Accumulative string
	}
	var acc Acc
	if err := orm.Eloquent.Raw(sql).Scan(&acc).Error; err != nil {
		return "", nil
	}
	return acc.Accumulative, nil
}

func (e *Investment) Check_stment() (err error) { // 生效订单 进行收益计算开始
	sqlx := `UPDATE  investment SET  active_mid=1 ,update_time=NOW(),start_time=DATE_FORMAT(NOW(),"%Y-%m-%d") ,end_time=DATE_FORMAT( DATE_ADD(NOW(), INTERVAL 180 DAY),"%Y-%m-%d")  WHERE status=1 AND is_del = 0`
	return orm.Eloquent.Debug().Exec(sqlx).Error
}
