package models

import (
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Finance struct {
	// 代码
	ID     int64  `gorm:"column:id" json:"id" `
	Name   string `gorm:"column:name" json:"name" `
	Value  string `gorm:"column:value" json:"value" `
	Locked string `gorm:"column:status" json:"status"`
	IsDel  int    `gorm:"column:is_del" json:"is_del"` //是否删除
}

func (f *Finance) FinanceConfigList() ([]Finance, error) {
	sql := `select * from sys_config where is_del = 0`
	var fi []Finance
	err := orm.Eloquent.Raw(sql).Scan(&fi).Error
	return fi, err
}

func (f *Finance) FinanceConfigEdit(param map[string]string) (interface{}, error) {
	sql := `update sys_config set value=:value ,status=:status  where id = :id`
	sql = utils.SqlReplaceParames(sql, param)
	fi := make([]Finance, 0)
	err := orm.Eloquent.Raw(sql).Scan(&fi).Error
	return fi, err
}

type ConfigSet struct {
	Value  string `gorm:"column:value" json:"value"`
	Locked int    `gorm:"column:status" json:"status"`
}

func (f *ConfigSet) Filfox() (ret []ConfigSet, err error) { // 查询当前结算 单T收益 折扣
	sqlx := `SELECT value,status   FROM sys_config where alies_name='filfox'   `
	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}

	return
}

func (f *ConfigSet) Damaged() (string, error) { // 查询当前结算 单T收益 折扣 0.98%
	sqlx := `SELECT   value ,status FROM sys_config where alies_name='damaged'`
	ret := make([]ConfigSet, 0)
	if err := orm.Eloquent.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		return "", nil
	}
	if len(ret) > 0 {
		value := ret[0].Value
		return value, nil
	}
	return "", nil
}
