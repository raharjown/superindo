package Routes

import (
	"superindo/Auth"
	"superindo/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("/")
	{
		// user
		grp1.POST("user", Controllers.CreateUser)
		grp1.POST("login", Controllers.Login)

		// category product
		grp1.GET("category", Auth.Auth, Controllers.GetCategory)

		// product
		grp1.GET("products/:category", Auth.Auth, Controllers.GetProductByCategory)
		grp1.GET("product/:id", Auth.Auth, Controllers.GetProductByID)

		// cart
		grp1.POST("cart", Auth.Auth, Controllers.CreateCart)

	}
	return r
}
