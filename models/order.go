package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not null; type:varchar(50)" json:"customer_name"`
	Items        *[]Item   `json:"items"`
	OrderedAt    time.Time `json:"ordered_at"`
}
