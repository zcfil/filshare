package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path"
	"strconv"
	"strings"
	"xAdmin/config"
	log "xAdmin/logrus"
	"xAdmin/models"
	"xAdmin/result"
	"xAdmin/utils"
)

//批量导入
func UploadAddress(c *gin.Context) {
	var a models.Addr
	file, err := c.FormFile("file")
	if err != nil {

	}
	fmt.Println(file)
	f, _ := file.Open()

	var res models.Response

	if err = a.UploadAddress(f, file.Size); err != nil {
		res.Msg = err.Error()
		res.Code = 400
		c.JSON(http.StatusOK, res)
	} else {
		res.Msg = "导入成功"
		c.JSON(http.StatusOK, res.ReturnOK())
	}
}

func UploadFile(ctx *gin.Context) {
	username := ctx.Request.FormValue("userid")
	adminname := ctx.Request.FormValue("adminid")
	// file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusNotFound, result.Fail(err))
		ctx.Abort()
		return
	}

	customer := models.NewCustomer()
	user_id, err := customer.UserId(username)
	if err != nil {
		log.Error("查询不到记录 ：", err.Error())
	}
	log.Info("User UUID ", user_id)
	var sysuserid models.SysUser
	admin_id, err := sysuserid.AminId(adminname)
	if err != nil {
		log.Error("查询管理级别用户的 查询不到异常：", err.Error())
		return
	}
	log.Info("admin UUID ", admin_id)

	UserTotal, err := models.MaxLimit(user_id)
	if err != nil {
		log.Error("查询管理级别用户的 查询不到异常", err.Error())
		return
	}
	if UserTotal >= config.UpdateFile.Total {
		ctx.JSON(http.StatusOK, gin.H{
			"total": "20",
			"msg":   "超出最大限制",
		})
		ctx.Abort()
		return
	}
	size := strconv.FormatInt(file.Size, 10)
	client := config.OSSConfig.Client
	bucket, err := client.Bucket(config.OSSConfig.BucketName)
	if err != nil {
		log.Error("oss bucket  found :", err.Error())
		return
	}
	filepath := path.Join("./"+config.UpdateFile.UploadPath, file.Filename)
	err = ctx.SaveUploadedFile(file, filepath)
	if err != nil {
		log.Error(logrus.Fields{"err": err.Error()}, "controller - admin - upload")
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	log.Info("file Filename ", file.Filename)
	//filename := path.Join("./"+config.UpdateFile.UploadPath, file.Filename)
	uploadFile := fmt.Sprintf("kdsystem/%v/%v", utils.TimeMonth(), file.Filename)
	err = bucket.PutObjectFromFile(uploadFile, filepath)

	if err != nil {
		log.Error(" put upload  :", err.Error())
	}
	filename := strings.Split(file.Filename, ".")
	filename1 := filename[0]
	fmt.Println("sss", file.Filename)
	urlk := fmt.Sprintf("%v/%v", config.UpdateFile.URL, uploadFile)
	fmt.Println(config.UpdateFile.URL, "ssss")
	var sysUser models.Addr
	err = sysUser.UploadFile(urlk, user_id, admin_id, size, filename1)

	if err != nil {
		log.Error("记录插入错误 ", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"url":  urlk,
	})
	log.Info("地址 ： ", urlk)
}

//  展示 用户上传文件信息详情

func GetUploadFile(ctx *gin.Context) {
	var sysUser models.Addr
	username := ctx.Request.FormValue("username")
	data, err := sysUser.GetUploadFile(username)
	if err != nil {
		ctx.JSON(http.StatusForbidden, result.Fail(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result.Ok(data))

}

// 用户删除上传文件

func DeleteFile(ctx *gin.Context) {
	var sysUser models.Addr
	filename := ctx.Request.FormValue("filename")
	err := sysUser.DelStatusFile(filename)
	if err != nil {
		ctx.JSON(http.StatusForbidden, result.Fail(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result.Code())
}
