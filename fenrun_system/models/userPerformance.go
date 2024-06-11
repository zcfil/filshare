package models

import (
	"errors"
	"strconv"
	"time"
	orm "xAdmin/database"
)

type UserPerformance struct {
	userID   int64
	NickName string `json:"nick_name"` // 用户昵称
	//TotalPerformance  int64			`json:"total_performance"`	 // 本人总业绩总业绩(真实业绩)
	TotalPerformance float64 `json:"total_performance"` // 本人总业绩总业绩(真实业绩)
	TodayPerformance float64 `json:"today_performance"` // 当日业绩(真实业绩)
	VipLevel         string  `json:"vip_level"`         // vip等级
	VipTitle         string  `json:"vip_title"`         // vip称号

	vipLevelPerformance int64 // 计算vip
	//vipLevelPerformance float64
}

func NewUserPerformance(userID int64) *UserPerformance {
	return &UserPerformance{
		userID: userID,
	}
}

func (this *UserPerformance) GetTotal() (err error) {
	// 先取出下级列表
	type findReferrers struct {
		Referrals string `gorm:"column:referrals"` // 下级id列表
	}
	sql := `select referrals from referrer where userid=` + strconv.FormatInt(this.userID, 10)
	var findRef findReferrers
	if err = orm.Eloquent.Raw(sql).Scan(&findRef).Error; err != nil {
		return
	}

	ref := findRef.Referrals
	sql1 := `select sum(i.amount) as total from investment as i where userid in(` + strconv.FormatInt(this.userID, 10)
	if ref != "" {
		if ref[0] == ',' {
			ref = ref[1:]
		}
		sql1 += `,`
		sql1 += ref
	}
	sql1 += `) and i.status = 0 and ifnull(manually_end,0)=0`
	type totalPerformance struct {
		//Total  int64		`gorm:"column:total"`
		Total float64 `gorm:"column:total"`
	}

	var findTotal totalPerformance
	if err = orm.Eloquent.Raw(sql1).Scan(&findTotal).Error; err != nil {
		return
	}
	sql2 := `select nick_name from sys_user where user_id=` + strconv.FormatInt(this.userID, 10)
	type nickName struct {
		NickName string `gorm:"column:nick_name"`
	}
	var findName nickName
	if err = orm.Eloquent.Raw(sql2).Scan(&findName).Error; err != nil {
		return
	}

	this.TotalPerformance = findTotal.Total
	this.NickName = findName.NickName
	return
}

func (this *UserPerformance) GetVipTotal() (err error) {
	sql := `select if(lifts<1,if(accumulative>set_level,accumulative,set_level),set_level) accumulative from sys_user WHERE user_id = ` + strconv.FormatInt(this.userID, 10)
	type setLevel struct {
		Accumulative int64 `gorm:"column:accumulative"`
	}
	var find setLevel
	if err = orm.Eloquent.Raw(sql).Scan(&find).Error; err != nil {
		return
	}
	this.vipLevelPerformance = find.Accumulative
	return
	// sql := `select ifnull(accumulative, 0) accumulative,ifnull(set_level, 0) set_level from sys_user where user_id=` + strconv.FormatInt(this.userID, 10)
	// type setLevel struct {
	// 	SetLevel     int64 `gorm:"column:set_level"`
	// 	Accumulative int64 `gorm:"column:accumulative"`
	// }

	// var find setLevel
	// if err = orm.Eloquent.Raw(sql).Scan(&find).Error; err != nil {
	// 	return
	// }

	// this.vipLevelPerformance = find.SetLevel
	// if find.Accumulative > find.SetLevel {
	// 	this.vipLevelPerformance = find.Accumulative
	// }
	// return
}

// GetVipLevel 先获取总业绩再获取vip等级
func (this *UserPerformance) GetVipLevel() (err error) {
	sql := `select * from user_level`
	list := make([]UserLevelConfig, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&list).Error; err != nil {
		return
	}

	if len(list) <= 0 {
		err = errors.New("符合查找条件的数量不对")
		return
	}

	if this.vipLevelPerformance <= 0 {
		tempVipLevel := int64(1)
		this.VipLevel = "V" + strconv.FormatInt(tempVipLevel, 10)
		this.setVipTitle(tempVipLevel)
		return
	}

	// 根据业务员的业绩获取业务员的vip等级
	tempVipLevel := int64(0)
	allLen := len(list)
	for i, m := range list {
		vl := int64(i + 1)
		if i+1 == allLen {
			tempVipLevel = vl
			break
		}
		next := list[i+1]
		if this.vipLevelPerformance >= m.LevelValue && this.vipLevelPerformance < next.LevelValue {
			tempVipLevel = vl
			break
		}
	}
	this.VipLevel = "V" + strconv.FormatInt(tempVipLevel, 10)
	this.setVipTitle(tempVipLevel)
	return
}

func (this *UserPerformance) GetToday() (err error) {
	referrals, err1 := this.getUserReferrals()
	if err1 != nil {
		err = err1
		return
	}

	l := len(referrals)
	sql := `select sum(i.amount) as total from investment as i where ( userid in (` + strconv.FormatInt(this.userID, 10)

	if l > 0 {
		sql += `,`
		sql += referrals
	}
	sql += `)`
	sql += `) and ` + ` create_time >= "`
	today := time.Now().Format("2006-01-02") + ` 00:00:00"`
	sql += today
	sql += ` and i.status = 0 and ifnull(manually_end,0)=0`

	type totalMsg struct {
		Total float64 `gorm:"column:total" json:"total"`
	}
	var find totalMsg
	if err = orm.Eloquent.Raw(sql).Scan(&find).Error; err != nil {
		return
	}

	this.TodayPerformance = find.Total
	return
}

// 获取推荐人列表
func (this *UserPerformance) getUserReferrals() (referrals string, err error) {
	// 先获取下级推荐人列表
	sql := `select referrals from referrer where userid = ` + strconv.FormatInt(this.userID, 10)
	type referral struct {
		Referrals string `gorm:"column:referrals" json:"referrals"`
	}
	find := make([]referral, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&find).Error; err != nil {
		return
	}

	l := len(find)
	if l <= 0 {
		return
	}
	if l != 1 {
		err = errors.New("查找推荐列表数量错误")
		return
	}

	referrals = find[0].Referrals
	if len(referrals) == 0 {
		return
	}
	if referrals[0] == ',' {
		referrals = referrals[1:]
	}

	return
}

// setVipTitle 设置vip称号
func (this *UserPerformance) setVipTitle(vipLevel int64) {
	switch vipLevel {
	case 0:
		fallthrough
	case 1:
		this.VipTitle = "业务员"
		return
	case 2:
		this.VipTitle = "银冠"
		return
	case 3:
		this.VipTitle = "金冠"
		return
	case 4:
		this.VipTitle = "皇冠"
		return
	default:
		this.VipTitle = "皇冠"
		return
	}
}
