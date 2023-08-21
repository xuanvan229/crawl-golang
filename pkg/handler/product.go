package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/crawl-golang/pkg/model"
	"math"
	"net/http"
	"strconv"
)

func GetProducts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	search := c.DefaultQuery("search", "")
	category := c.DefaultQuery("category", "")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	productCount, err := model.CountProduct(search, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const productPerPage = 20
	pageCount := int(math.Ceil(float64(productCount) / float64(productPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}

	if page < 1 || page > pageCount {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset := (page - 1) * productPerPage

	products, err := model.GetAllProducts(productPerPage, offset, search, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages":    pageCount,
		"products": products,
	})
}

func GetCategories(c *gin.Context) {
	categories, err := model.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
