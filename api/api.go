package ApiRoutes

import (
	"github.com/gin-gonic/gin"
)

func Handler(route *gin.RouterGroup) {
	route.POST("/login", login())
	route.POST("/register", register())
}
