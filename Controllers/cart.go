package Controllers

import (
	"net/http"
	"superindo/Models"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var cart Models.Cart
	var err error

	c.BindJSON(&cart)
	_cart := cart
	if cart.Jumlah <= 0 {
		c.JSON(http.StatusBadRequest, "jumlah tidak boleh kurang dari 1")
		c.Abort()
		return
	}

	var user Models.User
	err = Models.GetUserById(&user, cart.IdUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "User ID Not Found")
		c.Abort()
		return
	}

	var produk Models.Product
	err = Models.GetProductByID(&produk, cart.IdProduk)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Product ID Not Found")
		c.Abort()
		return
	}

	err = Models.GetCartByParam(&_cart)
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	if _cart.Id == 0 { // create new cart
		err = Models.CreateCart(&cart)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Fail to create cart")
			c.Abort()
			return
		} else {
			c.JSON(http.StatusCreated, cart)
			return
		}
	} else { // update new cart
		_cart.Jumlah = cart.Jumlah + _cart.Jumlah

		err = Models.UpdateCart(&_cart)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Fail to create cart")
			c.Abort()
			return
		} else {
			c.JSON(http.StatusCreated, _cart)
			return
		}
	}

}
