package router
import (
	"github.com/LiveShop/controllers"
	"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine{
	router:=gin.Default()
	router.GET("/",hello)
	router.POST("/register",controllers.Register)
	return router
}