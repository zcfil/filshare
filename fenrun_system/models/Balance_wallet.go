package models

import (
	"context"
	"fmt"
	"strconv"
	time2 "time"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/pkg/lotus"
)

type Default_Config struct {
	ID    int64  `json:"id" gorm:"column:configId;primary_key"` //编码
	Name  string `json:"name" gorm:"column:configKey;"`         //参数名称 //参数键名ConfigKey string `json:"configKey" gorm:"column:configKey"`
	Title string `json:"title" gorm:"column:configName"`        //变量标题  //参数名称ConfigName string `json:"Name" gorm:"column:name;primary_key"`

	//Type       string `json:"type" gorm:"column:type"`             //变量类型:string,text,int,bool,array,datetime,date,file
	Remark string `json:"remark" gorm:"column:remark"`    //变量描述 //Remark string `json:"remark" gorm:"column:remark"` //备注
	Group  string `json:"group" gorm:"column:configType"` //变量分组
	//Content    string `json:"content" gorm:"column:content"`       //变量字典数据
	//Rule       string `json:"rule" gorm:"column:rule"`             //验证规则
	//Extend     string `json:"extend" gorm:"column:extend"`         //扩展属性
	Value string `json:"value" gorm:"column:configValue"` //参数变量值 	//参数键值 //ConfigValue string `json:"configValue" gorm:"column:configValue"`
	IsDel string `json:"isDel" gorm:"column:is_del"`      //是否删除 0 正常使用，1 已删除
	//ConfigType string `json:"configType" gorm:"column:configType"` //变量类型
	//Params     string `json:"params" gorm:"column:params"`
	CreateBy   string     `json:"createBy" gorm:"column:create_by"`
	CreateTime time2.Time `json:"createTime" gorm:"column:create_time"`
	UpdateBy   string     `json:"updateBy" gorm:"column:update_by"`
	UpdateTime time2.Time `json:"updateTime" gorm:"column:update_time"`
	//DataScope  string `json:"dataScope" gorm:"_"`
}

func (e *Default_Config) GetConfig() (param map[string]string, err error) {
	param = make(map[string]string)
	sql := `SELECT  * from  default_config WHERE configKey="%s"`
	sql = fmt.Sprintf(sql, FROM_ADDRESS)
	if err = orm.Eloquent.Debug().Raw(sql).Scan(&e).Error; err != nil {
		log.Error("查询错误+++++++++", err.Error())
		return
	}
	address := e.Value
	total, err := lotus.Balance(context.Background(), address)
	if err != nil {
		return
	}
	ret := strconv.FormatFloat(total, 'f', -1, 64)
	param["address"] = address
	param["total"] = ret
	return param, err
}
