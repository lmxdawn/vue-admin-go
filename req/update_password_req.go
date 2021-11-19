package req

type UpdatePasswordReq struct {
	AdminId     int64  `json:"admin_id" binding:"required"`     // 管理员ID
	OldPassword string `json:"old_password" binding:"required"` // 旧密码
	NewPassword string `json:"new_password" binding:"required"` // 新密码
}
