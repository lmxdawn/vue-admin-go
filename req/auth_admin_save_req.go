package req

type AuthAdminSaveReq struct {
	Id       int64   `json:"id"`                                    // 用户ID（修改时必带）
	Username string  `json:"username" binding:"required"`           // 用户名
	Password string  `json:"password" binding:"required"`           // 密码
	Status   int     `json:"status" binding:"required,oneof=0 1 2"` // 状态（0：禁用； 1：正常 ；2：未验证）
	Roles    []int64 `json:"roles"`                                 // 角色IDs
}
