package dao

import (
	"gorm_demo/entity"

	"github.com/jinzhu/gorm"
)

type OrderDAOImpl struct{}

func (odao *OrderDAOImpl) FindAllOrders(db *gorm.DB) ([]entity.Order, error) {
	var orders []entity.Order
	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (odao *OrderDAOImpl) FindOrderByID(db *gorm.DB, id int) (entity.Order, error) {
	var order entity.Order
	if err := db.First(&order, id).Error; err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (odao *OrderDAOImpl) CreateOrder(db *gorm.DB, order entity.Order) error {
	if err := db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (odao *OrderDAOImpl) UpdateOrder(db *gorm.DB, order entity.Order) error {
	if err := db.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

func (odao *OrderDAOImpl) DeleteOrder(db *gorm.DB, order entity.Order) error {
	if err := db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}
