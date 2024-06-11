package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xAdmin/models"
	"xAdmin/utils"
)

func CustomerAuditList(c *gin.Context) {
	var u models.CustomerLog
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["status"] = "1"
	param["userid"] = utils.GetUserIdStr(c)
	param["role"] = utils.GetRolekey(c)
	//switch utils.GetRolekey(c) {
	//case "finance":
	//	param["status"] = "1"
	//case "boss":
	//	param["status"] = "1"
	//case "admin":
	//	param["status"] = "1"
	//default:
	//	c.JSON(http.StatusOK, re.ReturnError(401))
	//	return
	//}
	result, err := u.CustomerAuditList(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(401))
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}
func CustomerLog(c *gin.Context) {
	var u models.CustomerLog

	var re models.Response
	id := c.Request.FormValue("id")
	role := utils.GetRolekey(c)
	result, err := u.GetCustomerLog(id, role)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
	}
	re.Data = result
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	c.JSON(http.StatusOK, re.ReturnOK())
}
func CustomerLogList(c *gin.Context) {
	var u models.CustomerLog
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["userid"] = utils.GetUserIdStr(c)
	param["role"] = utils.GetRolekey(c)
	//switch utils.GetRolekey(c) {
	//case "finance":
	//	param["status"] = "1"
	//case "boss":
	//	param["status"] = "1"
	//case "admin":
	//	param["status"] = "1"
	//default:
	//	c.JSON(http.StatusOK, re.ReturnError(401))
	//	return
	//}
	result, err := u.CustomerLogList(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

func CustomerLogAudit(c *gin.Context) {
	var data models.CustomerLog
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	param["status"] = c.Request.FormValue("status")
	var res models.Response
	switch utils.GetRolekey(c) {
	case "finance":
		if param["status"] != "3" {
			param["status"] = "0"
		}
	case "boss":
		if param["status"] != "3" {
			param["status"] = "0"
		}
	case "admin":
		if param["status"] != "3" {
			param["status"] = "0"
		}
	default:
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	//pkg.AssertErr(err, "", 400)

	err := data.CustomerLogAudit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func CustomerLogEdit(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["customerid"] = c.Request.FormValue("id")
	param["namenew"] = c.Request.FormValue("name")
	param["identitynew"] = c.Request.FormValue("identity")
	param["phonenew"] = c.Request.FormValue("phone")
	param["userid"] = c.Request.FormValue("salesmanid")
	param["banknew"] = c.Request.FormValue("bank")
	param["banknumnew"] = c.Request.FormValue("banknum")
	param["sex"] = c.Request.FormValue("sex")
	var res models.Response
	err := data.CustomerEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
