package testing

import (
	"fmt"
	"testing"
	_ "xAdmin/config"
	"xAdmin/models"
)

func TestAdmin(t *testing.T) {
	var user models.InReviewUser
	list, count, _ := user.GetUserPassingList(1, 10, "common", "1")
	fmt.Println("list:", list)
	fmt.Println("count:", count)
}

func TestAdd(t *testing.T) {
	var sysUser models.SysUser

	submitterID := "120"

	sysUser.Username = "you1234"
	sysUser.NickName = "you1234"
	sysUser.Phone = "12345678901"
	sysUser.Password = "111111"
	// sysUser.Referrer = submitterID

	err := sysUser.SubmitNewUser(submitterID)
	fmt.Println("Submit new user err:", err)
}

func TestAllow(t *testing.T) {
	userID := "127"
	var sysUser models.SysUser
	err := sysUser.AllowNewPass(userID)
	fmt.Println("err:", err)
}

func TestNotAllow(t *testing.T) {
	userID := "122"
	var sysUser models.SysUser
	err := sysUser.NotAllowUserPass(userID)
	fmt.Println("err:", err)
}

func TestAuditList(t *testing.T) {
	audit := models.NewUserAudit()
	list, total, _ := audit.GetAuditList(1, 10, "1", "common")
	fmt.Println("list:", list)
	fmt.Println("total:", total)
}
