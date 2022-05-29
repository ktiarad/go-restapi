package params

type ItemCreate struct {
	ID          uint   `gorm:"primaryKey; unique" json:"item_id"`
	ItemCode    string `gorm:"not null; type:varchar(20)" json:"item_code"`
	Description string `gorm:"not null; type:varchar(50)" json:"description"`
	Quantity    int    `gorm:"not null; type:int" json:"quantity"`
	OrderID     uint
}
