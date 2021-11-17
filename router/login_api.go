package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lmxdawn/vue-admin-go/app"
)

// RegisterAPI ...
func registerAPI(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/byTel")
}

// LoginAPI ...
func loginAPI(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/byEmailPwd", app.ByEmailPwd)
}
