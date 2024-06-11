package models

import (
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Summarydetails struct {
	Summaryid  string    `gorm:"column:summaryid" json:"summaryid"`
	Cname      string    `gorm:"column:cname" json:"cname"`
	Cvalue     string    `gorm:"column:cvalue" json:"cvalue"`
	Profits    float64   `gorm:"column:profits" json:"profits"`
	Percent    float64   `gorm:"column:percent" json:"percent"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (e *Summarydetails) GetSummarydetails(param map[string]string) ([]Summarydetails, error) {
	sql1 := `select * from summarydetails where :start <= create_time and create_time <= :end`
	var once []Summarydetails
	var err error
	sql1 = utils.SqlReplaceParames(sql1, param)
	if err = orm.Eloquent.Raw(sql1).Scan(&once).Error; err != nil {
		return nil, err
	}
	return once, err
}
