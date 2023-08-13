package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	Image     string    `json:"image"`
	Url       string    `json:"url"`
	Category  string    `json:"category"`
	TotalSold string    `json:"total_sold"`
	ShopName  string    `json:"shop_name"`
}

func InsertProduct(product []Product) (*[]Product, error) {
	return &product, db.Create(&product).Error
}

func CountProduct() (int64, error) {
	var productCount int64
	if err := db.Model(&Product{}).Count(&productCount).Error; err != nil {
		return 0, err
	}
	return productCount, nil
}

func GetAllProducts(productPerPage int, offset int) (*[]Product, error) {
	var products []Product
	err := db.Limit(productPerPage).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}
