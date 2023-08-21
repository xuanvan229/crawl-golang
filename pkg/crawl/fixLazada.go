package crawl

import (
	"github.com/xuanvan229/crawl-golang/pkg/model"
)

func FixLazada() {
	products, err := model.GetProductWrongImage()
	if err != nil {
		panic(err)
	}

	for _, product := range *products {
		image, err := GetProductImage(product.Url)
		if err != nil {
			continue
		}
		if image != "" {
			product.Image = image
			_, err := model.UpdateProduct(product)
			if err != nil {
				panic(err)
			}
		}

	}
}
