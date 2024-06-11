package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xAdmin/models"
	"xAdmin/result"
	"xAdmin/utils"
)

//func GetWeekList1(c *gin.Context) {
//	pageSize := c.Request.FormValue("pageSize")
//	if pageSize == "" {
//		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空！"))
//		c.Abort()
//		return
//	}
//	pageIndex := c.Request.FormValue("pageIndex")
//	if pageIndex == "" {
//		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空！"))
//		c.Abort()
//		return
//	}
//
//	param := map[string]string{
//		"pageSize":  pageSize,
//		"pageIndex": pageIndex,
//	}
//	settle := models.NewSettle()
//	list, err := settle.GetWeekList1(param)
//	if err != nil {
//		c.JSON(http.StatusOK, result.Fail(err))
//		c.Abort()
//		return
//	}
//	data := utils.NewPageData(param, list)
//	c.JSON(http.StatusOK, data)
//}

func GetWeekCustomerList1(c *gin.Context) {
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
	list, err := settle.GetWeekCustomerList1(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	data := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, data)
}

func GetWeekCustomerListExport(c *gin.Context) {
	date := c.Request.FormValue("date")
	if date == "" {
		c.JSON(http.StatusOK, result.Failstr("date 不能为空！"))
		c.Abort()
		return
	}

	param := map[string]string{
		"pageSize":  "9999",
		"pageIndex": "1",
		"date":      date,
	}

	settle := models.NewSettle()
	var res models.Response
	result, err := settle.ExportWeekCustomerList(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetWeekCustomerInvestmentList1(c *gin.Context) {
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
	list, err := settle.GetWeekCustomerInvestmentList1(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	data := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, data)
}

func GetSettleList1(c *gin.Context) {
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
	ret, err := settle.SettleList1(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	data := utils.NewPageData(param, ret)
	c.JSON(http.StatusOK, data)
}
