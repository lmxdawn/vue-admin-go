package req

type AuthRoleAuthReq struct {
	RoleId    int64   `json:"role_id" binding:"required"`    // 角色ID
	AuthRules []int64 `json:"auth_rules" binding:"required"` // 权限规则
}
