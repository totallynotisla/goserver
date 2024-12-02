package tools

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const COOKIE_KEY = "auth"
const COOKIE_EXPIRE = 60 * 60 * 24 * 7
const COOKIE_USER = "user"
const JWT_SECRET_STRING = "mungkinbisajadimilikusemogalagucintainibersarangtepatdihati"

var JWT_METHOD = jwt.SigningMethodHS256
var JWT_SECRET = []byte(JWT_SECRET_STRING)

func Register(data RegisterData) (User, error) {
	uid := uuid.New().String()
	hashed := hashPassword(data.Password)

	_, err := Con.Exec(`INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)`, uid, data.Username, hashed, data.Email)
	if err != nil {
		return User{}, err
	}

	res := User{}
	Con.Get(&res, "SELECT * FROM users WHERE id = $1", uid)

	return res, nil
}

func Login(data LoginData, c *gin.Context) (User, Session, error) {
	user := User{}
	Con.Get(&user, "SELECT * FROM users WHERE username = $1 OR email = $1", data.Username)

	if user.ID == "" {
		return User{}, Session{}, fmt.Errorf("User not found")
	}
	if !verifyPassword(data.Password, user.Password) {
		return User{}, Session{}, nil
	}

	session, err := createSession(user, c)
	if err != nil {
		return User{}, Session{}, err
	}

	return user, session, nil
}

func IsAuthenticated(c *gin.Context) bool {
	//TODO: Implement this
	token, err := c.Cookie(COOKIE_KEY)
	if err != nil {
		return false
	}

	if !checkSession(token) {
		c.SetCookie(COOKIE_KEY, "", -1, "/", "", false, true)
		c.SetCookie(COOKIE_USER, "", -1, "/", "", false, true)
		return false
	}

	userCookie, err := c.Cookie(COOKIE_USER)
	if err != nil {
		return false
	}

	// user := User{}
	userJwt, err := readJwt[User](userCookie)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(userJwt)

	return true
}

func GetUserFromToken(token string) User {
	user := User{}
	Con.Get(&user, "SELECT * FROM users WHERE id = (SELECT userId FROM session WHERE token = $1)", token)
	return user
}

func createSession(user User, c *gin.Context) (Session, error) {
	token := uuid.New().String()
	expiresAt := time.Now().Add(time.Second * COOKIE_EXPIRE)

	_, err := Con.Exec("INSERT INTO session (token, expires_at, userId) VALUES ($1, $2, $3)", token, expiresAt, user.ID)
	if err != nil {
		return Session{}, err
	}

	userJwt, err := createJwt(user)
	c.SetCookie(COOKIE_USER, userJwt, int(COOKIE_EXPIRE), "/", "", false, true)
	c.SetCookie(COOKIE_KEY, token, int(COOKIE_EXPIRE), "/", "", false, true)

	return Session{Token: token, ExpiresAt: expiresAt, UserID: user.ID}, nil
}

type jwtWithClaims[T any] struct {
	jwt.RegisteredClaims
	Claims T
}

func createJwt[T any](data T) (string, error) {
	claims := jwtWithClaims[T]{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gerawana",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(COOKIE_EXPIRE)),
		},
		Claims: data,
	}

	token := jwt.NewWithClaims(JWT_METHOD, claims)
	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func readJwt[T any](tokenString string) (T, error) {
	var _zero T
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SECRET, nil
	})

	if err != nil {
		return _zero, err
	}

	return token.Claims.(jwtWithClaims[T]).Claims, nil
}

func checkSession(token string) bool {
	session := Session{}
	Con.Get(&session, "SELECT * FROM session WHERE token = $1 AND expires_at > $2", token, time.Now())

	return session.Token != ""
}

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
