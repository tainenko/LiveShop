package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	firstname := c.DefaultQuery("name", "kin")
	age := c.Query("age")
	c.String(http.StatusOK, "Hello %s %s", firstname, age)
}
