package params

import "go-restapi/models"

type ItemCreate struct {
	// ID          uint   `gorm:"primaryKey; unique" json:"item_id"`
	ItemCode    string `gorm:"not null; type:varchar(20)" json:"item_code"`
	Description string `gorm:"not null; type:varchar(50)" json:"description"`
	Quantity    int    `gorm:"not null; type:int" json:"quantity"`
	OrderID     uint
}

type ItemsCreate struct {
	Items []models.Item
}

type ItemUpdate struct {
	ID          uint   `gorm:"primaryKey; unique" json:"item_id"`
	ItemCode    string `gorm:"not null; type:varchar(20)" json:"item_code"`
	Description string `gorm:"not null; type:varchar(50)" json:"description"`
	Quantity    int    `gorm:"not null; type:int" json:"quantity"`
}
