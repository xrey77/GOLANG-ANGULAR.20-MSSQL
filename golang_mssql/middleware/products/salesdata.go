package middleware

import (
	"src/golang_mssql/dbconfig"
	"src/golang_mssql/models"
)

func GetSalesData() ([]float64, []string) {
	var sales []models.Sale
	db := dbconfig.Connection()

	db.Find(&sales) // In production, use .Where() or .Select() to filter

	var values []float64
	var labels []string
	for _, s := range sales {
		values = append(values, s.Amount)
		labels = append(labels, s.Date.Format("2006-01-02"))
	}
	return values, labels
}
