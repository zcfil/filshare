package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "xAdmin/logrus"
	"xAdmin/models"
	"xAdmin/result"
)

type Default_wallet struct {
	models.Default_Config
}

func GetWallet(ctx *gin.Context) {
	var default_wallet models.Default_Config
	data, err := default_wallet.GetConfig()
	if err != nil {
		log.Error("wallet address Error :", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result.Ok(data))
}
