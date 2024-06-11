package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Profit struct {
	ID          string  `gorm:"column:id" json:"id"`
	Percent     float64 `gorm:"column:percent" json:"percent"`       //
	Userid      string  `gorm:"column:userid" json:"userid"`         //
	Profittype  int     `gorm:"column:profittype" json:"profittype"` //
	Profitlevel int64   `gorm:"column:profitlevel" json:"profitlevel"`
	//A			[]Aa			`gorm:"column:a" json:"a"`
	//IsDel		int			`gorm:"column:is_del" json:"is_del"`	//是否删除
	NickName   string    `gorm:"column:nick_name" json:"nick_name"` //
	Username   string    `gorm:"column:username" json:"username"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` //创建时间
}
type Aa struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func (e *Profit) ProfitconfigList(param map[string]string) (result interface{}, err error) {
	//状态
	status := param["profittype"]
	// fixme 用户vip等级在 user_level表里面需要单独判断获取，后期有时间可以改一下独立的接口
	if status == "2" {
		var userLevelConfig UserLevelConfig
		dataList, err1 := userLevelConfig.getProfitConfigList()
		if err1 != nil {
			err = err1
			return
		}
		ret := make([]Profit, 0, len(dataList))
		for _, data := range dataList {
			var p Profit
			p.ID = strconv.FormatInt(data.ID, 10)
			p.Profittype = 2
			p.Profitlevel = data.LevelValue
			p.Percent = data.Percent
			ret = append(ret, p)
		}
		result = ret
		return
	}

	//拼凑筛选条件sql
	sql := ` select p.*,u.nick_name,u.username from profitconfig p
		left join sys_user u ON p.userid = u.user_id
		where 1=1 `

	if param["profittype"] != "" {
		sql += ` and profittype = '` + status + `'`

		if param["profittype"] == "1" {
			sql = `select if(c.count1=0,0,p.id)id,p.percent,1 profittype,u.user_id,p.profitlevel,u.nick_name,u.username FROM(									
					select count(1) count1 from profitconfig p
					left join sys_user u ON p.userid = u.user_id
					where userid = :userid and profittype =1
				)c
				left join profitconfig p on p.userid = if(c.count1=0,0,:userid) and p.profittype = if(c.count1=0,2,1)
				left join sys_user u ON u.user_id = :userid`
			sql = utils.SqlReplaceParames(sql, param)
		}
	}
	userid := param["userid"]
	if param["userid"] != "" {
		sql += ` and userid = '` + userid + `'`
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "profitlevel"
	param["order"] = "asc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Profit, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Profit) ProfitconfigOnce(param map[string]string) (err error) {

	//param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(),10)
	var count int
	orm.Eloquent.Table("profitconfig").Where("userid = ? and is_del =0 ", param["userid"]).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}
	sql := ` insert into profitconfig(percent,profittype,userid )value(:percent,0,:userid)`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
func (e *Profit) DelProfitconfigOnce(param map[string]string) (err error) {

	sql := ` delete from profitconfig where id =:id and profittype =0 `
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
func (e *Profit) UpdateProfitconfigOnce(param map[string]string) (err error) {

	//param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(),10)

	sql := ` update profitconfig set percent=:percent,userid=:userid where id=:id `
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}

type ProfitEdit struct {
	Ids string `json:"ids"`
	//Userid string `json:"userid"`
	Profit []ProfitConfig `json:"profit"`
}
type ProfitConfig struct {
	//ID           string     `gorm:"column:id" json:"id"`
	Percent     float64 `gorm:"column:percent" json:"percent"`       //
	Userid      string  `gorm:"column:userid" json:"userid"`         //
	Profittype  int     `gorm:"column:profittype" json:"profittype"` //
	Profitlevel float64 `gorm:"column:profitlevel" json:"profitlevel"`
	//A			[]Aa			`gorm:"column:a" json:"a"`
	//IsDel		int			`gorm:"column:is_del" json:"is_del"`	//是否删除
	//CreateTime	time.Time			`gorm:"column:create_time" json:"create_time"`	//创建时间
}

func (e *ProfitEdit) ProfitEdit() (err error) {
	orm1 := orm.Eloquent.Begin()

	//增加新配置
	flag := false
	sql1 := `insert into user_level(levelvalue, percent, percentreality)values `

	percentReality := 0.0
	for k, v := range e.Profit {
		flag = true
		if k == 0 && v.Profitlevel != 0 {
			return nil
		}

		percentReality += v.Percent
		sql1 += `(` + utils.Float64ToString(v.Profitlevel) + "," + utils.Float64ToString(v.Percent) + "," + utils.Float64ToString(percentReality) + `)`
		if k < len(e.Profit)-1 {
			sql1 += ","
		}
	}
	//删除原有配置
	sql := ` delete from user_level where id in ('` + strings.Replace(e.Ids, ",", "','", -1) + `')`
	if err = orm1.Exec(sql).Error; err != nil {
		orm1.Rollback()
		return err
	}
	if flag {
		if err = orm1.Exec(sql1).Error; err != nil {
			orm1.Rollback()
			return err
		}
	}
	orm1.Commit()
	return
}

type TotalProfit struct {
	Total     string `json:"total"`
	Yesterday string `json:"yesterday"`
	Today     string `json:"today"`
}

func (t *TotalProfit) Profit(userid string) TotalProfit {
	sql := `select sum(total) total,sum(yesterday) yesterday,sum(today)today from(
	select sum(p.profits) total,0 yesterday,0 today
		from investment i
		left join investmentprofit p on i.id = p.investmentid
		where p.userid = ?
	UNION
	select 0 total,round(sum(p.profits)) yesterday,0 today from investment i 
			left join investmentprofit p on p.investmentid = i.id
			where create_time >= DATE_SUB(date(now()),INTERVAL 1 DAY) 
			and create_time < date(now()) and p.userid = ?
	UNION
	select 0 total,0 yesterday,round(sum(p.profits)) profits from investment i 
			left join investmentprofit p on p.investmentid = i.id
			where create_time < DATE_ADD(date(now()),INTERVAL 1 DAY) 
			and create_time >= date(now()) and p.userid = ?
			)a`
	var tp TotalProfit
	orm.Eloquent.Raw(sql, userid, userid, userid).Scan(&tp)
	return tp
}
