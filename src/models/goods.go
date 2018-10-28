package models

// Goods s
type Goods struct {
	ID           string  `gorm:"column:id;type:varchar(50);primary_key"`
	GoodsName    string  `gorm:"column:name;"`
	GoodsCode    string  `gorm:"column:goods_code"`
	GoodsDesc    string  `gorm:"column:desc"`
	GoodsPicture string  `gorm:"column:picture"`
	GoodsPrice   float64 `gorm:"column:price"`
}
