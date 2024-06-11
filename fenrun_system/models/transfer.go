package models

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	orm "xAdmin/database"
	log "xAdmin/logrus"
	"xAdmin/pkg/lotus"
	"xAdmin/utils"
)

type Transfer struct {
	TransferId    int64     `json:"transfer_id"`
	Cid           string    `json:"cid"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	Amount        float64   `json:"amount"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Status        int       `json:"status"`
	ServiceCharge float64   `json:"service_charge"`
}
type TransferView struct {
	Transfer
	Name string `json:"name"`
}

func (t *Transfer) GetTransferList() ([]Transfer, error) {
	var ts []Transfer
	sql := ` select * from transfer where status = 0 and create_time < SUBDATE(now() ,INTERVAL 2 MINUTE ) `
	err := orm.Eloquent.Raw(sql).Scan(&ts).Error
	return ts, err
}

func (t *Transfer) SetServiceCharge(amount float64) error {
	sql := ` update transfer set status = 1,service_charge = ?,update_time = now() where transfer_id = ? `
	return orm.Eloquent.Exec(sql, amount, t.TransferId).Error
}

func (t *Transfer) Insert(param map[string]string) error {
	sql := " insert into transfer (cid,amount,`from`,`to`,user_id)value(:cid,:amount,:from,:to,:user_id) "
	sql = utils.SqlReplaceParames(sql, param)
	return orm.Eloquent.Exec(sql).Error
}

func (t *Transfer) Send(ctx context.Context, userid, from, addrTo string, amountk float64) (err error) {
	session := orm.Eloquent.Begin() // 保存到数据库  开始数据sql事务操作
	defer func() {
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
	}()
	wallet, err := t.getBabalce(userid, session)
	if err != nil {
		return
	}
	var amount float64
	if wallet <= amountk {
		amount = wallet
	} else {
		amount = amountk
	}
	msg, err := lotus.Send(ctx, from, addrTo, amount)
	if err != nil {
		return
	}
	param := make(map[string]string)
	param["cid"] = msg.Message.Cid().String()
	param["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)
	param["from"] = msg.Message.From.String()
	param["to"] = msg.Message.To.String()
	param["user_id"] = userid

	errl := t.updatewallet(userid, amount, session)
	if errl != nil {
		log.Error("更新用户钱包算信息错误, errl ：", errl)
	}
	errInsert := t.Insert(param)
	if errInsert != nil {
		log.Error("插入结算信息错误, errInsert：", errInsert)
	}
	return
}

func (t *Transfer) GetList(param map[string]string) ([]TransferView, error) {
	var ts []TransferView
	//原来的数据库语句
	//sql := ` select t.*,c.name from transfer t
	//		left join customer c on t.user_id = c.id
	//		where :start <= t.create_time and t.create_time <= :end`
	sql := `SELECT a.* FROM(SELECT t.*, c.name FROM transfer t LEFT JOIN customer c on t.user_id=c.id  
			WHERE t.status=1 and  t.create_time <=:end  and :start <= t.create_time GROUP BY t.transfer_id)a 
			LEFT JOIN customer d  on a.transfer_id=d.id `
	sql = utils.SqlReplaceParames(sql, param)
	param["total"] = GetTotalCount(sql)
	//param["sort"] = "t.create_time" //原来sql语句
	param["sort"] = "a.update_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	err := orm.Eloquent.Raw(sql).Scan(&ts).Error
	return ts, err
}

func (t *Transfer) updatewallet(userid string, amount float64, session *gorm.DB) (err error) {
	wallet, err := t.getBabalce(userid, session)
	if err != nil {
		return
	}
	wallet_count := wallet - amount
	total_wallet := strconv.FormatFloat(wallet_count, 'f', -1, 64)
	err = t.balance_Update(userid, total_wallet, session, utils.TimeHMS())
	if err != nil {
		return
	}
	return
}

type Customer_Balance struct {
	Wallet_Balance float64 `gorm:"column:wallet_balance" json:"wallet_balance"`
}

func (t *Transfer) getBabalce(userid string, session *gorm.DB) (wallet float64, err error) { // 获取用户余额  锁仓余额
	ret := make([]Customer_Balance, 0)
	sqlx := `SELECT  wallet_balance  FROM customer where id = "%s"`
	fmt.Println(userid)
	sqlx = fmt.Sprintf(sqlx, userid)
	if err = session.Debug().Raw(sqlx).Scan(&ret).Error; err != nil {
		log.Error("查询钱包错误 : ", err.Error())
		return
	}

	wallet = ret[0].Wallet_Balance
	return
}

func (t *Transfer) balance_Update(userid string, wallet string, session *gorm.DB, times string) (err error) {
	// 更新用户钱包 余额
	sqlx := `UPDATE  customer SET  wallet_balance="%s"  ,update_time="%s" WHERE id="%s"`
	sqlx = fmt.Sprintf(sqlx, wallet, times, userid)
	if err = session.Debug().Exec(sqlx).Error; err != nil {
		log.Error("更新用余额 锁仓余额  失败 : ", err.Error())
		return
	}
	return
}
