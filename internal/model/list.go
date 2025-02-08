package model

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required"`
}
