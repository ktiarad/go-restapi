package params

type OrderCreate struct {
	ID           uint         `gorm:"primaryKey; not null" json:"order_id"`
	CustomerName string       `gorm:"not null; type:varchar(50)" json:"customer_name"`
	Items        []ItemCreate `gorm:"omitempty" json:"items"`
}

type OrderUpdate struct {
	CustomerName string       `gorm:"not null; type:varchar(50)" json:"customer_name"`
	Items        []ItemUpdate `gorm:"omitempty" json:"items"`
}
