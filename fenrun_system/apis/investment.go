package apis

import (
	"net/http"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"

	"github.com/gin-gonic/gin"
)

func InvestmentList(c *gin.Context) {
	var u models.Investment
	var err error
	param := make(map[string]string)
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["user_id"] = utils.GetUserIdStr(c)
	result, err := u.InvestmentList(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

func InvestmentAdd(c *gin.Context) {
	var data models.Investment
	var d models.Config
	param := make(map[string]string)
	d, _ = d.GetConfig("customerratio")
	param["name"] = c.Request.FormValue("name")
	param["phone"] = c.Request.FormValue("phone")
	param["storage"] = c.Request.FormValue("storage")
	param["remark"] = c.Request.FormValue("remark")
	param["customer_id"] = c.Request.FormValue("customer_id")
	param["user_id"] = utils.GetUserIdStr(c)
	param["dayk"] = c.Request.FormValue("totalDay")
	param["days"] = c.Request.FormValue("days") //23/3/8 新增释放周期字段
	var res models.Response
	// var cus models.Customer
	// cus, _ = cus.NewCustomer(param["customer_id"])
	// if cus.Userid != param["user_id"] && utils.GetRolekey(c) != "admin" {
	// 	res.Msg = "无权限操作该客户"
	// 	c.JSON(http.StatusOK, res.ReturnError(400))
	// 	return
	// }
	err := data.InvestmentAdd(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	//var settement service.Order_Settlement
	//if err = settement.Settlement(order );err !=nil{
	//	return
	//}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentEdit(c *gin.Context) {
	var data models.Investment
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	param["customer_id"] = c.Request.FormValue("customer_id")
	param["name"] = c.Request.FormValue("name")
	param["storage"] = c.Request.FormValue("storage")
	param["remark"] = c.Request.FormValue("remark")
	param["user_id"] = utils.GetUserIdStr(c)
	param["time"] = c.Request.FormValue("time")
	param["dayk"] = c.Request.FormValue("totalDay")
	param["days"] = c.Request.FormValue("days") //23/3/13  释放周期字段
	//day := param["dayk"]
	//DaySk, _ := strconv.Atoi(day)
	//if DaySk <= 180 {
	//	log.Info("天数 ",day)
	//	errs   := xerrors.New("天数最小要求180天")
	//	c.JSON(http.StatusOK, result.Fail(errs))
	//	c.Abort()
	//	return
	//}

	var res models.Response
	// var cus models.Customer
	// cus, _ = cus.NewCustomer(param["customer_id"])
	// if cus.Userid != param["user_id"] && utils.GetRolekey(c) != "admin" {
	// 	res.Msg = "无权限操作该客户"
	// 	c.JSON(http.StatusOK, res.ReturnError(400))
	// 	return
	// }
	err := data.InvestmentEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func InvestmentDelete(c *gin.Context) {
	var data models.Investment
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	var res models.Response
	err := data.InvestmentDelete(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func InvestmentBreak(c *gin.Context) {
	var data models.Investment
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")

	var res models.Response
	// var cus models.Customer
	// cus, _ = cus.NewCustomer(param["customerid"])
	// if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
	// 	res.Msg = "无权限操作该客户"
	// 	c.JSON(http.StatusOK, res.ReturnError(400))
	// 	return
	// }
	err := data.InvestmentBreak(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentRevoke(c *gin.Context) {
	var data models.Investment
	param := make(map[string]string)
	param["investmentid"] = c.Request.FormValue("id")
	param["customerid"] = c.Request.FormValue("customerid")
	param["userid"] = utils.GetUserIdStr(c)

	var res models.Response
	var cus models.Customer
	cus, _ = cus.NewCustomer(param["customerid"])
	if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
		res.Msg = "无权限操作该客户"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	err := data.InvestmentRevoke(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetInvestmentByid(c *gin.Context) {
	var data models.Investment
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)

	id := c.Request.FormValue("id")

	var res models.Response
	re, err := data.InvestmentById(id)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}
	res.Data = re

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentExport(c *gin.Context) {
	var u models.Investment
	param := make(map[string]string)
	var res models.Response

	result, err := u.ExportInvestment(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentImport(c *gin.Context) {
	var res models.Response

	file, err := c.FormFile("file")
	if err != nil {
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	var u models.Investment
	f, _ := file.Open()
	err = u.InvestmentImport(f, file.Size)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	if file.Size == 0 {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Msg = "导入成功！"

	c.JSON(http.StatusOK, res.ReturnOK())
}

func EmptyInvestment(c *gin.Context) {
	var res models.Response
	var u models.Investment
	err := u.DeleteInvestment()
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	res.Msg = "删除成功！"

	c.JSON(http.StatusOK, res.ReturnOK())
}
