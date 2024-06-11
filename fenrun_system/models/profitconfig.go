package models

import (
	"strconv"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type SysProfitConfig struct {
	Id      string  `gorm:"column:id" json:"id"`
	Cvalue  string  `gorm:"column:cvalue" json:"cvalue"`
	Cname   string  `gorm:"column:cname" json:"cname"`
	Percent float64 `gorm:"column:percent" json:"percent"`
	Userid  string  `gorm:"column:userid" json:"userid"`
}
type SysProfitSalesman struct {
	Id      string  `gorm:"column:id" json:"id"`
	Amount  float64 `gorm:"column:amount" json:"amount"`
	Profits float64 `gorm:"column:profits" json:"profits"`
	Userid  string  `gorm:"column:userid" json:"userid"`
}

//一次性分配利润
func NewProfitconfig(profittype int) []SysProfitConfig {
	sql := `select p.*,u.nick_name cname,u.username cvalue from profitconfig p 
			left join sys_user u on p.userid = u.user_id
			where profittype=` + strconv.Itoa(profittype)
	var p []SysProfitConfig
	orm.Eloquent.Raw(sql).Scan(&p)
	return p
}

func NewProfitSalesman(param map[string]string) ([]SysProfitSalesman, Total) {
	sql := `select amount,a.userid,a.id,sum(p1.levelgain)+(amount-max(p1.profitlevel))*sum(percent) profits from (
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
				GROUP BY a.id`
	sql = utils.SqlReplaceParames(sql, param)
	var t Total
	orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t)

	var p []SysProfitSalesman
	orm.Eloquent.Raw(sql).Scan(&p)
	return p, t

}
