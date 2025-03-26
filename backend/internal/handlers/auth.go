package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type User struct {
	Id           string `json:"id" gorm:"primary_key"`
	Username     string `json:"username"`
	Password     string `json:"-"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LastLogin    int64  `json:"last_login"`
	LastUpdate   int64  `json:"last_update"`
}

type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func Register(c *gin.Context) {

	var body RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	var userId = generateId()
	var accessToken, _ = generateToken(userId, "accessToken")
	var refreshToken, _ = generateToken(userId, "refreshToken")

	//TODO: change hardCoded string by getting info in request body
	var user User = User{
		Id:           userId,
		Username:     body.Username,
		Password:     body.Password,
		Email:        body.Email,
		Role:         "User",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		LastLogin:    time.Now().Unix(),
		LastUpdate:   time.Now().Unix(),
	}

	c.JSON(200, gin.H{
		"message": "success",
		"user":    user,
	})
}

func generateId() string {
	var baseID = "U-" + strconv.Itoa(rand.Intn(10000)) + "-"

	var timestampLastSixDigits = strconv.FormatInt(time.Now().Unix(), 10)[6:]

	return baseID + timestampLastSixDigits
}

func generateToken(userID string, tokenType string) (string, error) {

	if tokenType != "accessToken" && tokenType != "refreshToken" {
		return "error", fmt.Errorf("invalid token type provided, %s", tokenType)
	}

	const accessTokenExpiry = time.Minute * 15
	const refreshTokenExpiry = time.Hour * 24 * 7
	var expiryDuration time.Duration

	if tokenType == "accessToken" {
		expiryDuration = accessTokenExpiry
	} else {
		expiryDuration = refreshTokenExpiry
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiryDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	signedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return signedToken.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})
}
