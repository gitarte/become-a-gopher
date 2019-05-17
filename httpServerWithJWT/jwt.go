package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret -
var Secret = []byte("TOP SECRET KEY")

// Credentials -
type Credentials struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

// JWTGetToken -
func JWTGetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cred := Credentials{}
		err := c.Bind(&cred)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if !(cred.User == "admin" && cred.Pass == "pass") {
			c.String(http.StatusUnauthorized, "go away!")
			return
		}

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 60, // one minute from now
				Issuer:    "test",
			},
		)

		tokenString, err := token.SignedString(Secret)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

// JWTCheckToken -
func JWTCheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header["Authorization"]
		if len(auth) == 0 {
			c.String(http.StatusUnauthorized, "no token")
			c.Abort()
			return
		}
		tokenString := auth[0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		if token.Valid {
			c.Next()
		}
		return
	}
}
