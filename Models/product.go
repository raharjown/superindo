package Models

import (
	"superindo/Database"
)

type Product struct {
	Id       int    `json:"id"`
	Kategori string `json:"kategori"`
	Nama     string `json:"nama"`
	Jenis    string `json:"jenis"`
	Varian   string `json:"varian"`
	Ukuran   string `json:"ukuran"`
	Harga    int    `json:"harga"`
}

func (p *Product) TableName() string {
	return "produk"
}

func GetProductByCategory(product *[]Product, category string) (err error) {
	err = Database.DB.
		Where("kategori=?", category).
		Find(product).
		Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id int) (err error) {
	err = Database.DB.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}
