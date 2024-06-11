package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type DataCrawler struct {
	periodTime int64
}

func NewDataCrawler(periodTime int64) *DataCrawler {
	dc := new(DataCrawler)
	dc.periodTime = periodTime
	return dc
}

func (this *DataCrawler) GetData() (bodyStr string, err error) {
	resp, err1 := http.Get("https://filfox.info")

	if err1 != nil {
		fmt.Println(err)
		err = err1
		return
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
		return
	}

	bodyStr = string(body)
	return
}

func (this *DataCrawler) WriteData() (ret string, h1 string, err error) {
	bodyStr, err := this.GetData()
	if err != nil {
		fmt.Println("Get data body error:", err)
		return
	}

	reg1 := regexp.MustCompile(`<div class="text-left lg:text-center text-sm lg:text-2xl items-start lg:mx-auto"> (.*?) </div>`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result := reg1.FindAllStringSubmatch(bodyStr, -1)
	info := &OfficialInfo{}
	info.Income24FT = result[4][1]
	Info := strings.Split(info.Income24FT, " ")
	ret = Info[0]
	info.Height = result[0][1] // 区块高度
	h1 = this.getHeight1(info.Height)
	return ret, h1, nil
}

func (this *DataCrawler) getHeight1(h string) (height string) {
	ret := make([]byte, 0)
	for _, b := range h {
		if b == ',' {
			continue
		}
		ret = append(ret, byte(b))
	}
	height = string(ret)
	return
}
