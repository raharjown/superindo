package Controllers

import (
	"net/http"
	"strconv"
	"superindo/Models"

	"github.com/gin-gonic/gin"
)

func GetProductByCategory(c *gin.Context) {
	var produk []Models.Product
	category := c.Params.ByName("category")
	err := Models.GetProductByCategory(&produk, category)
	if err != nil {
		c.JSON(http.StatusNotFound, "Category Not Found")
		c.Abort()
	} else {
		c.JSON(http.StatusOK, produk)
	}
}

func GetProductByID(c *gin.Context) {
	var produk Models.Product
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := Models.GetProductByID(&produk, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Id Not Found")
		c.Abort()
	} else {
		c.JSON(http.StatusOK, produk)
	}
}
