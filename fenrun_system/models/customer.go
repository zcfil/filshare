package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"xAdmin/common"
	orm "xAdmin/database"
	"xAdmin/define"
	"xAdmin/redisClient"
	"xAdmin/utils"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

//客户表
type Customer struct {
	ID            string    `gorm:"column:id" json:"id"`
	Name          string    `gorm:"column:name" json:"name"` //姓名
	Password      string    `gorm:"password" json:"password"`
	Identity      string    `gorm:"column:identity" json:"identity"` //身份证号
	Phone         string    `gorm:"column:phone" json:"phone"`       //手机号
	Wallet        string    `gorm:"wallet" json:"wallet"`
	Userid        string    `gorm:"column:user_id" json:"user_id"`         //业务员ID
	IsDel         int       `gorm:"column:is_del" json:"is_del"`           //是否删除
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	LockedBalance string    `gorm:"locked_balance" json:"locked_balance"`
	WalletBalance string    `gorm:"column:wallet_balance" json:"wallet_balance"` //业务员ID
	isKid         string    `gorm:"column:is_kid" json:"is_kid"`                 //业务员ID
}
type CustomerView struct {
	Customer
	Flag int `json:"flag"`
}

func NewCustomer() *Customer {
	customer := new(Customer)
	return customer
}

func (e *Customer) NewCustomer(customerid string) (Customer, error) {
	sql2 := `select c.*,u.nick_name,u.username from customer c
					left join sys_user u on c.user_id = u.user_id 
					where id = '` + customerid + "'"
	var c Customer
	if err := orm.Eloquent.Raw(sql2).Scan(&c).Error; err != nil {
		fmt.Errorf("获取客户信息失败：", err)
		return c, err
	}
	return c, nil
}

type Us struct {
	ID    string `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`   //姓名
	Phone string `gorm:"column:phone" json:"phone"` //手机号
}

func (e *Customer) GetSysList(param map[string]string) (interface{}, error) {
	sql := `select id,name,phone from customer where is_del = 0 `
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (name like '%` + keyword + `%') `
	}
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	var u []Us
	sql += utils.LimitAndOrderBy(param)
	err := orm.Eloquent.Raw(sql).Scan(&u).Error

	return u, err
}

func NewCustomer1(param map[string]string) error {
	sql2 := `select * from customer where id = '` + param["customer_id"] + "'"
	var c Customer
	if err := orm.Eloquent.Raw(sql2).Scan(&c).Error; err != nil {
		fmt.Errorf("获取客户信息失败：", err)
		return err
	}
	param["name"] = c.Name
	param["phone"] = c.Phone
	// param["bank"] = c.Bank
	// param["banknum"] = c.Banknum
	param["identity"] = c.Identity
	param["userid"] = c.Userid
	// param["sex"] = c.Sex
	return nil
}

// 客户管理页面数据
func (e *Customer) CustomerList(param map[string]string) (result interface{}, err error) {
	//sql := `select id, name, password, phone,identity, wallet,user_id, create_time,locked_balance,wallet_balance from customer where is_del <> 1`
	sql := `select id, name, password, phone,identity, wallet,user_id, create_time,  round(s.locked_balance,6) AS   locked_balance, round(s.wallet_balance,6)  wallet_balance from customer s  where is_del <> 1`
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (name like '%` + keyword + `%' or phone like '%` + keyword + `%')`
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Customer, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Customer) UpdateSetIsKid(userID string) error {
	sqlx := `update customer set  is_kid=1 Where id='` + userID + `'`
	if err := orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return err
	}
	return nil
}
func (e *Customer) getFilterUsers(userID string) (userIDs string, err error) {
	// 先判断部门，部门没有就取推荐人玩家列表
	userIDs, err = e.getDeptUsers(userID)
	if err == nil {
		return
	}
	return e.getUserReferralsIDs(userID)
}

// 获取下一级玩家列表
func (e *Customer) getNextUserIDs(userID string) (userIDs string, err error) {
	type findReferrer struct {
		UserID int64 `gorm:"column:user_id"`
	}
	sql1 := `select user_id from sys_user where referrer = ` + userID
	var finds []findReferrer
	err = orm.Eloquent.Raw(sql1).Scan(&finds).Error
	if err != nil {
		return
	}

	length := len(finds)
	if length <= 0 {
		err = errors.New("下级列表为空")
		return
	}

	userIDs = userID + ","
	for i, f := range finds {
		userIDs += strconv.FormatInt(f.UserID, 10)
		if i != length-1 {
			userIDs += ","
		}
	}
	return
}

func (e *Customer) getDeptUsers(userID string) (userIDs string, err error) {
	// 是否leader
	sql := `select deptId from sys_dept where leader_id=` + userID
	type findDept struct {
		Deptid int64 `gorm:"column:deptId;primary_key"`
	}
	var depts []findDept
	if err = orm.Eloquent.Raw(sql).Scan(&depts).Error; err != nil {
		return
	}

	if len(depts) <= 0 {
		err = errors.New("no records")
		return
	}

	deptStr := ""
	l := len(depts)
	for i, de := range depts {
		deptStr += strconv.FormatInt(de.Deptid, 10)
		if i != l-1 {
			deptStr += ","
		}
	}
	sql2 := `select user_id from sys_user where dept_id in (` + deptStr + `)`
	var findUsers []SysUserId
	if err = orm.Eloquent.Raw(sql2).Scan(&findUsers).Error; err != nil {
		return
	}

	userIDs = ""
	lu := len(findUsers)
	for i, u := range findUsers {
		userIDs += strconv.FormatInt(u.Id, 10)
		if i != lu-1 {
			userIDs += ","
		}
	}
	return
}

func (e *Customer) getUserReferralsIDs(userID string) (ret string, err error) {
	sql := `select referrals from referrer where userid=` + userID

	var referrers []Referrer
	if err = orm.Eloquent.Raw(sql).Scan(&referrers).Error; err != nil {
		return
	}
	if len(referrers) != 1 {
		err = errors.New("record error")
		return
	}
	ret = referrers[0].Referrals
	if len(ret) <= 0 {
		err = errors.New("no record")
		return
	}
	if len(ret) > 0 && ret[0] == ',' {
		ret = ret[1:]
	}
	ret += "," + userID // 追加上自己的ID
	return
}

//	for key, val := range param {
//		if val != "" && key != "id" {
//			con +=  key + "='" + val.(string) + "',"
//		}
//	}
func (e *Customer) CustomerAdd(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	sql := ` insert into customer(id,name,identity,password,wallet,phone,user_id,locked_balance,wallet_balance,is_kid)value(:id,:name,:identity,:password,:wallet,:phone,:user_id,0,0,0)`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}
	return
}

func (e *Customer) CustomerEdit(param map[string]string) (err error) {
	sql := ` update customer set name="%s",password="%s",phone="%s",identity="%s",wallet="%s", update_time=now() where id = %s`
	sql = fmt.Sprintf(sql, param["name"], param["password"], param["phone"], param["identity"], param["wallet"], param["id"])
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return
	}
	return
}

func (e *Customer) CustomerDelete(param map[string]string) (err error) {
	sql := ` update customer set is_del=1 where id =` + param["id"]
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return
	}
	return
}

func (e *Customer) CustomerProfitEdit(param map[string]string) (err error) {

	sql := ` update customer set profit =:profit where id = :customerid`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}
	return
}

func (e *Customer) GetCustomerByid(param map[string]string) (err error) {

	sql := ` update customer set profit =:profit where id = :customerid`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}
	return
}

// 查询用户id

func (this *Customer) UserId(username string) (ret int, err error) {
	sqlx := `SELECT id  FROM  customer WHERE  name = '` + username + `'`
	type find struct {
		CustomerID int `gorm:"column:id" json:"id"`
	}

	finds := make([]find, 0)
	if err = orm.Eloquent.Raw(sqlx).Scan(&finds).Error; err != nil {
		return
	}
	ret = finds[0].CustomerID
	return ret, err
}

func (this *Customer) GetCustomerByPhone(phone string) (err error) {
	sql := `select * from customer where phone="` + phone + `"`
	if err = orm.Eloquent.Raw(sql).Scan(this).Error; err != nil {
		return
	}
	return
}

func (this *Customer) IsExistByPhone(phone string) (isExist bool, err error) {
	isExist = false
	sql := `select count(id) as count from customer where is_del <> 1 and phone="` + phone + `"`
	type findCount struct {
		Count int64 `gorm:"column:count"`
	}

	findMsg := &findCount{}
	if err = orm.Eloquent.Raw(sql).Scan(findMsg).Error; err != nil {
		return
	}
	if findMsg.Count > 0 {
		isExist = true
		return
	}
	return
}
func (this *Customer) GetCustomerByToken(token string) (err error) {
	tokenKey := common.GenCustomerTokenKey(token)
	data, err := redisClient.RedisClient.Get(tokenKey).Result()
	if err == redis.Nil {
		err = errors.New("登录已过期，请重新登录")
		return
	}
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(data), this); err != nil {
		return
	}
	return
}
func (this *Customer) CustomerInvestmentList(param map[string]string) (ret interface{}, err error) {
	sql := `select * from investment where  is_del<>1 AND customer_id = "%s"`
	sql = fmt.Sprintf(sql, this.ID)
	param["total"] = GetTotalCount(sql)

	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}

	start := (pageIndex - 1) * pageSize
	sql += ` order by create_time desc limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]

	finds := make([]Investment, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}
	ret = finds
	return
}
func (this *Customer) ChangePassword(newPassword, token string) (err error) {
	session := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
	}()
	sql := `update customer set password = "%s" where id = %s`
	sql = fmt.Sprintf(sql, newPassword, this.ID)
	err = session.Exec(sql).Error
	if err != nil {
		return
	}

	this.Password = newPassword
	tokenKey := common.GenCustomerTokenKey(token)
	data, err := json.Marshal(this)
	if err != nil {
		return
	}
	if err = redisClient.RedisClient.Set(tokenKey, data, define.TOKEN_EXPIRATION_TIME).Err(); err != nil {
		return
	}

	return
}
func (this *Customer) GetTransferList(param map[string]string) (ret interface{}, err error) {
	sql := `select transfer_id, amount, create_time from transfer where status = 1 and user_id = %s`
	sql = fmt.Sprintf(sql, this.ID)
	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	param["total"] = GetTotalCount(sql)

	sql += ` order by create_time desc limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]
	type find struct {
		TransferID string    `gorm:"column:transfer_id" json:"transfer_id"`
		Amount     string    `gorm:"column:amount" json:"amount"`
		CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	}
	finds := make([]find, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}
	ret = finds
	return
}

func (this *Customer) SettlementList(param map[string]string) (ret interface{}, err error) {
	sql := `select * from settle_log where customer_id = %s  AND  types_of=2  `
	sql = fmt.Sprintf(sql, this.ID)
	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}

	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	param["total"] = GetTotalCount(sql)
	sql += ` order by time desc limit ` + strconv.FormatInt(start, 10) + `,` + param["pageSize"]
	type settleLog struct {
		ID                  string    `gorm:"column:id" json:"id"`
		InvestmentID        string    `gorm:"column:investment_id" json:"investment_id"`                 // 投资ID
		CustomerIncome      float64   `gorm:"column:customer_income" json:"customer_income"`             // 客户收益
		ToCustomerBalance   float64   `gorm:"column:to_customer_balance" json:"to_customer_balance"`     // 到客户余额
		ToCustomerLock      float64   `gorm:"column:to_customer_lock" json:"to_customer_lock"`           // 到客户锁仓
		CustomerLockRelease float64   `gorm:"column:customer_lock_release" json:"customer_lock_release"` // 客户锁仓释放
		Time                time.Time `gorm:"column:time" json:"time"`                                   // 结算时间
		IsTransfer          int8      `gorm:"column:is_transfer" json:"is_transfer"`                     // 是否已经转账
	}
	logList := make([]settleLog, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&logList).Error; err != nil {
		return
	}
	ret = logList
	return
}

// 客户信息展示内容
type CustomerHomepage struct {
	TotalHashrate float64 `gorm:"column:total_hashrate" json:"total_hashrate"` // 总算力
	TotalIncome   float64 `gorm:"column:total_income" json:"total_income"`     // 总收益
	ReleaseFil    float64 `gorm:"column:release_fil" json:"release_fil"`       // 已释放
	LockFil       float64 `gorm:"column:lock_fil" json:"lock_fil"`             // 锁定
}

func (this *Customer) Homepage() (ret interface{}, err error) {

	// 查找总算力
	type findHashrate struct {
		TotalHashrate float64 `gorm:"column:total_hashrate" json:"total_hashrate"` // 总算力
	}
	sql1 := `select sum(storage) as total_hashrate from investment where customer_id = %s  AND  is_del<>1`
	sql1 = fmt.Sprintf(sql1, this.ID)
	findh := &findHashrate{}
	err = orm.Eloquent.Raw(sql1).Scan(findh).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	// 查找总收益和总释放
	type findIncome struct {
		TotalIncome float64 `gorm:"column:total_income" json:"total_income"` // 总收益
		ReleaseFil  float64 `gorm:"column:release_fil" json:"release_fil"`   // 已释放
	}
	findI := &findIncome{}

	// 查找当前锁定
	type findLock struct {
		LockFil float64 `gorm:"column:locked_balance" json:"locked_balance"` // 锁定
	}
	findL := &findLock{}

	//sql2 := `select sum(to_customer_balance+customer_lock_release) as total_income,
	//	sum(customer_lock_release) as release_fil from settle_log where customer_id = %s`
	sql2 := `select  round(s.wallet_balance,6) total_income  , round(s.wallet_balance,6) release_fil  from customer s  where id = %s`
	sql2 = fmt.Sprintf(sql2, this.ID)
	err = orm.Eloquent.Raw(sql2).Scan(findI).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	sql3 := `select  round(s.locked_balance,6) locked_balance  from customer s  where id = %s`
	sql3 = fmt.Sprintf(sql3, this.ID)
	err = orm.Eloquent.Raw(sql3).Scan(findL).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	// 找不到数据err 置空
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	retMsg := &CustomerHomepage{
		TotalHashrate: findh.TotalHashrate,
		TotalIncome:   findI.TotalIncome,
		ReleaseFil:    findI.ReleaseFil,
		LockFil:       findL.LockFil,
	}
	ret = retMsg
	return
}
