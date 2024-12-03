package tools

import (
	"time"
)

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResetPassword struct {
	ID        string    `db:"id"`
	ExpiresAt time.Time `db:"expires_at"`
	UserID    string    `db:"userId"`
	OTP       string    `db:"otp"`
}

type Session struct {
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	UserID    string    `db:"userId"`
}

type SessionResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    string    `json:"userId"`
}

type ShortLink struct {
	ID        string    `db:"id"`
	Redirect  string    `db:"redirect"`
	Link      string    `db:"link"`
	AuthorID  string    `db:"authorId"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"createdAt"`
}

type Visit struct {
	ID          string    `db:"id"`
	ShortLinkID string    `db:"shortLinkId"`
	VisitedAt   time.Time `db:"visitedAt"`
}

type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
