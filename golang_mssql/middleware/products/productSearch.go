package middleware

import (
	"math"
	"src/golang_mssql/config"
	"src/golang_mssql/dto"
	"src/golang_mssql/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductSearch(c *gin.Context) {
	page := c.Param("page")
	key := c.Param("key")

	search := "%" + key + "%"

	perPage := 5
	db := config.Connection()

	var products []models.Product
	result := db.Where("descriptions LIKE ?", search).Find(&products)

	if len(products) == 0 {
		c.JSON(400, gin.H{"message": "Product(s) not found."})
		return
	}

	totrecs := result.RowsAffected
	total1 := float64(totrecs) / float64(perPage)
	totalPages := math.Ceil(total1)
	pg, _ := strconv.Atoi(page)
	offset := (pg - 1) * perPage

	var prods []dto.Products
	// return db.Where("name LIKE ? OR description LIKE ?", searchTermWithWildcards, searchTermWithWildcards)

	db.Where("descriptions LIKE ?", search).Limit(perPage).Offset(offset).Find(&prods)

	c.JSON(200, gin.H{
		"page":         page,
		"totpage":      totalPages,
		"totalrecords": totrecs,
		"products":     prods,
	})

}
