package repositories

import (
	"go-restapi/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	GetAllItems() (*[]models.Item, error)
	CreateItem(request *models.Item) error
	UpdateItem(request *models.Item, id uint) error
	DeleteItem(id uint) error
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{
		db: db,
	}
}

func (r *itemRepo) GetAllItems() (*[]models.Item, error) {
	var items []models.Item

	err := r.db.Find(&items).Error

	return &items, err
}

func (r *itemRepo) CreateItem(request *models.Item) error {
	err := r.db.Create(request).Error
	return err
}

func (r *itemRepo) UpdateItem(request *models.Item, id uint) error {
	var item models.Item

	err := r.db.Model(&item).Where("id = ?", id).Updates(models.Item{}).Error

	return err
}

func (r *itemRepo) DeleteItem(id uint) error {
	var item models.Item

	err := r.db.Model(&item).Where("id = ?", id).Delete(&item).Error

	return err
}
