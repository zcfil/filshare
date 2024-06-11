package service

import (
	"fmt"
	"strconv"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

// AverageFilConfig 利润计算
type AverageFilConfig struct {
	ID         int64     `gorm:"column:id"`
	Height     int64     `gorm:"column:height"`
	CreateTime time.Time `gorm:"column:create_time"`
	AverageFil float64   `gorm:"column:average_fil"`
}

type OrderSettlementFlow struct {
	Mid            int       `gorm:"column:mid" json:"mid,omitempty"`
	SettlementFlow string    `gorm:"column:settlement_flow" json:"settlement_flow"` //结算水号
	CrateTimes     time.Time `gorm:"column:crate_times" json:"crate_times"`         //创建时间
	UpdateTimes    time.Time `gorm:"column:update_times" json:"update_times"`       //操作更新时间
	TotalIncome    string    `gorm:"column:total_income" json:"total_income"`       //线性存储奖励75%
	ReleaseDays    int64     `gorm:"column:release_days" json:"release_days"`       //余额释放周期
	ActiveStatus   int64     `gorm:"column:active_status" json:"active_status"`     //此订单释放转态  0 释放中	1 释放完成
	OrderMid       string    `gorm:"column:order_mid" json:"order_mid"`             //订单号
	OneIncome      string    `gorm:"column:one_income" json:"one_income"`           //单天释放锁仓币
	DaysExecuted   int64     `gorm:"column:days_executed" json:"days_executed"`     //已经执行周期天数
	Locked         int64     `gorm:"column:locked" json:"locked"`                   //已经执行周期天数
	EndTime        time.Time `gorm:"column:end_time" json:"end_time"`
	UsersID        string    `gorm:"column:users_id" json:"users_id"` // 用户userid
}

type OrderSettlement struct {
	Mid                  int    `gorm:"column:mid" json:"mid,omitempty"`
	OrderFlow            int64  `gorm:"column:order_flow" json:"order_flow"`                         // 结算流水 ID
	DataTimes            string `gorm:"column:data_times" json:"data_times"`                         // 初次结算时间
	OrderMid             int64  `gorm:"column:order_mid" json:"order_mid"`                           // 订单ID
	OneTIncome           string `gorm:"column:one_t_income" json:"one_t_income"`                     // 单T收益
	StatusId             int    `gorm:"column:status_id" json:"status_id"`                           // 订单状态 是否生效
	TotalIncome          string `gorm:"column:total_income" json:"total_income"`                     // 总收益
	UpdateTimes          string `gorm:"column:update_times" json:"update_times"`                     // 更新日期
	InstantStorageEWarDs string `gorm:"column:instant_storage_ewards" json:"instant_storage_ewards"` // 即时存储奖励25%
	LinearStorageReward  string `gorm:"column:linear_storage_reward" json:"linear_storage_reward"`   // 线性存储奖励75%
	Damaged              string `gorm:"column:damaged" json:"damaged"`                               // 0.98折损
	BillingCycle         int    `gorm:"column:billing_cycle" json:"billing_cycle"`                   // 结算周期 180
	OneIncome            string `gorm:"column:one_income" json:"one_income"`
}

type Invetment struct {
	DateTimes string `gorm:"column:update_time" json:"update_time"` //更新日期
}

type CustomerBalance struct {
	LockedBalance string `gorm:"column:locked_balance" json:"locked_balance"`
	WalletBalance string `gorm:"column:wallet_balance" json:"wallet_balance"`
}

type OrderValid struct {
	Id         string  `gorm:"column:order_id" json:"order_id"`
	Storage    float64 `gorm:"column:storage" json:"storage"`
	Day        int     `gorm:"column:total_day" json:"total_day"`
	Days       int     `gorm:"column:days" json:"days"`               //释放锁仓周期  23/3/8修改
	CustomerID string  `gorm:"column:customer_id" json:"customer_id"` //客户id
}

type OrderSettEMtFlow1 struct {
	SettlementFlowId string    `gorm:"column:settlement_flow" json:"settlement_flow"`
	EndTime          time.Time `gorm:"column:end_time" json:"end_time"`
}

type OrderId struct {
	Id         string    `gorm:"column:id" json:"id"`
	EndTime    time.Time `gorm:"column:end_time" json:"end_time"`
	CustomerId string    `gorm:"column:customer_id" json:"customer_id"`
}

type FlowOrder struct {
	SettlementFlow int64 `gorm:"column:settlement_flow" json:"settlement_flow"` //结算水号
	UpdateTimes    int64 `gorm:"column:update_times" json:"update_times"`       // 更新日期
}

type Migrate struct {
	OrderMid      string    `gorm:"column:order_mid" json:"order_mid"`
	Time          time.Time `gorm:"column:time" json:"time"`
	LockedBalance string    `gorm:"column:locked_balance" json:"locked_balance"`
	TotalDays     int       `gorm:"column:total_days" json:"total_days"`
	StartTime     time.Time `gorm:"column:start_time" json:"start_time"`
	EndTIme       time.Time `gorm:"column:end_time" json:"end_time"`
	OneReward     string    `gorm:"column:one_reward" json:"one_reward"`
	UserId        string    `gorm:"column:user_id" json:"user_id"`
	PlusOne       string    `gorm:"column:plus_one" json:"plus_one"`
	Status        string    `gorm:"column:status" json:"status"`
	ActiveStatus  string    `gorm:"column:active_status" json:"active_status"`
	IsDel         string    `gorm:"column:is_del" json:"is_del"`
	Active        string    `gorm:"column:active" json:"active"`
	AdminId       string    `gorm:"column:admin_id" json:"admin_id"`
}

type CheckTimes struct {
	CreateTime   string `gorm:"column:create_time" json:"create_time"`
	ActiveStatus string `gorm:"column:active_status" json:"active_status"`
}

//configKey 配置key
const (
	FROM_ADDRESS             = "from_address"             //转出钱包配置
	COLLECTION_ADDRESS       = "collection_address"       //归集钱包
	CUSTOMER_INCOME_FLOATING = "customer_income_floating" // 客户收益浮动比例
)

type Finance struct {
	// 代码
	ID          int64   `json:"id" gorm:"column:configId;primary_key"` //编码
	ConfigKey   string  `json:"name" gorm:"column:configKey;"`         //参数名称 //参数键名ConfigKey string `json:"configKey" gorm:"column:configKey"`
	ConfigName  string  `json:"title" gorm:"column:configName"`        //变量标题  //参数名称ConfigName string `json:"Name" gorm:"column:name;primary_key"`
	ConfigValue string  `json:"value" gorm:"column:configValue"`       //参数变量值 	//参数键值 //ConfigValue string `json:"configValue" gorm:"column:configValue"`
	Balance     float64 `json:"balance"`
}

func (f *Finance) FinanceConfigList() ([]Finance, error) {
	sql := `select * from sys_config where is_del = 0`
	var fi []Finance
	err := orm.Eloquent.Raw(sql).Scan(&fi).Error
	return fi, err
}

func (f *Finance) UpdateConfigById(param map[string]string) ([]Finance, error) {
	sql := `update sys_config set configName=:title,configvalue=:value where configId = :id`
	sql = utils.SqlReplaceParames(sql, param)
	var fi []Finance
	err := orm.Eloquent.Raw(sql).Scan(&fi).Error
	return fi, err
}

func (f *Finance) GetFromWallet() (err error) {
	sql := `select * from sys_config where configKey = "%s"`
	sql = fmt.Sprintf(sql, FROM_ADDRESS)
	if err = orm.Eloquent.Raw(sql).Scan(f).Error; err != nil {
		return
	}
	return
}

func (f *Finance) GetCustomerIncomeFloating() (ratio int64, err error) {
	sql := `select * from sys_config where configKey = "%s"`
	sql = fmt.Sprintf(sql, CUSTOMER_INCOME_FLOATING)
	if err = orm.Eloquent.Raw(sql).Scan(f).Error; err != nil {
		return
	}

	ratio, err = strconv.ParseInt(f.ConfigValue, 10, 64)
	if err != nil {
		return
	}
	return
}

// LockBalanceRecord 锁定的历史记录
type LockBalanceRecord struct {
	UpdateDay int32     `json:"updateDay"`
	Arr       []float64 `json:"arr"`
}
