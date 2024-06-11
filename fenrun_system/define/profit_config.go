package define

// 分润配置数据表ID 对应数据表profit_config
const (
	PROFIT_CUSTOMER_RATIO = 1 // 收益客户占比
	PROFIT_COMPANY_RATIO  = 2 // 收益公司占比

	PROFIT_CUSTOMER_BALANCE_RATIO = 3 // 收益到余额占比
	PROFIT_CUSTOMER_LOCK_RATIO    = 4 // 收益到锁仓占比

	//CUSTOMER_INCOME_FLOATIND_RATE = 5			// 客户收益浮动比例
)

// 公司收益ID
const COMPANY_INCOME_ID = 1 // 公司收益ID

const LOCK_BALANCE_DAY_COUNT = 180 // 锁定天数

const (
	INSTEAN_REWARD = 0.25
	LINEAR_REWARD  = 0.75
	YMDHMS         = "%Y-%m-%d %H:%i:%s"
)
