package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
