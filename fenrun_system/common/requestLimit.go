package common

import (
	"crypto/md5"
	"fmt"
	"time"
	log "xAdmin/logrus"
	"xAdmin/redisClient"
)

const CHECK_REQUEST_TIME = time.Minute * 10

func CheckRequest(reqMsg string) bool {
	md5Data := md5.Sum([]byte(reqMsg))
	dataStr := fmt.Sprintf("%x", md5Data)

	key := genCheckKey(dataStr)
	isExist, err := redisClient.RedisClient.Exists(key).Result()
	if err != nil {
		log.Error("请求redis 错误, err:", err.Error())
		return false
	}
	if isExist == 1 {
		return false
	}

	if err = redisClient.RedisClient.Set(key, 1, CHECK_REQUEST_TIME).Err(); err != nil {
		return false
	}
	return true
}

func genCheckKey(data string) string {
	return fmt.Sprintf("checkRequest:%s", data)
}
