package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lmxdawn/vue-admin-go/internal/code"
	"github.com/lmxdawn/vue-admin-go/pkg/util"
	"github.com/lmxdawn/vue-admin-go/req"
	"github.com/lmxdawn/vue-admin-go/res"
)

// ByEmailPwd ...
// @Tags 授权
// @Summary 邮箱加密码登录
// @Produce json
// @Security ApiKeyAuth
// @Param login body req.LoginEmailPwdReq true "登录参数"
// @Success 200 {object} res.Response{data=res.LoginInfoRes}
// @Router /login/byEmailPwd [post]
func ByEmailPwd(c *gin.Context) {

	var req req.LoginEmailPwdReq

	if err := c.ShouldBindJSON(&req); err != nil {
		res.APIResponse(c, code.ErrParam, nil)
	}

	if req.Email == "" || req.Pwd == "" {
		res.APIResponse(c, code.ErrParam, nil)
	}

	uid := int64(1)

	token := util.CreateToken(uid)

	loginInfo := res.LoginInfoRes{
		Uid:   uid,
		Token: token,
	}

	res.APIResponse(c, nil, loginInfo)
}

func Test() {

	//if uid, ok := c.Keys["uid"]; ok {
	//
	//	loginInfo := res.LoginInfoRes {
	//		Uid: uid.(int64),
	//		Token: "tttt",
	//	}
	//
	//	res.APIResponse(c, nil, loginInfo)
	//	return
	//}
	//res.APIResponse(c, code.ErrToken, nil)
}
