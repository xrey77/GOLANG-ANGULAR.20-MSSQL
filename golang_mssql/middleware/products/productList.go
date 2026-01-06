package middleware

import (
	"math"
	"src/golang_mssql/dbconfig"
	"src/golang_mssql/dto"
	"src/golang_mssql/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	page := c.Param("page")
	perPage := 5
	db := dbconfig.Connection()

	var products []models.Product
	result := db.Find(&products)

	totrecs := result.RowsAffected
	total1 := float64(totrecs) / float64(perPage)
	totalPages := math.Ceil(total1)
	pg, _ := strconv.Atoi(page)
	offset := (pg - 1) * perPage

	var prods []dto.Products

	db.Order("id").Limit(perPage).Offset(offset).Find(&prods)

	c.JSON(200, gin.H{
		"page":         page,
		"totpage":      totalPages,
		"totalrecords": totrecs,
		"products":     prods,
	})

}
