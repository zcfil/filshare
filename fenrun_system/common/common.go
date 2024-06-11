package common

import "time"

func TimeToDay(time time.Time) int32 {
	year := time.Year()
	month := time.Month()
	day := time.Day()
	return int32(year)*10000 + int32(month)*100 + int32(day)
}

func DefaultString(des, def string) string {
	if des != "" {
		return des
	}
	return def
}

const MaxAddressStringLength = 2 + 84
const MainnetPrefix = "f"
const TestnetPrefix = "t"

// 钱包的地址的两种长度
const (
	TYPE_2_LEN = 41
	TYPE_1_LEN = 86
)

func CheckAddress(a string) bool {
	l := len(a)
	if l != TYPE_1_LEN && l != TYPE_2_LEN {
		return false
	}

	if l > MaxAddressStringLength || l < 3 {
		return false
	}

	if string(a[0]) != MainnetPrefix && string(a[0]) != TestnetPrefix {
		return false
	}
	switch a[1] {
	case '1':
	case '2':
	case '3':
	default:
		return false
	}
	return true
}
