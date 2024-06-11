package models

import (
	"fmt"
	"strconv"
	"strings"
	orm "xAdmin/database"
	"xAdmin/utils"

	"github.com/jinzhu/gorm"
)

type InvestmentShareProfit struct {
	ID           string  `gorm:"column:id" json:"id"`
	Userid       string  `gorm:"column:userid" json:"userid"`
	Investmentid string  `gorm:"column:investmentid" json:"investmentid"`
	Profits      float64 `gorm:"column:profits" json:"profits"`
}

func (in *InvestmentShareProfit) AddInvestmentShareProfit(orm1 *gorm.DB, param map[string]string) error {
	var le UserLevel
	var re Referrer
	//获取订单信息
	ins := GetInvestmentById(param["investmentid"])
	re = re.GetReferrer(param["userid"])
	//获取自己等级
	mylevel, err := le.GetSetUserByUserid(param["userid"])
	if err != nil {
		return err
	}
	//获取上级等级
	re.Referrers = strings.Trim(re.Referrers, ",")
	var ulevel []UserLevel

	if re.Referrers != "" {
		ulevel, err = mylevel.GetReferrerLevel(re.Referrers, mylevel.Accumulative)
		if err != nil {
			return err
		}
		//ulevel = append(ulevel, le...)
	}

	//获取设置等级
	var setlevel []UserLevel
	setlevel, err = mylevel.GetSetUserLevel(mylevel.Levelvalue)
	//升级所需
	type upgrade struct {
		need       float64
		levelvalue float64
		percent    float64
	}
	var up []upgrade
	//计算自己的利润
	var isps []InvestmentShareProfit
	var ip InvestmentShareProfit
	// camount := ins.Amount
	ip.Userid = ins.Userid
	ip.Investmentid = param["investmentid"]

	sql1 := `select accumulative from sys_user WHERE user_id = ` + ins.Userid
	type Accnum struct {
		Accumulative float64
	}
	var accnum Accnum
	orm.Eloquent.Raw(sql1).Scan(&accnum)
	type Numactivue struct {
		Tivue float64
	}
	// var numactivue Numactivue
	// if mylevel.Accumulative > accnum.Accumulative {
	// 	if camount > mylevel.Accumulative {
	// 		numactivue.Tivue = camount
	// 		mylevel.Accumulative = accnum.Accumulative
	// 	}
	// 	if camount <= mylevel.Accumulative && camount > mylevel.Accumulative-100000 {
	// 		if accnum.Accumulative <= 0 {
	// 			numactivue.Tivue = camount
	// 			mylevel.Accumulative = 100000
	// 		}
	// 		if accnum.Accumulative > 0 {
	// 			numactivue.Tivue = 100000
	// 			mylevel.Accumulative = accnum.Accumulative
	// 		}
	// 	}
	// 	if camount <= mylevel.Accumulative-100000 {
	// 		numactivue.Tivue = camount
	// 	}
	// }
	// for i := len(setlevel) - 1; i >= 0; i-- {
	// 	if mylevel.Accumulative+ins.Amount > setlevel[i].Levelvalue {
	// 		//平级
	// 		if mylevel.Accumulative+numactivue.Tivue > setlevel[i].Levelvalue {
	// 			ip.Profits += camount * setlevel[i].Percentreality
	// 			break
	// 		}
	// 		//越级
	// 		a := mylevel.Accumulative + camount - setlevel[i].Levelvalue
	// 		ip.Profits += a * setlevel[i].Percentreality
	// 		camount -= a
	// 	}
	// }
	// acc := mylevel.Accumulative
	//记录差多少越级
	// for i, v := range setlevel {
	// 	var u upgrade
	// 	if mylevel.Accumulative < v.Levelvalue && mylevel.Accumulative+ins.Amount >= v.Levelvalue {
	// 		u.need = v.Levelvalue - acc
	// 		u.levelvalue = mylevel.Levelvalue
	// 		//u.percent = mylevel.Percent
	// 		u.percent = mylevel.Percentreality
	// 		if i > 0 {
	// 			u.levelvalue = setlevel[i-1].Levelvalue
	// 			u.percent = setlevel[i-1].Percentreality
	// 		}
	// 		up = append(up, u)
	// 		acc = v.Levelvalue
	// 		continue
	// 	}
	// 	//最后一个也要放进去
	// 	// if mylevel.Accumulative+ins.Amount < v.Levelvalue && len(up) > 0 {
	// 	// 	u.levelvalue = setlevel[i-1].Levelvalue
	// 	// 	u.percent = setlevel[i-1].Percentreality
	// 	// 	if mylevel.Accumulative > v.Levelvalue {
	// 	// 		u.need = ins.Amount
	// 	// 	} else {
	// 	// 		u.need = mylevel.Accumulative + ins.Amount - setlevel[i-1].Levelvalue
	// 	// 	}
	// 	// 	up = append(up, u)
	// 	// 	break
	// 	// }
	// }
	isps = append(isps, ip)

	//判断是否存在升级情况
	if len(up) > 0 {
		mp := make(map[string]InvestmentShareProfit)
		for _, u := range up {
			for i, v := range ulevel {
				if u.levelvalue < v.Levelvalue {
					//var isp InvestmentShareProfit
					isp := mp[v.UserId]
					isp.Userid = v.UserId
					isp.Investmentid = param["investmentid"]
					camount := u.need
					prelevel := u.percent
					if i > 0 {
						//prelevel = ulevel[i-1].Percent
						prelevel = ulevel[i-1].Percentreality
					}
					for j := len(setlevel) - 1; j >= 0; j-- {
						if v.Accumulative+camount > setlevel[j].Levelvalue {
							//平级
							if v.Accumulative > setlevel[j].Levelvalue {
								//isp.Profits += camount * (v.Percent-prelevel)
								isp.Profits += camount * (v.Percentreality - prelevel)
								break
							}
							//越级
							a := v.Accumulative + camount - setlevel[j].Levelvalue
							//isp.Profits += a * (v.Percent-prelevel)
							isp.Profits += a * (v.Percentreality - prelevel)
							camount -= a
						}
					}
					mp[v.UserId] = isp
					//isps = append(isps, isp)
				}
			}
		}
		for k, _ := range mp {
			isps = append(isps, mp[k])
		}
	} else {
		for i, v := range ulevel {
			//if mylevel.Levelvalue < v.Levelvalue{
			var isp InvestmentShareProfit
			isp.Userid = v.UserId
			isp.Investmentid = param["investmentid"]
			// camount := ins.Amount
			//prelevel := mylevel.Percent
			// prelevel := mylevel.Percentreality
			if i > 0 {
				//prelevel = ulevel[i-1].Percent
				// prelevel = ulevel[i-1].Percentreality
			}
			// for j := len(setlevel) - 1; j >= 0; j-- {
			// 	if v.Accumulative+camount > setlevel[j].Levelvalue {
			// 		//平级
			// 		if v.Accumulative > setlevel[j].Levelvalue {
			// 			//isp.Profits += camount * (v.Percent-prelevel)
			// 			isp.Profits += camount * (v.Percentreality - prelevel)
			// 			break
			// 		}
			// 		//越级
			// 		a := v.Accumulative + camount - setlevel[j].Levelvalue
			// 		//isp.Profits += a * (v.Percent-prelevel)
			// 		isp.Profits += a * (v.Percentreality - prelevel)
			// 		camount -= a
			// 	}
			// }
			isps = append(isps, isp)
			//}
		}
	}

	flag := false
	sql := `insert into investmentprofit(investmentid,userid,profits)values`
	for _, v := range isps {
		if v.Profits <= 0 {
			continue
		}
		sql += `(` + v.Investmentid + `,` + v.Userid + `,` + utils.Float64ToString(v.Profits) + `),`
		flag = true
	}
	sql = strings.TrimRight(sql, ",")
	if flag {
		err = orm1.Exec(sql).Error
	}
	return err
}

//func (in *InvestmentShareProfit)AddInvestmentShareProfitImport(orm1 *gorm.DB,param map[string]string)(err error){
func (in *InvestmentShareProfit) AddInvestmentShareProfitImport(param map[string]string) (err error) {
	var le UserLevel
	var re Referrer
	re = re.GetReferrer(param["userid"])
	//获取自己等级
	mylevel, _ := le.GetSetUserByUserid(param["userid"])
	mylevel.UserId = param["userid"]
	//获取上级等级
	re.Referrers = strings.Trim(re.Referrers, ",")
	var ulevel []UserLevel

	if re.Referrers != "" {
		ulevel, err = mylevel.GetReferrerLevel(re.Referrers, mylevel.Accumulative)
		if err != nil && gorm.ErrRecordNotFound != err {
			return
		}
		//ulevel = append(ulevel, le...)
	}

	//获取设置等级
	var setlevel []UserLevel
	setlevel, _ = mylevel.GetSetUserLevel(mylevel.Levelvalue)
	//升级所需
	type upgrade struct {
		need       float64
		levelvalue float64
		percent    float64
	}
	var up []upgrade
	//计算自己的利润
	var isps []InvestmentShareProfit
	var ip InvestmentShareProfit
	camount, err := strconv.ParseFloat(param["amount"], 64)
	if err != nil {
		return
	}
	Amount := camount
	ip.Userid = param["userid"]
	ip.Investmentid = param["investmentid"]

	for i := len(setlevel) - 1; i >= 0; i-- {
		if mylevel.Accumulative+Amount > setlevel[i].Levelvalue {
			//平级
			if mylevel.Accumulative > setlevel[i].Levelvalue {
				ip.Profits += camount * setlevel[i].Percentreality
				break
			}
			//越级
			a := mylevel.Accumulative + camount - setlevel[i].Levelvalue
			ip.Profits += a * setlevel[i].Percentreality
			camount -= a
		}
	}
	acc := mylevel.Accumulative
	//记录差多少越级
	for i, v := range setlevel {
		var u upgrade
		if mylevel.Accumulative < v.Levelvalue && mylevel.Accumulative+Amount >= v.Levelvalue {
			u.need = v.Levelvalue - acc
			u.levelvalue = mylevel.Levelvalue
			//u.percent = mylevel.Percent
			u.percent = mylevel.Percentreality
			if i > 0 {
				u.levelvalue = setlevel[i-1].Levelvalue
				u.percent = setlevel[i-1].Percentreality
			}
			up = append(up, u)
			acc = v.Levelvalue
			continue
		}
		//最后一个也要放进去
		if mylevel.Accumulative+Amount < v.Levelvalue && len(up) > 0 {
			u.levelvalue = setlevel[i-1].Levelvalue
			u.percent = setlevel[i-1].Percentreality
			if mylevel.Accumulative > v.Levelvalue {
				u.need = Amount
			} else {
				u.need = mylevel.Accumulative + Amount - setlevel[i-1].Levelvalue
			}
			up = append(up, u)
			break
		}
	}
	isps = append(isps, ip)
	if param["investmentid"] == "1442492126456713216" {
		fmt.Println("sdfsdfsdf")
	}
	//判断是否存在升级情况
	if len(up) > 0 {
		mp := make(map[string]InvestmentShareProfit)
		for _, u := range up {
			for i, v := range ulevel {
				if u.levelvalue < v.Levelvalue {
					//var isp InvestmentShareProfit
					isp := mp[v.UserId]
					isp.Userid = v.UserId
					isp.Investmentid = param["investmentid"]
					camount := u.need
					prelevel := u.percent
					if i > 0 {
						//prelevel = ulevel[i-1].Percent
						prelevel = ulevel[i-1].Percentreality
					}
					for j := len(setlevel) - 1; j >= 0; j-- {
						if v.Accumulative+camount > setlevel[j].Levelvalue {
							//平级
							if v.Accumulative > setlevel[j].Levelvalue {
								isp.Profits += camount * (v.Percentreality - prelevel)
								break
							}
							//越级
							a := v.Accumulative + camount - setlevel[j].Levelvalue
							isp.Profits += a * (v.Percentreality - prelevel)
							camount -= a
						}
					}
					mp[v.UserId] = isp
					//isps = append(isps, isp)
				}
			}
		}
		for k, _ := range mp {
			isps = append(isps, mp[k])
		}
	} else {
		for i, v := range ulevel {
			//if mylevel.Levelvalue < v.Levelvalue{
			var isp InvestmentShareProfit
			isp.Userid = v.UserId
			isp.Investmentid = param["investmentid"]
			camount := Amount
			//prelevel := mylevel.Percent
			prelevel := mylevel.Percentreality
			if i > 0 {
				//prelevel = ulevel[i-1].Percent
				prelevel = ulevel[i-1].Percentreality
			}
			for j := len(setlevel) - 1; j >= 0; j-- {
				if v.Accumulative+camount > setlevel[j].Levelvalue {
					//平级
					if v.Accumulative > setlevel[j].Levelvalue {
						isp.Profits += camount * (v.Percentreality - prelevel)
						break
					}
					//越级
					a := v.Accumulative + camount - setlevel[j].Levelvalue
					isp.Profits += a * (v.Percentreality - prelevel)
					camount -= a
				}
			}
			isps = append(isps, isp)
			//}
		}
	}

	flag := false
	sql := `insert into investmentprofit(investmentid,userid,profits)values`
	for _, v := range isps {
		if v.Profits <= 0 {
			continue
		}
		sql += `(` + v.Investmentid + `,` + v.Userid + `,` + utils.Float64ToString(v.Profits) + `),`
		flag = true
	}
	sql = strings.TrimRight(sql, ",")
	if flag {
		//if err = orm1.Exec(sql).Error;err!=nil{
		if err = orm.Eloquent.Exec(sql).Error; err != nil {
			return
		}
	}
	return
}
