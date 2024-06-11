package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xAdmin/models"
)

func FinanceConfigList(c *gin.Context) {
	var res models.Response
	var f models.Finance
	fs, err := f.FinanceConfigList()
	f.FinanceConfigList()
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = fs
	c.JSON(http.StatusOK, res.ReturnOK())
}

func FinanceConfigEdit(c *gin.Context) {
	var res models.Response
	var f models.Finance
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	param["value"] = c.Request.FormValue("value")
	param["status"] = c.Request.FormValue("status")
	fs, err := f.FinanceConfigEdit(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = fs
	c.JSON(http.StatusOK, res.ReturnOK())
}
