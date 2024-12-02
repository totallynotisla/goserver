package ApiRoutes

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/totallynotisla/goserver/tools"
)

func register() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}{}
		json.NewDecoder(c.Request.Body).Decode(&body)
		body.Username = strings.ToLower(body.Username)
		body.Email = strings.ToLower(body.Email)

		if body.Username == "" || body.Password == "" || body.Email == "" {
			c.JSON(400, gin.H{
				"message": "Missing required fields",
				"status":  "FAILED",
			})
			return
		}

		user := tools.User{}
		tools.Con.Get(&user, `SELECT * FROM users WHERE username=$1 OR email=$2`, body.Username, body.Email)

		if user.ID != "" {
			c.JSON(409, gin.H{
				"message": "User already exists",
				"status":  "FAILED",
			})
			return
		}

		insertedUser, err := tools.Register(tools.RegisterData(body))
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Server Error",
				"status":  "FAILED",
			})
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Registering users with username:", body.Username, "email:", body.Email)
		c.JSON(200, gin.H{
			"message": "Success",
			"status":  "OK",
			"data":    insertedUser,
		})
	}
}
