package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Category       string          `gorm:"type:varchar(255)"`
	Descriptions   string          `gorm:"type:varchar(255);uniqueIndex;not:null"`
	Qty            int             `gorm:"type:integer;default:0"`
	Unit           string          `gorm:"type:varchar(255)"`
	Costprice      decimal.Decimal `gorm:"type:numeric(10,2);default:0.00"`
	Sellprice      decimal.Decimal `gorm:"type:numeric(10,2);default:0.00"`
	Saleprice      decimal.Decimal `gorm:"type:numeric(10,2);default:0.00"`
	Productpicture string          `gorm:"type:varchar(255);default:null"`
	Alertstocks    int             `gorm:"type:integer;default:0"`
	Criticalstocks int             `gorm:"type:integer;default:0"`
}
