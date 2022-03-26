package Controllers

import (
	"net/http"
	"superindo/Models"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	var category []Models.Category
	err := Models.GetCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}
