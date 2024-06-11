package service

//
//import (
//	"errors"
//	"strings"
//	"time"
//	orm "xAdmin/database"
//	log "xAdmin/logrus"
//	"xAdmin/utils"
//)
//
//type TransferList struct {
//	ID           int       `gorm:"column:id" json:"id,omitempty"`             //ID
//	Name         string    `gorm:"column:name" json:"name"`                   //ID
//	UserId       string    `gorm:"column:userid" json:"user_id"`              // 用户id
//	ToBalance    string    `gorm:"column:to_balance" json:"to_balance"`       // 用户收益
//	LockRelease  string    `gorm:"column:lock_release" json:"lock_release"`   // 用户收益
//	Amount       string    `gorm:"column:amount" json:"amount"`               // 用户收益
//	Active       string    `gorm:"column:active" json:"active"`               // 总监确认
//	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`     // 创建时间
//	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`     // 创建时间
//	OrderMid     string    `gorm:"column:order_mid" json:"order_mid"`         // 订单id
//	TotalCompany string    `gorm:"column:total_company" json:"total_company"` // 公司收益
//}
//
//func (this *TransferList) start() {
//	if err := this.transfer(); err != nil {
//		log.Error("生成清单失败")
//		return
//	}
//}
//
//func (this *TransferList) transfer() (err error) {
//	session := orm.Eloquent.Begin()
//	defer func() {
//		if err != nil {
//			session.Rollback()
//			return
//		}
//		session.Commit()
//	}()
//
//	times := utils.TimeHMS() // 获得时间
//	date, err := this.getweekes(times)
//	if len(date) <= 0 {
//		return
//	}
//	if err != nil {
//		return
//	}
//	if err = this.getdata(date); err != nil {
//		return
//	}
//	return
//
//}
//
//func (this *TransferList) getdata(date string) (error) {
//	if len(date) > 0 {
//		startDate, endDate, err := this.getStartEndTime(date)
//		if err !=nil{
//			return err
//		}
//		sqlx := `update settle_log s   SET enabled=0 where UNIX_TIMESTAMP(s.time) >= UNIX_TIMESTAMP('` + startDate + `')
//      and UNIX_TIMESTAMP(s.time) <= UNIX_TIMESTAMP('` + endDate + `') and is_transfer <> 1`
//		if err = orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
//			return err
//		}
//		return err
//	}
//
//	return nil
//}
//
//func (this *TransferList) getStartEndTime(date string) (start, end string, err error) {
//	list := strings.Split(date, "--")
//	if len(list) != 2 {
//		err = errors.New("日期格式不正确")
//		log.Warning("date:", date)
//		return
//	}
//
//	start = list[0] + " 00:00:00"
//	end = list[1] + " 23:59:59"
//	return
//}
//
//func (this *TransferList) getweekes(times string) (ret string, err error) {
//	sqlx := `select if(DATE_FORMAT(time, '%u')='00','52',DATE_FORMAT(time, '%u')) weeks,  sum(to_customer_balance + customer_lock_release) amount, CONCAT(
//               DATE_FORMAT(  subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 1),  '%Y-%m-%d'), '--',
//                DATE_FORMAT(  subdate(max(time), if(date_format(max(time), '%w') = 0, 7, date_format(max(time), '%w')) - 7), '%Y-%m-%d'))   date
//              from (select s.*  from settle_log s  inner join customer c on c.id = s.customer_id) a
//              where is_transfer <> 1  AND YEARWEEK(date_format(time, '%Y-%m-%d'), 1) = YEARWEEK('` + times + `', 1)   GROUP BY weeks  `
//	type dates struct { //'`+ utils.TimeHMS() +`'
//		Date string `gorm:"column:date" json:"date"`
//	}
//	rets := make([]dates, 0)
//	log.Info(sqlx)
//	if err = orm.Eloquent.Debug().Raw(sqlx).Scan(&rets).Error; err != nil {
//		return
//	}
//	var ret1 string
//	if len(rets) > 0 {
//		ret1 = rets[0].Date
//		return ret1, nil
//	}
//	return ret1, nil
//}
//
//func (this *TransferList) transferTest(times string) (err error) {
//	session := orm.Eloquent.Begin()
//	defer func() {
//		if err != nil {
//			session.Rollback()
//			return
//		}
//		session.Commit()
//	}()
//
//	date, err := this.getweekes(times)
//	if err != nil {
//		return
//	}
//
//	if err = this.getdata(date); err != nil {
//		return
//	}
//
//	log.Info("for data ok" + "ok")
//	return
//
//}
