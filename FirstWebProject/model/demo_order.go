package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
type Order struct {
	gorm.Model
	Order_no string `gorm:"column:order_no" json:"order_no"`
	User_name string  `gorm:"column:user_name" json:"user_name"`
	Amount float64   `gorm:"column:amount" json:"amount"`
	Status string   `gorm:"column:status" json:"status"`
	File_url string  `gorm:"file_url" json:"file_url"`
}