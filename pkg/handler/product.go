package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/crawl-golang/pkg/model"
	"math"
	"net/http"
	"strconv"
)

func GetProducts(c *gin.Context) {
	fmt.Print("check heheheheheheh")
	pageStr := c.DefaultQuery("page", "1")
	search := c.DefaultQuery("search", "")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	productCount, err := model.CountProduct(search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const productPerPage = 80
	pageCount := int(math.Ceil(float64(productCount) / float64(productPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}

	if page < 1 || page > pageCount {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset := (page - 1) * productPerPage

	products, err := model.GetAllProducts(productPerPage, offset, search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages":    pageCount,
		"products": products,
	})
}
