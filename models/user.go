package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:150;not null;unique" json:"username"`
	Email string `gorm:"size:150;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Phone string `gorm:"size:255;not null;" json:"phone"`
	status uint8 `gorm:"default:1;" json:"status"`//1-active
}