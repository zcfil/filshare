package models

import (
	"strconv"
	"time"
	orm "xAdmin/database"
)

// UserAudit 用户审核记录表
type UserAudit struct {
	UserID   int64     `gorm:"column:user_id" json:"user_id"`     // 用户ID
	NickName string    `gorm:"column:nick_name" json:"nick_name"` // 昵称
	UserName string    `gorm:"column:username" json:"username"`   // 用户名
	Phone    string    `gorm:"column:phone" json:"phone"`         // 电话号码
	Referrer int64     `gorm:"column:Referrer" json:"Referrer"`   // 推荐人
	IsPass   int32     `gorm:"column:is_pass" json:"is_pass"`     // 1 通过 0 拒绝
	PassTime time.Time `gorm:"column:pass_time" json:"pass_time"` // 通过(拒绝)时间
}

type UserAuditRet struct {
	UserAudit
	UserName string `gorm:"column:referrer_name" json:"referrer_name"` // 推荐者用户名
}

func NewUserAudit() *UserAudit {
	ua := new(UserAudit)
	return ua
}

func (this *UserAudit) GetAuditList(pageIndex int64, pageSize int64, curUserID, role string) (ret interface{}, total string, err error) {
	start := (pageIndex - 1) * pageSize
	sql := `select u1.*, u2.username as referrer_name from user_audit u1 left join sys_user u2 on u2.user_id = u1.referrer`

	if _, ok := AUDIT[role]; !ok {
		sql += ` where u1.referrer = ` + curUserID
	}

	total = GetTotalCount(sql)
	sql += ` limit `
	sql += strconv.FormatInt(start, 10)
	sql += `,`
	sql += strconv.FormatInt(pageSize, 10)
	findList := make([]UserAuditRet, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	ret = findList
	return
}
