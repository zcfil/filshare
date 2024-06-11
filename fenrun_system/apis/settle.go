package apis

import (
	"net/http"
	"time"
	"xAdmin/models"
	"xAdmin/result"
	"xAdmin/utils"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func GetWeekList(c *gin.Context) {
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空！"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空！"))
		c.Abort()
		return
	}

	param := map[string]string{
		"pageSize":  pageSize,
		"pageIndex": pageIndex,
	}
	settle := models.NewSettle()
	list, err := settle.GetWeekList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	data := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, data)
}

func GetWeekCustomerList(c *gin.Context) {
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空！"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空！"))
		c.Abort()
		return
	}
	date := c.Request.FormValue("date")
	if date == "" {
		c.JSON(http.StatusOK, result.Failstr("date 不能为空！"))
		c.Abort()
		return
	}

	param := map[string]string{
		"pageSize":  pageSize,
		"pageIndex": pageIndex,
		"date":      date,
	}

	settle := models.NewSettle()
	list, err := settle.GetWeekCustomerList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	data := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, data)
}

func GetWeekCustomerInvestmentList(c *gin.Context) {
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空！"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空！"))
		c.Abort()
		return
	}

	date := c.Request.FormValue("date")
	if date == "" {
		c.JSON(http.StatusOK, result.Failstr("date 不能为空！"))
		c.Abort()
		return
	}
	customerID := c.Request.FormValue("customer_id")
	if customerID == "" {
		c.JSON(http.StatusOK, result.Failstr("customer_id 不能为空！"))
		c.Abort()
		return
	}

	param := map[string]string{
		"pageSize":   pageSize,
		"pageIndex":  pageIndex,
		"date":       date,
		"customerID": customerID,
	}

	settle := models.NewSettle()
	list, err := settle.GetWeekCustomerInvestmentList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	data := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, data)
}

func TransferWeek(c *gin.Context) {
	date := c.Request.FormValue("date")
	if date == "" {
		c.JSON(http.StatusOK, result.Failstr("date 不能为空！"))
		c.Abort()
		return
	}

	settle := models.NewSettle()
	if err := settle.TransferWeek(date); err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
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
	loginlog.Msg = "给日期为 " + date + " 的周列表转账"
	loginlog.Create()

	c.JSON(http.StatusOK, result.Ok(nil))
}

func TransferWeekCustomer(c *gin.Context) {
	date := c.Request.FormValue("date")
	if date == "" {
		c.JSON(http.StatusOK, result.Failstr("date 不能为空！"))
		c.Abort()
		return
	}
	customerID := c.Request.FormValue("customer_id")
	if customerID == "" {
		c.JSON(http.StatusOK, result.Failstr("customer_id 不能为空！"))
		c.Abort()
		return
	}

	param := map[string]string{
		"date":       date,
		"customerID": customerID,
	}
	settle := models.NewSettle()
	if err := settle.TransferWeekCustomer(param); err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
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
	loginlog.Msg = "给日期为 " + date + " 的客户 " + customerID + "转账"
	loginlog.Create()

	c.JSON(http.StatusOK, result.Ok(nil))
}

func GetTransferList(c *gin.Context) {
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	var t models.Transfer
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		c.JSON(http.StatusOK, result.Failstr("时间不能为空"))
		return
	}
	param["start"] += " 00:00:00"
	param["end"] += " 23:59:59"
	list, err := t.GetList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, utils.NewPageData(param, list))
}

func GetSettleList(c *gin.Context) {
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空！"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空！"))
		c.Abort()
		return
	}

	param := map[string]string{
		"pageSize":  pageSize,
		"pageIndex": pageIndex,
	}
	settle := models.NewSettle()
	ret, err := settle.SettleList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	data := utils.NewPageData(param, ret)
	c.JSON(http.StatusOK, data)
}
