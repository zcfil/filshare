package models

import (
	"fmt"
	orm "xAdmin/database"
)

//configKey 配置key
const (
	FROM_ADDRESS             = "from_address"             //转出钱包配置
	COLLECTION_ADDRESS       = "collection_address"       //归集钱包
	CUSTOMER_INCOME_FLOATING = "customer_income_floating" // 客户收益浮动比例
)

type Default_Finance struct {
	// 代码
	ID          int64   `json:"id" gorm:"column:configId;primary_key"` //编码
	ConfigKey   string  `json:"name" gorm:"column:configKey;"`         //参数名称 //参数键名ConfigKey string `json:"configKey" gorm:"column:configKey"`
	ConfigName  string  `json:"title" gorm:"column:configName"`        //变量标题  //参数名称ConfigName string `json:"Name" gorm:"column:name;primary_key"`
	ConfigValue string  `json:"value" gorm:"column:configValue"`       //参数变量值 	//参数键值 //ConfigValue string `json:"configValue" gorm:"column:configValue"`
	Balance     float64 `json:"balance"`
}

func (f *Default_Finance) GetFromWallet() (err error) {
	sql := `select * from default_config where configKey = "%s"`
	sql = fmt.Sprintf(sql, FROM_ADDRESS)
	if err = orm.Eloquent.Raw(sql).Scan(f).Error; err != nil {
		return
	}
	return
}
