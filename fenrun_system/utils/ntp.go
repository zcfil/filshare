package utils

import (
	"github.com/beevik/ntp"
	"time"
)

func NewNtp() time.Time {
	response, err := ntp.Query("ntp.aliyun.com")
	if err != nil {
		return time.Time{}
	}
	time := time.Now().Add(response.ClockOffset)
	return time
}
