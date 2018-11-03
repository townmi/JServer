package models

// Order s
type Order struct {
	ID        string  `gorm:"column:id;type:varchar(50);primary_key"`
	CreatedAt int64   `gorm:"column:created_at;"`
	PaidAt    int64   `gorm:"column:paid_at;"`
	UpdatedAt int64   `gorm:"column:updated_at;"`
	GoodsID   string  `gorm:"column:goods_id;type:varchar(50)"`
	UserID    string  `gorm:"column:user_id;type:varchar(50)"`
	Price     float64 `gorm:"column:price"`
	Status    int     `gorm:"column:status;type:int(10);"`
}

// TableName s
func (Order) TableName() string {
	return "orders"
}
