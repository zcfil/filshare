package main

import "fmt"

func main() {

	//q := NewQQwry("./temp/logs/qqwry.dat")
	//q.Find("1.24.41.0")
	//fmt.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
	// output:
	// 2014/02/22 22:10:32 ip:180.89.94.90, Country:北京市, City:长城宽带
	//ipstr:= p.Get("113.87.167.223")
	//fmt.Println(ipstr)

	fmt.Println("err")
	region, err := New("./temp/logs/ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	ip, err := region.MemorySearch("113.87.167.223")
	fmt.Println(ip, err)
}
