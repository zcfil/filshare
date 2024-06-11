package apis

import (
	"encoding/json"
	"errors"
	"net/http"
	"xAdmin/common"
	"xAdmin/define"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/redisClient"
	"xAdmin/result"
	"xAdmin/utils"

	"github.com/gin-gonic/gin"
)

func GetcustomerList(c *gin.Context) {
	var u models.Customer
	var err error
	param := make(map[string]string)

	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")

	result, err := u.CustomerList(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

func GetUserList(c *gin.Context) {
	var data models.Customer

	param := make(map[string]string)
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")

	result, err := data.GetSysList(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}
func CustomerAdd(c *gin.Context) {
	var data models.Customer
	var d models.Config
	param := make(map[string]string)
	param["name"] = c.Request.FormValue("name")
	param["phone"] = c.Request.FormValue("phone")
	param["user_id"] = utils.GetUserIdStr(c)
	param["identity"] = c.Request.FormValue("identity")
	param["wallet"] = c.Request.FormValue("wallet")
	param["password"] = c.Request.FormValue("password")

	d, _ = d.GetConfig("customerratio")
	isExist, err := data.IsExistByPhone(param["phone"])
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	if isExist {
		c.JSON(http.StatusOK, result.Failstr("该手机号码已注册"))
		c.Abort()
		return
	}
	var res models.Response
	err = data.CustomerAdd(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func CustomerEdit(c *gin.Context) {
	var data models.Customer
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	param["name"] = c.Request.FormValue("name")
	param["phone"] = c.Request.FormValue("phone")
	param["user_id"] = utils.GetUserIdStr(c)
	param["identity"] = c.Request.FormValue("identity")
	param["wallet"] = c.Request.FormValue("wallet")
	param["password"] = c.Request.FormValue("password")
	var res models.Response
	// var cus models.Customer
	// cus, _ = cus.NewCustomer(param["id"])
	// if cus.Userid != param["user_id"] && utils.GetRolekey(c) != "admin" {
	// 	res.Msg = "无权限操作该客户"
	// 	c.JSON(http.StatusOK, res.ReturnError(400))
	// 	return
	// }
	err := data.CustomerEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func CustomerDelete(c *gin.Context) {
	var data models.Customer
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	var res models.Response
	sql1 := `select * from investment where is_del <> 1 and customer_id=` + param["id"] + `'`
	param["total"] = models.GetTotalCount(sql1)
	if param["total"] > "0" {
		res.Msg = "该用户有订单不能删除！"
		c.JSON(http.StatusOK, res.ReturnError(201))
		return
	}
	err := data.CustomerDelete(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func CustomerProfitEdit(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["customerid"] = c.Request.FormValue("customerid")
	param["profit"] = c.Request.FormValue("profit")
	var res models.Response
	err := data.CustomerProfitEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetCustomerByid(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	id := c.Request.FormValue("customerid")
	var res models.Response
	re, err := data.NewCustomer(id)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}
	res.Data = re
	c.JSON(http.StatusOK, res.ReturnOK())
}
func CustomerLogin(c *gin.Context) {
	phone := c.Request.FormValue("phone")
	if phone == "" {
		c.JSON(http.StatusOK, result.Failstr("phone不能为空"))
		c.Abort()
		return
	}

	pwd := c.Request.FormValue("password")
	if pwd == "" {
		c.JSON(http.StatusOK, result.Failstr("password不能为空"))
		c.Abort()
		return
	}

	customer := models.NewCustomer()
	if err := customer.GetCustomerByPhone(phone); err != nil {
		c.JSON(http.StatusOK, result.Failstr("未注册的手机号码"))
		c.Abort()
		return
	}

	//cryptoPwd := utils.EncodePassword(phone, pwd)
	//if customer.Password != cryptoPwd {
	//	c.JSON(http.StatusOK, result.Failstr("密码错误"))
	//	c.Abort()
	//	return
	//}

	if pwd != customer.Password {
		c.JSON(http.StatusOK, result.Failstr("密码错误"))
		c.Abort()
		return
	}
	// 创建token
	token := common.GenCustomerToken()
	tokenKey := common.GenCustomerTokenKey(token)

	data, err := json.Marshal(customer)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	if err = redisClient.RedisClient.Set(tokenKey, data, define.TOKEN_EXPIRATION_TIME).Err(); err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(token))
}
func CustomerInvestmentList(c *gin.Context) {
	token := c.GetHeader(define.TOKEN_STR)
	if token == "" {
		c.JSON(http.StatusOK, result.LoginTimeout(errors.New("token 不能为空")))
		c.Abort()
		return
	}

	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空"))
		c.Abort()
		return
	}

	param := make(map[string]string)
	param["pageSize"] = pageSize
	param["pageIndex"] = pageIndex

	customer := models.NewCustomer()
	if err := customer.GetCustomerByToken(token); err != nil {
		c.JSON(http.StatusOK, result.LoginTimeout(err))
		c.Abort()
		return
	}

	list, err := customer.CustomerInvestmentList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	res := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, res)
}
func CustomerChangePassword(c *gin.Context) {
	oldPassword := c.Request.FormValue("oldPassword")
	if oldPassword == "" {
		c.JSON(http.StatusOK, result.Failstr("旧密码不能为空"))
		c.Abort()
		return
	}

	newPassword := c.Request.FormValue("newPassword")
	if newPassword == "" {
		c.JSON(http.StatusOK, result.Failstr("新密码不能为空"))
		c.Abort()
		return
	}

	token := c.GetHeader(define.TOKEN_STR)
	if token == "" {
		c.JSON(http.StatusOK, result.Failstr("token不能为空"))
		c.Abort()
		return
	}

	customer := models.NewCustomer()
	if err := customer.GetCustomerByToken(token); err != nil {
		c.JSON(http.StatusOK, result.LoginTimeout(err))
		c.Abort()
		return
	}

	if customer.Password != oldPassword {
		c.JSON(http.StatusOK, result.Failstr("旧密码错误"))
		c.Abort()
		return
	}

	if err := customer.ChangePassword(newPassword, token); err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}
func CustomerTransferList(c *gin.Context) {
	token := c.GetHeader(define.TOKEN_STR)
	if token == "" {
		c.JSON(http.StatusOK, result.LoginTimeout(errors.New("token 不能为空")))
		c.Abort()
		return
	}

	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空"))
		c.Abort()
		return
	}

	param := make(map[string]string)
	param["pageSize"] = pageSize
	param["pageIndex"] = pageIndex

	customer := models.NewCustomer()
	if err := customer.GetCustomerByToken(token); err != nil {
		c.JSON(http.StatusOK, result.LoginTimeout(err))
		c.Abort()
		return
	}

	list, err := customer.GetTransferList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	pageData := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, pageData)
}

func CustomerSettlementList(c *gin.Context) {
	token := c.GetHeader(define.TOKEN_STR)
	if token == "" {
		c.JSON(http.StatusOK, result.LoginTimeout(errors.New("token 不能为空")))
		c.Abort()
		return
	}

	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空"))
		c.Abort()
		return
	}

	param := make(map[string]string)
	param["pageSize"] = pageSize
	param["pageIndex"] = pageIndex

	customer := models.NewCustomer()
	if err := customer.GetCustomerByToken(token); err != nil {
		c.JSON(http.StatusOK, result.LoginTimeout(err))
		c.Abort()
		return
	}
	list, err := customer.SettlementList(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	pageData := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, pageData)
}
func CustomerHomepage(c *gin.Context) {
	token := c.GetHeader(define.TOKEN_STR)
	if token == "" {
		c.JSON(http.StatusOK, result.LoginTimeout(errors.New("token 不能为空")))
		c.Abort()
		return
	}
	customer := models.NewCustomer()
	if err := customer.GetCustomerByToken(token); err != nil {
		c.JSON(http.StatusOK, result.LoginTimeout(err))
		c.Abort()
		return
	}

	ret, err := customer.Homepage()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(ret))
}

func CustomerSettlementListMigrate(c *gin.Context) {
	token := c.GetHeader(define.TOKEN_STR)
	if token == "" {
		c.JSON(http.StatusOK, result.LoginTimeout(errors.New("token 不能为空")))
		c.Abort()
		return
	}

	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空"))
		c.Abort()
		return
	}

	param := make(map[string]string)
	param["pageSize"] = pageSize
	param["pageIndex"] = pageIndex

	customer := models.NewCustomer()
	if err := customer.GetCustomerByToken(token); err != nil {
		c.JSON(http.StatusOK, result.LoginTimeout(err))
		c.Abort()
		return
	}
	migrage := models.NewMigrate()
	list, err := migrage.BillingImpressions(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	pageData := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, pageData)
}
