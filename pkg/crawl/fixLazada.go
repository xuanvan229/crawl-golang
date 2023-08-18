package crawl

import (
	"fmt"
	"github.com/xuanvan229/crawl-golang/pkg/model"
)

func FixLazada() {
	products, err := model.GetProductWrongImage()
	if err != nil {
		panic(err)
	}

	for _, product := range *products {
		image := CrawlDetailProduct(product.Url)
		fmt.Println("image", image)
		if image != "" {
			product.Image = image
			_, err := model.UpdateProduct(product)
			if err != nil {
				panic(err)
			}
		}

	}
}
