package res

import "time"

type LoginInfoRes struct {
	Uid   int64  // 用户ID
	Token string // 登录token
	Time  *time.Time
}
