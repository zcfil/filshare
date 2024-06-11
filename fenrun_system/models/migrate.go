package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/utils"
)

type Migrate struct {
	OrderMid      string    `gorm:"column:order_mid" json:"order_mid"`
	Time          time.Time `gorm:"column:time" json:"time"`
	LockedBalance string    `gorm:"column:locked_balance" json:"locked_balance"`
	TotalDays     string    `gorm:"column:total_days" json:"total_days"`
	StartTime     time.Time `gorm:"column:start_time" json:"start_time"`
	EndTIme       time.Time `gorm:"column:end_time" json:"end_time"`
	OneReward     string    `gorm:"column:one_reward" json:"one_reward"`
	UserId        string    `gorm:"column:user_id" json:"user_id"`
	PlusOne       string    `gorm:"column:plus_one" json:"plus_one"`
	Status        string    `gorm:"column:status" json:"status"`
	ActiveStatus  string    `gorm:"column:active_status" json:"active_status"`
	IsDel         string    `gorm:"column:is_del" json:"is_del"`
	Active        string    `gorm:"column:active" json:"active"`
	AdminId       string    `gorm:"column:admin_id" json:"admin_id"`
}

func NewMigrate() *Migrate {
	s := new(Migrate)
	return s
}

func (kthis *Migrate) GetListData(param map[string]string) (result interface{}, err error) {
	sqlx := `select order_mid, (select c.name FROM customer c where id = s.user_id) AS user_name,locked_balance AS balance,
       total_days AS totalDay,  start_time, end_time,  (select a.username from sys_user a where a.user_id = s.admin_id) admin_name,
       status ,remark ,user_id AS customer_id from migrate_order s WHERE is_del<>1`
	keyword := param["keyword"]
	if keyword != "" {
		sqlx += ` and (name like '%` + keyword + `%' or order_mid like '%` + keyword + `%')`
	}
	//总数
	param["total"] = GetTotalCount(sqlx)
	//分页 and 排序
	param["sort"] = "order_mid"
	param["order"] = "desc"
	sqlx += utils.LimitAndOrderBy(param)

	type magrate struct {
		OrderMid     string `gorm:"column:order_mid" json:"order_mid"`
		UserName     string `gorm:"column:user_name" json:"user_name"`
		UserId       string `gorm:"column:customer_id" json:"customer_id"`
		LockedBlance string `gorm:"column:balance" json:"balance"`
		TotalDays    string `gorm:"column:totalDay" json:"totalDay"`
		StartTime    string `gorm:"column:start_time" json:"start_time"`
		EndTime      string `gorm:"column:end_time" json:"end_time"`
		AdminName    string `gorm:"column:admin_name" json:"admin_name"`
		Status       string `gorm:"column:status" json:"status"`
		//Active       string `gorm:"column:active" json:"active"`
		Remark string `gorm:"column:remark" json:"remark"`
	}

	ret := make([]magrate, 0)
	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		log.Error("查询失败  ", err.Error())
		return
	}
	result = ret
	return result, nil
}

func (khtis *Migrate) RetrieveMigrate(orderMid string) error {
	sqlx := `Select   *  from migrate_order WHERE order_mid='` + orderMid + `' IS NULL `
	err := orm.Eloquent.Debug().Raw(sqlx).Error
	if err != nil {
		err = gorm.ErrRecordNotFound
		return err
	}
	return nil

}

func (kthis *Migrate) DeleteMigreate(orderMid string) error {
	sqlx := `UPDATE  migrate_order SET  is_del=1 ,active_status=1  WHERE order_mid='` + orderMid + `'`
	if err := orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		log.Error("ERROR  ", err.Error())
		return err
	}
	return nil
}

func (kthis *Migrate) EditMigrate(param map[string]string) error {
	reward1 := utils.StringToFloat64(param["balance"]) / utils.StringToFloat64(param["days"])
	reward := utils.Float64ToString(reward1)
	sqlx := `update  migrate_order set total_days='` + param["days"] + `',start_time='` + param["Time"] + `', end_time=date_add('` + param["Time"] + `', interval '` + param["days"] + `' day), locked_balance='` + param["balance"] + `' ,remark='` + param["remark"] + `',one_reward='` + reward + `' where order_mid='` + param["orderMid"] + `' AND user_id='` + param["customer_id"] + `'
`
	if err := orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return err
	}
	return nil
}

func (kthis *Migrate) BreakMigreate(orderMid string) error {
	sqlx := `UPDATE  migrate_order SET  status=1  ,active=0  WHERE order_mid='` + orderMid + `'`
	if err := orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		log.Error("ERROR  ", err.Error())
		return err
	}
	return nil
}

func (kthis *Migrate) BillingImpressions(param map[string]string) (ret interface{}, err error) {
	type settleLog struct {
		CustomerName        string    `gorm:"column:customer_name" json:"customer_name"`                 // 投资者名字
		CustomerID          string    `gorm:"column:customer_id" json:"customer_id"`                     // 投资者ID
		InvestmentID        string    `gorm:"column:investment_id" json:"investment_id"`                 // 投资ID
		CustomerLockRelease float64   `gorm:"column:customer_lock_release" json:"customer_lock_release"` // 用户锁仓释放
		TypesOf             string    `gorm:"column:name" json:"name"`
		SettleTime          time.Time `gorm:"column:time" json:"time"`               // 结算时间
		IsTransfer          string    `gorm:"column:is_transfer" json:"is_transfer"` // 是否已转账 1已转账， 0未转账
		//ToCustomerBalance   float64   `gorm:"column:to_customer_balance" json:"to_customer_balance"`     // 到用户余额
	}

	sql := `select c.name    customer_name,  s.investment_id,  s.customer_id,  s.customer_lock_release ,
       (SELECT name from tables_name WHERE kid = s.types_of)    name,  s.time,  (select  name from  tables_name where kid=s.is_transfer) is_transfer
     from settle_log s  left join customer c on s.customer_id = c.id WHERE type_kid=2 order by time desc`

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
	return ret, nil
}

func (kthis *Migrate) MigrateInsert(param map[string]string) error {
	id := strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	times := utils.TimeHMS()
	reward1 := utils.StringToFloat64(param["balance"]) / utils.StringToFloat64(param["days"])
	reward := utils.Float64ToString(reward1)
	sqlx := `INSERT INTO migrate_order ( order_mid, time, locked_balance, total_days, start_time, end_time, one_reward, user_id, plus_one,status,active_status,is_del,admin_id,remark) VALUES
                ('` + id + `','` + times + `','` + param["balance"] + `','` + param["days"] + `','` + times + `',date_add('` + times + `', interval ` + param["days"] + ` day),'` + reward + `','` + param["customer_id"] + `',0,0,0,0,'` + param["user_id"] + `','` + param["remark"] + `')`
	if err := orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		log.Error("insert 失败 ", err.Error())
		return err
	}
	return nil
}
