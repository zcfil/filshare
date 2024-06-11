package models

import orm "xAdmin/database"

type Referrer struct {
	//Id int64
	Userid    int64
	Referrers string
	Referrals string
}

func (r *Referrer) GetReferrer(userid string) (a Referrer) {
	sql := `select * from referrer where userid = ?`
	orm.Eloquent.Raw(sql, userid).Scan(&a)
	return
}
