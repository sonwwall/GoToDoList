package model

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description"`
	GroupID     uint   `form:"group_id" json:"group_id" gorm:"default:0"`
	Tag         string `form:"tag" json:"tag"`
	UserID      uint
	DescPicture string
	Tasks       []Task `gorm:"foreignKey:list_id"`
}
