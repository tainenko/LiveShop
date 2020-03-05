package middlewares

import (
	"fmt"
	"github.com/LiveShop/controllers"
	"github.com/LiveShop/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"time"
)

var identityKey = "Hello World"

type JwtAuthorizator func(data interface{}, c *gin.Context) bool

func AuthMiddleware(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "gin jwt",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user:=controllers.GetUserFromContext(c)
			fmt.Println(user)
			if err := c.Bind(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			fmt.Print(user.Email,user.Name)
			models.HashPassword(&user)
			id := models.QueryUserId(user.Email, user.Password)
			user.Password = ""
			if id > 0 {
				return &models.RegisterInfo{
					Name:  user.Name,
					Email: user.Email,
				}, nil
			} else {
				return nil, jwt.ErrFailedAuthentication
			}
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return

}
func AllUserAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}
func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}
