package models

import (
	"gorm.io/gorm"
	"time"
)

type Sale struct {
	gorm.Model
	Amount float64
	Date   time.Time `gorm:"type:datetimeoffset"`
}
