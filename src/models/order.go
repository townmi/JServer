package models

// Order s
type Order struct {
	ID string `gorm:"column:id;type:varchar(50);primary_key"`
}
