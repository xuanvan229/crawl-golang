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

func UpdateProduct(product Product) (*Product, error) {
	return &product, db.Save(&product).Error
}

func CountProduct(name string) (int64, error) {
	var productCount int64
	if name != "" {
		err := db.Model(&Product{}).Where("name ILIKE ?", "%"+name+"%").Count(&productCount).Error
		if err != nil {
			return 0, err
		}
		return productCount, nil
	}
	if err := db.Model(&Product{}).Count(&productCount).Error; err != nil {
		return 0, err
	}
	return productCount, nil
}

func GetAllProducts(productPerPage int, offset int, name string) (*[]Product, error) {
	var products []Product
	if name != "" {
		err := db.Limit(productPerPage).Offset(offset).Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			return nil, err
		}
		return &products, nil
	}
	err := db.Limit(productPerPage).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func GetProductWrongImage() (*[]Product, error) {
	var products []Product
	name := "TB13MLwbOLaK1RjSZFxXXamPFXa"
	err := db.Where("image ILIKE ?", "%"+name+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}
