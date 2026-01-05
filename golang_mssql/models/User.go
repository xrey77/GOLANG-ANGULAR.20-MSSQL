package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Lastname    string `gorm:"type:varchar(32)"` // Nullable if pointer type *string, otherwise defaults to ""
	Firstname   string `gorm:"type:varchar(32)"`
	Email       string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Mobile      string `gorm:"type:varchar(32)"`
	Username    string `gorm:"type:varchar(32);uniqueIndex;not null"`
	Password    string `gorm:"type:varchar(200)"`
	Isactivated bool   `gorm:"default:true"`  // Use bool for clean boolean semantics
	Isblocked   bool   `gorm:"default:false"` // Use bool
	Userpicture string `gorm:"default:http://127.0.0.1:5000/assets/users/pix.png"`
	Mailtoken   int    `gorm:"default:0"` // Use int or int32/int64
	Secret      string `gorm:"type:text;default:null"`
	Qrcodeurl   string `gorm:"type:text;default:null"`
	Role_id     int    `gorm:"type:integer"`

	Roles []Role `gorm:"many2many:user_roles;"`
}
