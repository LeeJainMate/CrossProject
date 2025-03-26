package app

import (
	"Project-Go/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ค้นหาสินค้าจากคำอธิบาย (description)
func SearchByDescription(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/products/search/description")
	{
		routes.GET("/", func(c *gin.Context) {
			query := c.DefaultQuery("query", "")

			var products []model.Product
			if query != "" {
				// ค้นหาจาก description ของสินค้า
				result := db.Where("description LIKE ?", "%"+query+"%").Find(&products)
				if result.Error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
					return
				}
				if len(products) == 0 {
					c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"products": products})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
			}
		})
	}
}

// ค้นหาสินค้าจากช่วงราคา (price range)
func SearchByPrice(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/products/search/price")
	{
		routes.GET("/", func(c *gin.Context) {
			minPrice := c.DefaultQuery("min_price", "0")
			maxPrice := c.DefaultQuery("max_price", "10000")

			var products []model.Product
			// แปลงค่า minPrice และ maxPrice เป็น float64 เพื่อใช้ในการคำนวณ
			// กำหนดค่า min และ max ที่จะใช้ในการค้นหาตามช่วงราคา
			result := db.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&products)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			if len(products) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"products": products})
		})
	}
}
