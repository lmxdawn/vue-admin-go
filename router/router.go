package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lmxdawn/vue-admin-go/middleware"
)

// InitRouter ...
func InitRouter() *gin.Engine {

	server := gin.Default()

	// 中间件
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	register := server.Group("/register", middleware.AuthRequired(false))
	{
		registerAPI(register)
	}

	login := server.Group("/login", middleware.AuthRequired(false))
	{
		loginAPI(login)
	}

	auth := server.Group("/auth", middleware.AuthRequired(true))
	{
		authAPI(auth)
	}

	return server
}
