package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint
	Product    string
	Price      float32
}
