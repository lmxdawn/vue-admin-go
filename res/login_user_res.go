package res

type LoginUserRes struct {
	Id        int64    `json:"id"`         // 用户ID
	Username  string   `json:"username"`   // 用户名
	Avatar    string   `json:"avatar"`     // 头像
	AuthRules []string `json:"auth_rules"` // 权限列表
}
