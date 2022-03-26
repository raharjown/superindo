package Controllers

import (
	"net/http"
	"superindo/Models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLogin Models.User

	c.BindJSON(&userLogin)
	err := Models.Login(&userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
	} else {
		sign := jwt.New(jwt.GetSigningMethod("HS256"))
		secret := "secret"
		token, err := sign.SignedString([]byte(secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error Generate Token")
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
	} else {
		c.JSON(http.StatusCreated, user)
	}
}
