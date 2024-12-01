package ApiRoutes

import (
	"github.com/gin-gonic/gin"
)


func LoginRoutes(route *gin.RouterGroup) {
	route.GET("/login", login())
}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	}
}