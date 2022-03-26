package Auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		c.JSON(http.StatusUnauthorized, "Missing Token")
		c.Abort()
	} else {
		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if token != nil && err == nil {
			fmt.Println("token verified")
		} else {
			result := gin.H{
				"message": "Not Authorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		}
	}
}
