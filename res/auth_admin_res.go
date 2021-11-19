package res

import "time"

type AuthAdminRes struct {
	Id            int64     `json:"id"`              // 管理员ID
	Username      string    `json:"username"`        // 用户名
	LastLoginIp   string    `json:"last_login_ip"`   // 最后登录ip
	LastLoginTime time.Time `json:"last_login_time"` // 最后登录时间
	Status        int       `json:"status"`          // 用户状态（0：禁用； 1：正常 ；2：未验证）
	Roles         []int64   `json:"roles"`           // 角色列表
}
