package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ApiRoutes "github.com/totallynotisla/goserver/api"
	db "github.com/totallynotisla/goserver/tools"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	db.Con = db.DbConnect()
	err = db.InitDB(db.Con)

	if err != nil {
		panic(err.Error())
	}

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
	ApiRoutes.Handler(api)

	reloader.Run(":8080")
	defer db.Con.Close()
}
