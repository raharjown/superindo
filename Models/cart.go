package Models

import (
	"fmt"
	"superindo/Database"
)

type Cart struct {
	Id       int    `json:"id"`
	IdUser   string `json:"id_user"`
	IdProduk int    `json:"id_produk"`
	Jumlah   int    `json:"jumlah"`
}

func (c *Cart) TableName() string {
	return "keranjang"
}

func GetIsExistCart(cart *Cart) (isExist bool, err error) {
	var result struct {
		Found bool
	}
	err = Database.DB.
		Select("count(*) > 0").
		Where("id_user=? and id_produk=?", cart.IdUser, cart.IdProduk).
		Find(&result).
		Error
	fmt.Println(result)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetCartByParam(_cart *Cart) (err error) {
	err = Database.DB.
		Where("id_user=? and id_produk=?", _cart.IdUser, _cart.IdProduk).
		Find(_cart).
		Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCart(cart *Cart) (err error) {
	if err = Database.DB.Save(cart).Error; err != nil {
		return err
	}
	return nil
}

func CreateCart(cart *Cart) (err error) {
	if err = Database.DB.Create(cart).Error; err != nil {
		return err
	}
	return nil
}
