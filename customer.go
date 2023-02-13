package entity

import (
	"github.com/jinzhu/gorm"
)

// Customer model
type Customer struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	Address string `gorm:"type:varchar(100);not null"`
	Email   string `gorm:"type:varchar(100);not null"`
}
