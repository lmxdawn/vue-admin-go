package req

type AuthRoleQueryReq struct {
	*ListPage
	Name   string `json:"name" binding:"required"`             // 权限名称
	Status int    `json:"status" binding:"required,oneof=0 1"` // 状态（1：正常，0：禁用）
}
