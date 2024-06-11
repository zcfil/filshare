package auth

import (
	"fmt"
	"time"
)

const (
	OnceStatement     = 0
	SalesmanStatement = 1
	CustomerStatement = 2
	SummaryStatement  = 3
)

func AddStatement() {
	go func() {
		for i := 0; ; i++ {
			if i != 0 {
				now := time.Now()
				next := now.AddDate(0, 1, 0)

				next = time.Date(next.Year(), next.Month(), 1, 0, 0, 0, 0, next.Location())

				fmt.Println(next.Sub(now))

				c := time.NewTimer(next.Sub(now))
				<-c.C
			}
			//var s models.Statement
			now := time.Now()
			pre := time.Date(now.Year(), now.Month(), 0, 23, 59, 59, 0, now.Location()).Format("2006-01-02 15:04:05")

			fmt.Println(pre)
			//if !s.GetStatementBool(pre,OnceStatement){
			//	if err := s.AddStatementConfigOnce(pre);err!=nil{
			//		log.Println("保存一次性分配报表：",err.Error())
			//	}
			//}
			//if !s.GetStatementBool(pre,SalesmanStatement){
			//	if err := s.AddStatementSalesman(pre);err!=nil{
			//		log.Println("保存业务员报表：",err.Error())
			//	}
			//
			//}

			param := make(map[string]string)
			param["date"] = pre
			//if !s.GetSummaryBool(pre){
			//	if err := s.AddStatementSummary(param);err!=nil{
			//		log.Println("保存汇总报表报表：",err.Error())
			//	}
			//}
		}
	}()

}
