package testing

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"xAdmin/config"
	orm "xAdmin/database"
	"xAdmin/models"
)

/*
	本测试单元是用来生成sys_user表里面用户的业绩的
*/

func Test_CreatePerformance(t *testing.T) {
	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}

	sql := `select user_id from sys_user`
	type findUser struct {
		UserID int64 `gorm:"column:user_id" json:"user_id"`
	}
	findUsers := make([]findUser, 0)
	if err := orm.Eloquent.Raw(sql).Scan(&findUsers).Error; err != nil {
		return
	}

	for _, m := range findUsers {
		total := getUserPerformance(m.UserID)
		if total == 0 {
			continue
		}

		sql2 := `update sys_user set accumulative=%d where user_id = %d`
		sql2 = fmt.Sprintf(sql2, total, m.UserID)
		if err := orm.Eloquent.Exec(sql2).Error; err != nil {
			fmt.Println("err:", err)
			continue
		}
	}
}

func getUserPerformance(userID int64) (total int64) {
	// 先获取下级的业绩， 在加上自己的
	// 先取出下级列表
	type findReferrers struct {
		Referrals string `gorm:"column:referrals"` // 下级id列表
	}
	sql := `select referrals from referrer where userid=` + strconv.FormatInt(userID, 10)
	var findRef findReferrers
	if err := orm.Eloquent.Raw(sql).Scan(&findRef).Error; err != nil {
		return
	}

	ref := findRef.Referrals
	sql1 := `select sum(i.amount) as total from investment as i where userid in(` + strconv.FormatInt(userID, 10)
	if ref != "" {
		if ref[0] == ',' {
			ref = ref[1:]
		}
		sql1 += `,`
		sql1 += ref
	}
	sql1 += `) and i.status <> 1`
	type totalPerformance struct {
		Total float64 `gorm:"column:total"`
	}

	var findTotal totalPerformance
	if err := orm.Eloquent.Raw(sql1).Scan(&findTotal).Error; err != nil {
		return
	}
	total = int64(findTotal.Total)
	return
}
