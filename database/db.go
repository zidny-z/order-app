package database

import (
	"errors"
	"order-app/models"

	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d Database) CreateOrder(order models.Order) (models.Order, error) {
	if err := d.db.Create(&order).Error; err != nil {
		return models.Order{}, err
	}

	newOrder := models.Order{
		OrderID:      order.OrderID,
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Items:        order.Items,
	}

	return newOrder, nil
}

func (d Database) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := d.db.Find(&orders).Error; err != nil {
		return []models.Order{}, err
	}

	return orders, nil
}

func (d Database) UpdateOrder(id int, data models.Order) (models.Order, error, bool) {
	result := d.db.Where("order_id", id).Omit("Items").Updates(&data)

	if result.RowsAffected == 0 && result.Error == nil {
		return models.Order{}, errors.New("order not found"), false
	}

	if result.Error != nil {
		return models.Order{}, result.Error, true
	}

	order := models.Order{OrderID: id}
	for _, v := range data.Items {
		if v.ItemID != 0 {
			d.db.Where("item_id", v.ItemID).Updates(&v)
		}
	}

	if err := d.db.Model(&order).Association("Items").Replace(data.Items); err != nil {
		return models.Order{}, err, true
	}

	updatedOrder := models.Order{
		OrderID:      id,
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
		Items:        data.Items,
	}

	return updatedOrder, nil, true
}

func (d Database) DeleteOrder(id int) (error, bool) {
	result := d.db.Delete(&models.Order{}, "order_id", id)

	if result.RowsAffected == 0 && result.Error == nil {
		return errors.New("order not found"), false
	}

	return result.Error, true
}