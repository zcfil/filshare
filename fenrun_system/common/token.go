package common

import (
	"fmt"
	"xAdmin/define"

	uuid "github.com/satori/go.uuid"
)

func GenCustomerTokenKey(token string) string {
	return fmt.Sprintf(define.TOKEN_PREFIX+":%s", token)
}

func GenCustomerToken() string {
	return uuid.NewV4().String()
}
