package testing

import (
	"fmt"
	"log"
	"testing"
	"xAdmin/config"
	"xAdmin/models"
)

func Test_showPerformance(t *testing.T) {
	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}

	list := []int64{77, 87, 88, 92}
	for _, id := range list {
		total := getUserPerformance(id)
		fmt.Println("id:", id, "业绩:", total)
	}
}
