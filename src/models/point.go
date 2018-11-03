package models

// Point s
type Point struct {
	ID        string `gorm:"column:id;type:varchar(50);primary_key"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
	UserID    string `gorm:"column:user_id;type:varchar(50)"`
	Status    int    `gorm:"column:status;type:int(10);"`
}

// TableName s
func (Point) TableName() string {
	return "points"
}
