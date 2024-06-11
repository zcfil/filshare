package service

import (
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/utils"
)

func NewOrderCheck() *OrderSettlement {
	OrderChek := new(OrderSettlement)
	OrderChek.startSettleLoop()
	return OrderChek
}

func (this *OrderSettlement) startSettleLoop() { // 每天凌晨零时1分开始生成前提用户收益清单
	spec := "01, 59, 23, *, *, *" // 每天0 点 01 分
	c := utils.CronNew()
	if err := c.AddFunc(spec, this.start); err != nil {
		log.Error("Add Order_check func error:", err.Error())
		return
	}
	c.Start()
}

func (this *OrderSettlement) start() {
	if !this.checkTime() { // check active status true or false
		log.Info("正常结算开启")
		if err := this.settlement(); err != nil {
			log.Error("每天凌晨开启结算 : ", err.Error())
			return
		}
		log.Info("结算线性释放结束")
		// check  Insert active true
		if err := this.checkTimeInsert(); err != nil {
			return
		}
	}
}

func (this *OrderSettlement) checkTime() bool {
	sqlx := `SELECT  date_format(create_time,'%Y-%m-%d') AS  create_time,active_status  from  times_status WHERE create_time='` + utils.TimeHMS() + `'`
	var check []CheckTimes
	if err := orm.Eloquent.Debug().Raw(sqlx).Scan(&check).Error; err != nil {
		return false
	}
	if len(check) > 0 { // 0
		log.Info("不为空")
		times := check[0].CreateTime
		log.Info(times)
		if times == "" {
			log.Info("不为空3")
			return true
		}
		TimeInt64 := utils.TImeInt64(times)
		TimesInt641 := utils.TimeDayInt64()
		if TimeInt64 == TimesInt641 {
			log.Info("不为空4")
			return true
		}
	}
	return false
}

func (this *OrderSettlement) checkTimeInsert() error {
	sqlx := `insert into times_status (create_time, active_status) VALUES ('` + utils.TimeHMS() + `',1)`
	if err := orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return err
	}
	return nil
}
