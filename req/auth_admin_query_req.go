package req

type AuthAdminQueryReq struct {
	ListPage
	Username string  `json:"username"` // 用户名
	Status   int     `json:"status"`   // 用户状态（-1: 全部，0：禁用； 1：正常 ；2：未验证）
	RoleId   int64   `json:"role_id"`  // 角色ID列表
	Ids      []int64 `json:"ids"`      // 用户ID列表
}
