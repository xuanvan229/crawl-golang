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

type CategoryLazada struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Url     string    `json:"url"`
	Crawled bool      `json:"crawled"`
}

func InsertCategory(categories []CategoryLazada) (*[]CategoryLazada, error) {
	return &categories, db.Create(&categories).Error
}

func UpdateCategory(category CategoryLazada) (*CategoryLazada, error) {
	return &category, db.Save(&category).Error
}

func GetCategoryNotCrawled() (*[]CategoryLazada, error) {
	var categories []CategoryLazada
	err := db.Where("crawled = ?", false).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}

func InsertProduct(product []Product) (*[]Product, error) {
	return &product, db.Create(&product).Error
}

func UpdateProduct(product Product) (*Product, error) {
	return &product, db.Save(&product).Error
}

func CountProduct(name string, category string) (int64, error) {
	var productCount int64
	if name != "" {
		if category != "" {
			err := db.Model(&Product{}).Where("name ILIKE ? AND category = ?", "%"+name+"%", category).Count(&productCount).Error
			if err != nil {
				return 0, err
			}
			return productCount, nil
		}
		err := db.Model(&Product{}).Where("name ILIKE ?", "%"+name+"%").Count(&productCount).Error
		if err != nil {
			return 0, err
		}
		return productCount, nil
	}

	if category != "" {
		err := db.Model(&Product{}).Where("category = ?", category).Count(&productCount).Error
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

func GetAllProducts(productPerPage int, offset int, name string, category string) (*[]Product, error) {
	var products []Product
	if name != "" {
		if category != "" {
			err := db.Limit(productPerPage).Offset(offset).Where("name ILIKE ? AND category = ?", "%"+name+"%", category).Find(&products).Error
			if err != nil {
				return nil, err
			}
			return &products, nil
		}
		err := db.Limit(productPerPage).Offset(offset).Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			return nil, err
		}
		return &products, nil
	}

	if category != "" {
		err := db.Limit(productPerPage).Offset(offset).Where("category = ?", category).Find(&products).Error
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
	name := "data:image/png;base6"
	err := db.Where("image ILIKE ?", "%"+name+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func GetAllCategories() (*[]string, error) {
	var categories []string
	err := db.Model(&Product{}).Select("category").Group("category").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}
