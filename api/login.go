package ApiRoutes

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/totallynotisla/goserver/tools"
)

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		json.NewDecoder(c.Request.Body).Decode(&body)
		body.Username = strings.ToLower(body.Username)

		if body.Username == "" || body.Password == "" {
			c.JSON(400, gin.H{
				"message": "Missing required fields",
				"status":  "FAILED",
			})
			return
		}

		user := tools.User{}
		tools.Con.Get(&user, `SELECT * FROM users WHERE username=$1 OR email=$1`, body.Username)

		if user.ID == "" {
			c.JSON(404, gin.H{
				"message": "User not found",
				"status":  "FAILED",
			})
			return
		}

		userDb, session, err := tools.Login(tools.LoginData(body), c)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Server Error",
				"status":  "FAILED",
			})
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Login users with username:", body.Username)
		c.JSON(200, gin.H{
			"message": "Success",
			"status":  "OK",
			"data": gin.H{
				"user": tools.UserResponse{
					ID:       userDb.ID,
					Username: userDb.Username,
					Email:    userDb.Email,
				},
				"session": tools.SessionResponse{
					Token:     session.Token,
					ExpiresAt: session.ExpiresAt,
					UserID:    session.UserID,
				},
			},
		})
	}
}
