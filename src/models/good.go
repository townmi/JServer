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

type GoodPictures struct {
	ID     int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	GoodID string `gorm:"column:good_id;type:varchar(50)"`
	Url    string `gorm:"column:url"`
}

// TableName s
func (Goods) TableName() string {
	return "goods"
}

// TableName s
func (GoodPictures) TableName() string {
	return "good_pictures"
}
