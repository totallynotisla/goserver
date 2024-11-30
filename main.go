package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	ApiRoutes "github.com/mangadi3859/goserver/api"
	db "github.com/mangadi3859/goserver/tools"
)

var DB *sql.DB

func main() {
    DB = db.DbConnect();

	reloader := gin.Default()
    api := reloader.Group("/api")
    api.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "welcome to api",
        })
    })

    api.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    //Routes
    ApiRoutes.LoginRoutes(api)
    
    reloader.Run(":8080")
    defer DB.Close()
}
