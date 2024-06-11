package models

import (
	"fmt"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/utils"
)

type Sys_Config struct {
	Value       string `gorm:"column:value" json:"value"`
	Create_Time int64  `gorm:"column:create_time" json:"create_time"`
	Allies_Name int64  `gorm:"column:allies_name" json:"allies_name"`
	Is_Del      int64  `gorm:"column:is_del" json:"is_del"`
}

type Sex struct {
	Fox        DataCrawler
	periodTime int64
}

func (this Sex) GetIncome24() (err error) {
	filfox, h1, err := this.Fox.WriteData()
	if err != nil {
		return
	}
	sqlx := `INSERT INTO  filearnings (height,average_fil,create_time) VALUES ("%s",%s,"%s")`
	sqlx = fmt.Sprintf(sqlx, h1, filfox, utils.TimeHMS())
	err = orm.Eloquent.Debug().Exec(sqlx).Error
	return
}

func (this Sex) gets() {
	err := this.GetIncome24()
	if err != nil {
		log.Error("插入数据失败 ", err.Error())
		return
	}
}
func NewDetails() *Sex {
	details := new(Sex)
	details.startSettleLoop()
	return details
}

func (this Sex) startSettleLoop() {
	spec := "01, 58, 23, *, *, *" // 每天0 点 01 分
	c := utils.CronNew()
	if err := c.AddFunc(spec, this.gets); err != nil {
		log.Error("Add settle func error:", err.Error())
		return
	}
	c.Start()
}
