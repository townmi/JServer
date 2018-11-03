package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Goods s
type Goods struct {
	ID                  string  `gorm:"column:id;type:varchar(50);primary_key"`
	GoodsName           string  `gorm:"column:name;"`
	GoodsSKU            string  `gorm:"column:sku"`
	GoodsDesc           string  `gorm:"column:desc"`
	GoodsPicture        string  `gorm:"column:picture"`
	GoodsPrice          float64 `gorm:"column:price"`
	GoodsFirstTypeID    string  `gorm:"column:f_type_id;type:varchar(50)"`
	GoodsFirstTypeName  string  `gorm:"column:f_type_name;"`
	GoodsSecondTypeID   string  `gorm:"column:s_type_id;type:varchar(50)"`
	GoodsSecondTypeName string  `gorm:"column:s_type_name;"`
	GoodsThirdTypeID    string  `gorm:"column:t_type_id;type:varchar(50)"`
	GoodsThirdTypeName  string  `gorm:"column:t_type_name;"`
	GoodsBrandID        string  `gorm:"column:brand_id;type:varchar(50)"`
	GoodsBrandName      string  `gorm:"column:brand_name;"`
	CreatedAt           int64   `gorm:"column:created_at;"`
	UpdatedAt           int64   `gorm:"column:updated_at;"`
}

// GoodsPicture s
type GoodsPicture struct {
	ID      int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	GoodsID string `gorm:"column:goods_id;type:varchar(50)"`
	URL     string `gorm:"column:url"`
}

// GoodsFirstType s
type GoodsFirstType struct {
	ID      string `gorm:"column:id;type:varchar(50);primary_key"`
	Name    string `gorm:"column:name;"`
	Picture string `gorm:"column:picture"`
}

// GoodsSecondType s
type GoodsSecondType struct {
	ID      string `gorm:"column:id;type:varchar(50);primary_key"`
	Name    string `gorm:"column:name;"`
	Picture string `gorm:"column:picture"`
}

// GoodsThirdType s
type GoodsThirdType struct {
	ID      string `gorm:"column:id;type:varchar(50);primary_key"`
	Name    string `gorm:"column:name;"`
	Picture string `gorm:"column:picture"`
}

// GoodsBrand s
type GoodsBrand struct {
	ID   string `gorm:"column:id;type:varchar(50);primary_key"`
	Name string `gorm:"column:name;"`
	Logo string `gorm:"column:logo"`
}

// TableName s
func (Goods) TableName() string {
	return "goods"
}

// BeforeCreate s
func (goods *Goods) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}

// TableName s
func (GoodsPicture) TableName() string {
	return "goods_pictures"
}

// TableName s
func (GoodsFirstType) TableName() string {
	return "goods_first_types"
}

// TableName s
func (GoodsSecondType) TableName() string {
	return "goods_second_types"
}

// TableName s
func (GoodsThirdType) TableName() string {
	return "goods_third_types"
}

// TableName s
func (GoodsBrand) TableName() string {
	return "goods_brands"
}
