package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通登录
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //被禁用的用户
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	switch r {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "被禁言的用户"
	default:
		return "其他"
	}
}
