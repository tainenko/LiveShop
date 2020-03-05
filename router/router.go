package router

import (
	"github.com/LiveShop/controllers"
	"github.com/LiveShop/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/register", controllers.RegisterPost)
		v1.POST("/login", middlewares.AuthMiddleware(middlewares.AllUserAuthorizator).LoginHandler)
	}
	auth := router.Group("/auth")
	auth.Use(middlewares.AuthMiddleware(middlewares.AllUserAuthorizator).MiddlewareFunc())
	{
		auth.GET("/hello", hello)
	}
	return router
}
