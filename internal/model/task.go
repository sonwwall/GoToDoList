package model

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ListID      uint       `form:"list_id" json:"list_id" binding:"required"`
	Name        string     `form:"name" json:"name" binding:"required"`
	Description string     `form:"description" json:"description"`
	Priority    string     `form:"priority" json:"priority" gorm:"type:enum('p0', 'p1', 'p2');default:'p2'"`
	DueDate     *time.Time `form:"due_date" json:"due_date" gorm:"type:datetime;default:NULL"`
	Completed   bool       `form:"completed" json:"completed" gorm:"default:false"`
	UserID      uint
}
