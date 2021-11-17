package req

type LoginEmailPwdReq struct {
	Email string `json:"email" binding:"required,min=2,max=30"` // 邮箱
	Pwd   string `json:"pwd" binding:"required,min=2,max=30"`   // 密码
}
