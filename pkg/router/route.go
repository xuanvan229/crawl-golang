package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/crawl-golang/pkg/config"
	"github.com/xuanvan229/crawl-golang/pkg/handler"
	"net/http"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Setting.Host},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	products := r.Group("/products")
	{
		products.GET("/", handler.GetProducts)
		products.GET("/category/", handler.GetCategories)
		//products.PUT("/:id", handler.UpdateProduct)
		//products.DELETE("/:id", handler.DeleteProduct)
	}

	return r
}
