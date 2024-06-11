package models

const (
	PASS_VERIFICATION    = 1 // 通过审核
	NO_PASS_VERIFICATION = 0 // 没通过审核
)

//审核角色权限
var AUDIT = map[string]bool{
	"admin":   true,
	"finance": true,
	"boss":    true,
}
