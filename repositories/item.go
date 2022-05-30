package repositories

import (
	"go-restapi/models"
	"log"

	"gorm.io/gorm"
)

type ItemRepo interface {
	GetAllItems() (*[]models.Item, error)
	GetItemsById(id uint) (*[]models.Item, error)
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

func (r *itemRepo) GetItemsById(id uint) (*[]models.Item, error) {
	log.Default().Printf("peeking GetItemsById, id :%d", id)

	var items []models.Item

	err := r.db.Where("order_id = ?", id).Find(&items).Error
	if &items == nil {
		log.Default().Printf("DATA NOT FOUND")
	}

	return &items, err
}

func (r *itemRepo) CreateItem(request *models.Item) error {
	err := r.db.Create(request).Error
	log.Default().Println("CreateItem at itemrepo")
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
