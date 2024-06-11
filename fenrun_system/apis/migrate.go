package apis

import (
	"errors"
	"github.com/gin-gonic/gin"

	"net/http"

	"xAdmin/models"
	"xAdmin/result"
	"xAdmin/utils"
)

func AddMigrate(ctx *gin.Context) {
	param := make(map[string]string)
	param["balance"] = ctx.Request.FormValue("balance")
	param["customer_id"] = ctx.Request.FormValue("customer_id")
	param["days"] = ctx.Request.FormValue("totalDay")
	param["remark"] = ctx.Request.FormValue("remark")
	param["user_id"] = utils.GetUserIdStr(ctx)
	var customer models.Customer
	if err := customer.UpdateSetIsKid(param["customer_id"]); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}

	var migrate models.Migrate
	if err := migrate.MigrateInsert(param); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result.Code())
}

func GetMigrate(ctx *gin.Context) {
	param := make(map[string]string)
	param["keyword"] = ctx.Request.FormValue("keyword")
	param["pageSize"] = ctx.DefaultQuery("pageSize", "10")
	param["pageIndex"] = ctx.DefaultQuery("pageIndex", "1")
	param["user_id"] = utils.GetUserIdStr(ctx)
	var migrate models.Migrate
	data, err := migrate.GetListData(param)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	res := utils.NewPageData(param, data)
	ctx.JSON(http.StatusOK, result.Ok(res))
}

func DeleteMigrate(ctx *gin.Context) {
	orderMid := ctx.Request.FormValue("order_mid")
	var migrate models.Migrate
	err := migrate.RetrieveMigrate(orderMid)
	if err != nil {
		err = errors.New("订单不存在")
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	err = migrate.DeleteMigreate(orderMid)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result.Code())
}
func EditMigrate(ctx *gin.Context) {
	param := make(map[string]string)
	param["orderMid"] = ctx.Request.FormValue("order_mid")
	param["balance"] = ctx.Request.FormValue("balance")
	param["customer_id"] = ctx.Request.FormValue("customer_id")
	param["remark"] = ctx.Request.FormValue("remark")
	param["Time"] = ctx.Request.FormValue("time")
	param["days"] = ctx.Request.FormValue("totalDay")

	var migrate models.Migrate
	err := migrate.RetrieveMigrate(param["orderMid"])
	if err != nil {
		err = errors.New("订单不存在")
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}

	err = migrate.EditMigrate(param)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result.Code())
}

func BreakMigreate(ctx *gin.Context) {
	orderMid := ctx.Request.FormValue("order_mid")
	var migrate models.Migrate
	err := migrate.RetrieveMigrate(orderMid)
	if err != nil {
		err = errors.New("订单不存在")
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	err = migrate.BreakMigreate(orderMid)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result.Code())
}

func BillingImpressions(ctx *gin.Context) {
	pageSize := ctx.Request.FormValue("pageSize")
	if pageSize == "" {
		ctx.JSON(http.StatusOK, result.Failstr("pageSize 不能为空！"))
		ctx.Abort()
		return
	}
	pageIndex := ctx.Request.FormValue("pageIndex")
	if pageIndex == "" {
		ctx.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空！"))
		ctx.Abort()
		return
	}

	param := map[string]string{
		"pageSize":  pageSize,
		"pageIndex": pageIndex,
	}
	settle := models.NewMigrate()
	ret, err := settle.BillingImpressions(param)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}

	data := utils.NewPageData(param, ret)
	ctx.JSON(http.StatusOK, data)
}
