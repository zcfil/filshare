package models

import (
	"time"
	orm "xAdmin/database"
)

type OfficialInfo struct {
	Height             string `gorm:"column:height" json:"height"`                             // 区块高度
	TotalValidHashRate string `gorm:"column:total_valid_hashrate" json:"total_valid_hashrate"` // 全网有效算力
	ActiveProvider     string `gorm:"column:active_provider" json:"active_provider"`           // 活跃提供者
	EveryAward         string `gorm:"column:every_award" json:"every_award"`                   // 每区块奖励
	Income24FT         string `gorm:"column:income_24ft" json:"income_24ft"`                   // 24h平均提供服务收益
	Output24F          string `gorm:"column:output_24f" json:"output_24f"`                     // 近24h产出量
	CurSectorPledge    string `gorm:"column:cur_sector_pledge" json:"cur_sector_pledge"`       // 当前扇区质押量
	FILPledge          string `gorm:"column:fil_pledge" json:"fil_pledge"`                     // FIL 质押量
	MessageCount24     string `gorm:"column:message_count24" json:"message_count24"`           // 24h消息数
	FILTurnover        string `gorm:"column:fil_turnover" json:"fil_turnover"`                 // FIL流通量
	TotalAccount       string `gorm:"column:total_account" json:"total_account"`               // 总账户数
	AverageInterval    string `gorm:"column:average_interval" json:"average_interval"`         // 平均区块间隔
	AverageHeightCount string `gorm:"column:average_height_count" json:"average_height_count"` // 平均高度区块数量
	NewAddHashrateFT   string `gorm:"column:new_add_hashrate_ft" json:"new_add_hashrate_ft"`   // 新增算力成本
	CurBasicHashrate   string `gorm:"column:cur_basic_hashrate" json:"cur_basic_hashrate"`     // 当前基础费率
	FILDestory         string `gorm:"column:fil_destroy" json:"fil_destroy"`                   // fil销毁量
	FILSupply          string `gorm:"column:fil_supply" json:"fil_supply"`                     // fil供给量
	FlowRate           string `gorm:"column:flow_rate" json:"flow_rate"`                       // FIL流通率
	//Price			  			string 		`gorm:"column:price" json:"price"`										// 最新价格
	Date time.Time `gorm:"column:date" json:"date"` // 日期
}

func (this *OfficialInfo) GetNewData() (err error) {
	sql := `select * from filfox_info where id = (select MAX(id) from official_info)`
	if err = orm.Eloquent.Raw(sql).Scan(this).Error; err != nil {
		return
	}
	return
}
