package params

type OrderCreate struct {
	// ID           uint   `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"not null; type:varchar(50)" json:"customer_name"`
	// Items        []Item
	// OrderedAt    time.Time
}

type OrderUpdate struct {
	// ID           uint   `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"not null; type:varchar(50)" json:"customer_name"`
	// Items        []Item
	// OrderedAt    time.Time
}
