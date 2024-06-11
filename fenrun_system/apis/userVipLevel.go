package apis

import (
	"net/http"
	"time"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func GetUserVipLevel(c *gin.Context) {
	userID := c.Request.FormValue("userid")
	var conf models.UserLevel
	resultData, err := conf.GetVipLevelList(userID)
	pkg.AssertErr(err, "未找到相关信息", -1)
	var rsp models.Response
	rsp.Data = resultData
	c.JSON(http.StatusOK, rsp.ReturnOK())
}

func EditUserVipLevel(c *gin.Context) {
	userID := c.Request.FormValue("userid")
	levelID := c.Request.FormValue("levelid")
	vipLevel := c.Request.FormValue("vipLevel")
	nick_name := c.Request.FormValue("nick_name")
	//var conf models.UserLevelConfig
	//err := conf.EditUserVipLevel(userID, levelID)
	var conf models.UserLevel
	err := conf.EditUserVipLevel(userID, levelID)
	pkg.AssertErr(err, "跟新vip等级失败", -1)
	var rsp models.Response

	var loginlog models.LoginLog
	ua := user_agent.New(c.Request.UserAgent())
	loginlog.Ipaddr = c.ClientIP()
	location := utils.GetLocation(c.ClientIP())
	loginlog.LoginLocation = location
	loginlog.LoginTime = time.Now()
	loginlog.CreateTime = time.Now()
	loginlog.Status = "0"
	loginlog.IsDel = "0"
	loginlog.Remark = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginlog.Browser = browserName + " " + browserVersion
	loginlog.Os = ua.OS()
	loginlog.Platform = ua.Platform()
	loginlog.UserName = utils.GetUserName(c)
	loginlog.Msg = "调整 " + nick_name + " 的等级为 " + vipLevel
	loginlog.Create()
	c.JSON(http.StatusOK, rsp.ReturnOK())
}
