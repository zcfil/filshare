package apis

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"xAdmin/models"
//	"xAdmin/pkg"
//	"xAdmin/result"
//	"xAdmin/service"
//	"xAdmin/utils"
//)
//
//type Transfer struct {
//	models.SysUser
//	service.Transfer_List
//}
//
//func (this *Transfer) GettransferList(ctx *gin.Context) {
//	param := make(map[string]string)
//
//	//usernaem := ctx.Request.FormValue("username")
//	param["pageSize"] = ctx.DefaultQuery("pageSize", "10")
//	param["pageIndex"] = ctx.DefaultQuery("pageIndex", "1")
//	//_, err := this.Getadmindiff(usernaem)
//	//if err != nil {
//	//	ctx.JSON(http.StatusOK, result.Fail(err))
//	//	ctx.Abort()
//	//	return
//	//}
//
//	//if flag {
//	//	ctx.JSON(http.StatusOK, gin.H{
//	//		"code": 200,
//	//		"msg":  "此用户权限不足",
//	//	})
//	//	ctx.Abort()
//	//	return
//	//}
//
//	var transfer service.Transfer_List
//	//resultk, err := transfer.GetTransfer(param)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)
//	total, err := transfer.TotalCount()
//	pkg.AssertErr(err, "抱歉未找到相关信息", -2)
//	data := utils.NewPageDataTotal(param, resultk, total)
//	pkg.AssertErr(err, "抱歉未找到相关信息", -3)
//	ctx.JSON(http.StatusOK, result.Ok(data))
//}
//
//func (this *Transfer) SetTransfer(ctx *gin.Context) {
//	Nubmer := ctx.Request.FormValue("nubmer")
//	transfer_total := ctx.Request.FormValue("total_balance")
//	resultk, err := this.SetTransfe(Nubmer, transfer_total)
//	if err != nil {
//		ctx.JSON(http.StatusOK, result.Fail(err))
//		ctx.Abort()
//		return
//	}
//	ctx.JSON(http.StatusOK, result.Ok(resultk))
//}
//
//func (this *Transfer) Updata_Active(ctx *gin.Context) {
//	username := ctx.Request.FormValue("username")
//	active := ctx.Request.FormValue("active")
//
//	_, err := this.Getadmindiff(username)
//	if err != nil {
//		return
//	}
//
//	flag, err := this.Update_Active(active)
//	if err != nil {
//		ctx.JSON(http.StatusOK, result.Fail(err))
//		ctx.Abort()
//		return
//	}
//	ctx.JSON(http.StatusOK, result.Ok(flag))
//}

//func WeekPostTest(ctx *gin.Context) {
//	var this service.TransferList
//	date := ctx.Request.FormValue("date")
//	log.Info("sssssssss1111111")
//	if err := this.WeekStart(date); err != nil {
//		ctx.JSON(http.StatusOK, result.Fail(err))
//		ctx.Abort()
//		return
//	}
//	ctx.JSON(http.StatusOK, result.Code())
//
//}
