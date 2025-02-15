package model

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description"`
	UserID      uint
	DescPicture string
	Tasks       []Task `gorm:"foreignKey:list_id"`
}
