package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lmxdawn/vue-admin-go/internal/code"
	"github.com/lmxdawn/vue-admin-go/pkg/util"
	"github.com/lmxdawn/vue-admin-go/res"
)

// AuthRequired 认证中间件
// @required 是否必须token
func AuthRequired(required bool) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("x-token")
		uid, err := util.VerifyToken(token)
		if required && err != nil {
			c.Abort()
			res.APIResponse(c, code.ErrToken, nil)
		}

		c.Set("uid", uid)

	}

}
