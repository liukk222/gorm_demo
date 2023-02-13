package dao

import (
	"gorm_demo/entity"

	"github.com/jinzhu/gorm"
)

// CustomerDAOImpl customer data access object implementation
type CustomerDAOImpl struct{}

// CreateCustomer method to create a new customer
func (dao CustomerDAOImpl) CreateCustomer(db *gorm.DB, customer entity.Customer) (entity.Customer, error) {
	err := db.Create(&customer).Error
	return customer, err
}

// GetCustomer method to get a customer by ID
func (dao CustomerDAOImpl) GetCustomer(db *gorm.DB, id uint) (entity.Customer, error) {
	var customer entity.Customer
	err := db.First(&customer, id).Error
	return customer, err
}

// GetAllCustomers method to get all customers
func (dao CustomerDAOImpl) GetAllCustomers(db *gorm.DB) ([]entity.Customer, error) {
	var customers []entity.Customer
	err := db.Find(&customers).Error
	return customers, err
}

// UpdateCustomer method to update a customer
func (dao CustomerDAOImpl) UpdateCustomer(db *gorm.DB, customer entity.Customer) (entity.Customer, error) {
	err := db.Save(&customer).Error
	return customer, err
}

func (dao *CustomerDAOImpl) DeleteCustomer(db *gorm.DB, customer entity.Customer) error {
	if err := db.Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}
