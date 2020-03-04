package router

import (
	"github.com/LiveShop/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", hello)
	v1 := router.Group("/v1")
	{
		v1.POST("/register", controllers.Register)
	}
	return router
}
