package router
import (
	"github.com/gin-gonic/gin"
)
func InitRouter(){
	router:=gin.Default()
	router.GET("/",hello)
	router.Run(":8080")
}