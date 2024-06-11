package models

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type UserLevel struct {
	Percentreality float64 `gorm:"column:percentreality" json:"percentreality"`
	UserId         string  `gorm:"column:user_id" json:"user_id"`
	Accumulative   float64 `gorm:"column:accumulative" json:"accumulative"`
	Levelvalue     float64 `gorm:"column:levelvalue" json:"levelvalue"`
	Percent        float64 `gorm:"column:percent" json:"percent"`
	Levelname      string  `gorm:"column:levelname" json:"levelname"`
}

//func (u *UserLevel)GetUserLevelList(ids string)([]UserLevel,error){
//	con := `and user_id in (`+ids+`)`
//	sql := `select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
//					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
//					where 1=1 `+con+`
//					GROUP BY user_id order by levelvalue`
//	var ul []UserLevel
//	err := orm.Eloquent.Raw(sql).Scan(&ul).Error
//	for i:=0;i<len(ul);i++{
//		if i < len(ul)-1{
//			if ul[i].Levelvalue == ul[i+1].Levelvalue{
//
//			}
//		}
//	}
//	return ul,err
//}
////排除一样等级的
//func (u *UserLevel)GetUserLevel(ids string,levelvalue float64)([]UserLevel,error){
//	con := `and user_id in (`+ids+`)`
//	sql := `select u.user_id,u.accumulative, b.levelvalue, b.percent from sys_user u
//					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
//					left join (
//							select count(1) count,levelvalue,sum(percent) percent from (
//							select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
//							left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
//							where 1=1 `+con+` and l.levelvalue > ?
//							GROUP BY user_id order by levelvalue
//							)a GROUP BY levelvalue
//					)b on b.levelvalue = l.levelvalue
//					where b.count = 1 `+con+` and l.levelvalue > ?
//					order by levelvalue`
//	var ul []UserLevel
//	err := orm.Eloquent.Raw(sql,levelvalue,levelvalue).Scan(&ul).Error
//	return ul,err
//}
//func (u *UserLevel)GetReferrerLevel(ids string,accumulative float64)([]UserLevel,error){
//	sql := `select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
//					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
//					where user_id in (`+ids+`) and is_del = 0 and status = 0 and u.accumulative > ?
//					GROUP BY user_id order by levelvalue`
//	var ul []UserLevel
//	err := orm.Eloquent.Raw(sql,accumulative).Scan(&ul).Error
//	//mp := make(map[string]UserLevel)
//	//for i:=0;i<len(ul);i++ {
//	//	mp[ul[i].UserId] = ul[i]
//	//}
//	//同级排序，直接上级在前面
//	refs := strings.Split(ids,",")
//	mp := make(map[string]int)
//	for i := len(refs)-1;i>=0;i--{
//		mp[refs[i]] = i
//	}
//	var res []UserLevel
//	//一级只取一个
//	for i:=0;i<len(ul);i++{
//		pre := ul[i]
//		for j:=i+1;j<len(ul);j++{
//			if pre.Levelvalue < ul[j].Levelvalue{
//				//res = append(res, ul[i])
//				break
//			}else{
//				i++
//			}
//			if mp[pre.UserId]<mp[ul[j].UserId]{
//				pre = ul[j]
//			}
//		}
//		res = append(res, pre)
//	}
//	return res,err
//}
//func (e *UserLevel) GetAccumulative(id string) (float64,error) {
//	var ref Referrer
//	ref = ref.GetReferrer(id)
//	als := strings.Trim(ref.Referrals,",")
//	if als!=""{
//		als += ","
//	}
//	als += id
//	sql := `select sum(amount)accumulative from investment where userid in(`+als+`) and status `
//	type Acc struct {
//		Accumulative float64
//	}
//	var acc Acc
//	if err := orm.Eloquent.Raw(sql).Scan(&acc).Error;err!=nil{
//		return 0,nil
//	}
//	con := acc.Accumulative
//	//获取手动设置
//	ul,err := e.GetUserSetLevel(id)
//	if err!=nil{
//		return 0, err
//	}
//	if con < ul{
//		con = ul
//	}
//	return con,nil
//}
func (e *UserLevel) GetUserSetLevel(id string) (float64, error) {
	sql := `select set_level from sys_user WHERE user_id = ` + id
	type Set struct {
		SetLevel float64
	}
	var set Set
	if err := orm.Eloquent.Raw(sql).Scan(&set).Error; err != nil {
		return 0, nil
	}
	return set.SetLevel, nil
}
func (e *UserLevel) GetAccumulative(id string) (float64, error) {
	sql := `select accumulative from sys_user WHERE user_id = ` + id
	type Acc struct {
		Accumulative float64
	}
	var acc Acc
	if err := orm.Eloquent.Raw(sql).Scan(&acc).Error; err != nil {
		return 0, nil
	}
	return acc.Accumulative, nil
}
func (e *UserLevel) GetMaxAccumulative(id string) (float64, error) {
	// sql := `select if(accumulative>set_level,accumulative,set_level) accumulative from sys_user WHERE user_id = ` + id
	sql := `select if(lifts<1,if(accumulative>set_level,accumulative,set_level),set_level) accumulative from sys_user WHERE user_id = ` + id
	type Acc struct {
		Accumulative float64
	}
	var acc Acc
	if err := orm.Eloquent.Raw(sql).Scan(&acc).Error; err != nil {
		return 0, nil
	}
	return acc.Accumulative, nil
}
func (u *UserLevel) GetReferrerLevel(ids string, accumulative float64) ([]UserLevel, error) {
	refs := strings.Split(ids, ",")
	if len(refs) == 0 {
		return nil, nil
	}
	var ul []UserLevel
	for _, id := range refs {
		acc, err := u.GetMaxAccumulative(id)
		if err != nil {
			return nil, err
		}
		astr := utils.Float64ToString(acc)
		//sql := `select u.user_id,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
		//			left join (SELECT * from user_level)l on if(u.set_level>`+acc+`,u.set_level,`+acc+`) >= l.levelvalue
		//			where user_id = `+id+` and is_del = 0 and status = 0 and if(u.set_level>`+acc+`,u.set_level,`+acc+`) > ?
		//			GROUP BY user_id order by levelvalue`
		sql := `select *,` + astr + ` accumulative from user_level 
					where levelvalue =
					(
						SELECT max(levelvalue) from user_level
						where  ` + astr + ` >= levelvalue
					)`
		var u UserLevel
		if err = orm.Eloquent.Raw(sql).Scan(&u).Error; err != nil {
			return nil, err
		}
		if u.Levelvalue > accumulative {
			u.UserId = id
			ul = append(ul, u)
		}
	}
	sort.Slice(ul, func(i, j int) bool {
		return ul[i].Levelvalue < ul[j].Levelvalue
	})

	//同级排序，直接上级在前面
	mp := make(map[string]int)
	for i := len(refs) - 1; i >= 0; i-- {
		mp[refs[i]] = i
	}
	var res []UserLevel
	ulmap := make(map[float64]bool)
	//一级只取一个
	for i := 0; i < len(ul); i++ {
		pre := ul[i]
		for j := i + 1; j < len(ul); j++ {
			if pre.Levelvalue < ul[j].Levelvalue {
				//res = append(res, pre)
				break
			}
			if mp[pre.UserId] < mp[ul[j].UserId] {
				pre = ul[j]
			}
		}
		//res = append(res, pre)
		if _, ok := ulmap[pre.Levelvalue]; !ok {
			res = append(res, pre)
			ulmap[pre.Levelvalue] = true
		}

	}
	return res, nil
}

//获取用户等级
//func (u *UserLevel)GetSetUserByUserid(userid string)(UserLevel,error){
//	sql := `select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent, max(l.percentreality) percentreality
//			from sys_user u
//			left join (select * from user_level ) l on  u.accumulative >= l.levelvalue
//			where u.user_id = `+userid+`
//			GROUP BY user_id
//			order by levelvalue `
//	var ul UserLevel
//	err := orm.Eloquent.Raw(sql).Scan(&ul).Error
//	return ul,err
//}

func (u *UserLevel) GetSetUserByUserid(userid string) (UserLevel, error) {
	acc, err := u.GetMaxAccumulative(userid)
	if err != nil {
		return UserLevel{}, err
	}

	astr := utils.Float64ToString(acc)
	sql := `select *,` + astr + ` accumulative from user_level 
         where levelvalue =
         (
            SELECT max(levelvalue) from user_level
            where  ` + astr + `  >= levelvalue
         )`
	var ul UserLevel
	err = orm.Eloquent.Raw(sql).Scan(&ul).Error
	return ul, err
}

//获取设置等级
func (u *UserLevel) GetSetUserLevel(levelvalue float64) ([]UserLevel, error) {
	sql := `select * from user_level where levelvalue>=? order by levelvalue `
	var ul []UserLevel
	err := orm.Eloquent.Raw(sql, levelvalue).Scan(&ul).Error
	return ul, err
}

type FindLevel struct {
	SetLevel int64 `gorm:"column:set_level" json:"set_level"`
}

//设置用户等级
func (this *UserLevel) EditUserVipLevel(userID string, levelID string) (err error) {
	sql := `select if(accumulative>set_level,accumulative,set_level) set_level from sys_user where user_id = ` + userID
	setLevelVip := make([]FindLevel, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&setLevelVip).Error; err != nil {
		return
	}
	// SetLevelStr := strconv.FormatInt(setLevelVip[0].SetLevel, 10)

	sql1 := `select levelvalue from user_level where id = ` + levelID
	findVip := make([]UserLevelConfig, 0)
	if err = orm.Eloquent.Raw(sql1).Scan(&findVip).Error; err != nil {
		return
	}
	if len(findVip) <= 0 {
		err = errors.New("没找到vip配置")
		return
	}
	levelValueStr := strconv.FormatInt(findVip[0].LevelValue, 10)
	nums := "0"
	if findVip[0].LevelValue < setLevelVip[0].SetLevel {
		nums = "1"
	}
	sql2 := `update sys_user set set_level = ` + levelValueStr + `,lifts=` + nums + `,update_time = now() where user_id = ` + userID
	if err = orm.Eloquent.Exec(sql2).Error; err != nil {
		return
	}

	return
}

func (this *UserLevel) GetVipLevelList(userID string) (ret interface{}, err error) {
	//accumulative, err := this.getUserSetLevel(userID)
	//if err != nil {
	//	return
	//}
	accumulative, err := this.GetMaxAccumulative(userID)
	if err != nil {
		return
	}
	//if accumulative < accumulative1{
	//	accumulative = accumulative1
	//}

	sql := `select CONCAT('V',(@rowNO := @rowNo+1)) AS vipLevel,a.* from  user_level a,(select @rowNO :=0) b ORDER BY levelvalue asc`
	vips := make([]vipLevelData, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&vips).Error; err != nil {
		return
	}

	if len(vips) <= 0 {
		err = errors.New("没找到配置数据")
		return
	}
	//
	if accumulative <= 0 {
		retData := map[string]interface{}{
			"curLevelID": vips[0].ID,
			"curLevel":   vips[0].VipLevel,
			"vipList":    vips,
		}
		ret = retData
		return
	}

	curLevel := ""
	curLevelID := int64(0)
	allLen := len(vips)
	for i, m := range vips {
		if i+1 == allLen {
			curLevel = m.VipLevel
			curLevelID = m.ID
			break
		}
		next := vips[i+1]
		if accumulative >= m.LevelValue && accumulative < next.LevelValue {
			curLevel = m.VipLevel
			curLevelID = m.ID
			break
		}
	}

	retData := map[string]interface{}{
		"curLevelID": curLevelID,
		"curLevel":   curLevel,
		"vipList":    vips,
	}
	ret = retData
	return
}

func (this *UserLevel) getUserSetLevel(userID string) (val float64, err error) {
	sql := `select set_level from sys_user where user_id=` + userID

	type findData struct {
		SetLevel float64 `gorm:"column:set_level"`
	}
	finds := make([]findData, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}

	if len(finds) != 1 {
		err = errors.New("用户表数据量不对")
		return
	}

	val = finds[0].SetLevel
	return
}
