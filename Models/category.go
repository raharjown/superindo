package Models

import "superindo/Database"

type Category struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

func (c *Category) TableName() string {
	return "kategori_produk"
}

func GetCategory(category *[]Category) (err error) {
	err = Database.DB.Find(category).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryByID(category *[]Category) (err error) {
	err = Database.DB.Find(category).Error
	if err != nil {
		return err
	}
	return nil
}
