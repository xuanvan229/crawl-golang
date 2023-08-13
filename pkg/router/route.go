package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/crawl-golang/pkg/handler"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	products := r.Group("/products")
	{
		products.GET("/", handler.GetProducts)
		//products.POST("/", handler.CreateProduct)
		//products.PUT("/:id", handler.UpdateProduct)
		//products.DELETE("/:id", handler.DeleteProduct)
	}

	return r
}
