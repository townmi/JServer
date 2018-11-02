package models

// Goods s
type Goods struct {
	ID           string  `gorm:"column:id;type:varchar(50);primary_key"`
	GoodsName    string  `gorm:"column:goods_name;"`
	GoodsSKU     string  `gorm:"column:goods_sku"`
	GoodsDesc    string  `gorm:"column:desc"`
	GoodsPicture string  `gorm:"column:picture"`
	GoodsPrice   float64 `gorm:"column:price"`
}

// GoodsPictures s
type GoodsPictures struct {
	ID      int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	GoodsID string `gorm:"column:goods_id;type:varchar(50)"`
	URL     string `gorm:"column:url"`
}

// TableName s
func (Goods) TableName() string {
	return "goods"
}

// TableName s
func (GoodsPictures) TableName() string {
	return "goods_pictures"
}
