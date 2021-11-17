package router

import "github.com/gin-gonic/gin"

// authAPI ...
func authAPI(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/admin/info")
}
