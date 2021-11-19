package req

type LoginEmailPwdReq struct {
	UserName string `json:"user_name" binding:"required"` // 邮箱
	Pwd      string `json:"pwd" binding:"required"`       // 密码
}
