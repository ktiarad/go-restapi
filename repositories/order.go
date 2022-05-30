package repositories

import (
	"go-restapi/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	GetAllOrders() (*[]models.Order, error)
	CreateOrder(request *models.Order) (uint, error)
	UpdateOrder(request *models.Order, id uint) error
	DeleteOrder(id uint) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) GetAllOrders() (*[]models.Order, error) {
	var orders []models.Order

	err := r.db.Preload("Items").Find(&orders).Error

	return &orders, err
}

func (r *orderRepo) CreateOrder(request *models.Order) (uint, error) {
	result := r.db.Create(request)
	err := result.Error
	id := request.ID
	return id, err
}

func (r *orderRepo) UpdateOrder(request *models.Order, id uint) error {
	var order models.Order

	err := r.db.Model(&order).Where("id = ?", id).Updates(models.Order{CustomerName: request.CustomerName}).Error
	return err
}

func (r *orderRepo) DeleteOrder(id uint) error {
	var order models.Order

	err := r.db.Model(&order).Where("id = ?", id).Delete(&order).Error

	return err
}
