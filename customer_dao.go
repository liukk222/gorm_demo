package dao

import (
	"gorm_demo/entity"

	"github.com/jinzhu/gorm"
)

// CustomerDAO interface
type CustomerDAO interface {
	CreateCustomer(db *gorm.DB, customer entity.Customer) (entity.Customer, error)
	GetCustomer(db *gorm.DB, id uint) (entity.Customer, error)
	GetAllCustomers(db *gorm.DB) ([]entity.Customer, error)
	UpdateCustomer(db *gorm.DB, customer entity.Customer) (entity.Customer, error)
	DeleteCustomer(db *gorm.DB, customer entity.Customer) error
}
