package models

import (
	"log"
	"xAdmin/config"
	orm "xAdmin/database"
)

func RunInit() {
	var conf Config
	c, err := conf.GetConfig("customerratio")
	if err != nil || (c == Config{}) {
		sql := `insert into sys_config(name,value)value('customerratio',?),('salesmanratio',?)`
		if err = orm.Eloquent.Exec(sql, config.ApplicationConfig.CustomerRatio, config.ApplicationConfig.SalesmanRatio).Error; err != nil {
			log.Fatal(err)
		}
	}
}
