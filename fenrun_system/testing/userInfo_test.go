package testing

import (
	"fmt"
	"testing"
	_ "xAdmin/config"
	"xAdmin/models"
	"xAdmin/pkg"
)

func Test_performance(t *testing.T) {
	userID := int64(73)
	per := models.NewUserPerformance(userID)
	err := per.GetTotal()
	pkg.AssertErr(err, "获取总业绩失败", -1)
	err = per.GetToday()
	pkg.AssertErr(err, "获取当天业绩失败", -1)
	err = per.GetVipLevel()
	pkg.AssertErr(err, "获取vip等级失败", -1)
	fmt.Println("基础信息:", per)
	fmt.Println("昵称:", per.NickName)
	fmt.Println("总业绩:", per.TotalPerformance)
	fmt.Println("当日业绩:", per.TodayPerformance)
	fmt.Println("vip等级:", per.VipLevel)
	fmt.Println("vip称号:", per.VipTitle)
}
