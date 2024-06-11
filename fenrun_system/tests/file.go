package main

import (
	"fmt"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"xAdmin/utils"
)

func main() {
	//s := "0.014022222222222222222222"
	//stw, _ := strconv.ParseFloat(s, 64)
	//
	//s1, err := utils.ParseFIL(fmt.Sprintf("%.18f", stw))
	//if err != nil {
	//	fmt.Println("异常")
	//	return
	//}
	//
	//total := 100 * stw
	//fmt.Println("总收益100%", total)
	//stk := float64(0.25)
	//s3 := float64(total) * stk
	//fmt.Println("总收益25%", s3)
	//s13, err := utils.ParseFIL(fmt.Sprintf("%.18f", s3))
	//if err != nil {
	//	fmt.Println("异常")
	//	return
	//}
	//s133 := abi.TokenAmount(s13)
	//fmt.Println(utils.FIL(s133))
	//fmt.Println(s133.Int64())
	//fmt.Println("总收益25%1", s13)
	//
	//stk1 := float64(0.75)
	//s31 := float64(total) * stk1
	//fmt.Println("总收益75%", s31)
	//tol := s31 / 180
	//fmt.Println("180天", tol)
	//fmt.Println(s1.String())
	//s2 := abi.TokenAmount(s1)
	//fmt.Println(s2.String())
	//fmt.Println(utils.FIL(s2))
	type s struct {
		T float64 `json:"s"`
	}
	k := s{}

	fmt.Println(k.T)

	ties := time.Now()
	fmt.Println(utils.TimeHMSStr(ties.Unix()))
}

func strtobig(str string) (s abi.TokenAmount) {
	s1, err := utils.ParseFIL(str)
	if err != nil {
		fmt.Println("异常")
		return abi.TokenAmount{}
	}
	s = abi.TokenAmount(s1)
	return
}
