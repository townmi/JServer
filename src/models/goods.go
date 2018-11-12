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

// GoodsResType s
type GoodsResType struct {
	ID           string  `gorm:"column:id;" json:"id"`
	GoodsName    string  `gorm:"column:name;" json:"name"`
	GoodsSKU     string  `gorm:"column:sku;" json:"sku"`
	GoodsDesc    string  `gorm:"column:desc;" json:"desc"`
	GoodsPicture string  `gorm:"column:picture;" json:"picture"`
	GoodsPrice   float64 `gorm:"column:price;" json:"price"`
}

// GoodsPicture s
type GoodsPicture struct {
	ID        int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	GoodsID   string `gorm:"column:goods_id;type:varchar(50)"`
	URL       string `gorm:"column:url"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
}

// GoodsFirstType s
type GoodsFirstType struct {
	ID        string `gorm:"column:id;type:varchar(50);primary_key"`
	Name      string `gorm:"column:name;"`
	Picture   string `gorm:"column:picture"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
}

// GoodsSecondType s
type GoodsSecondType struct {
	ID        string `gorm:"column:id;type:varchar(50);primary_key"`
	Name      string `gorm:"column:name;"`
	Picture   string `gorm:"column:picture"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
}

// GoodsThirdType s
type GoodsThirdType struct {
	ID        string `gorm:"column:id;type:varchar(50);primary_key"`
	Name      string `gorm:"column:name;"`
	Picture   string `gorm:"column:picture"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
}

// GoodsBrand s
type GoodsBrand struct {
	ID        string `gorm:"column:id;type:varchar(50);primary_key"`
	Name      string `gorm:"column:name;"`
	Logo      string `gorm:"column:logo"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
}

// GoodsBrandResType s
type GoodsBrandResType struct {
	ID   string `gorm:"column:id;" json:"id"`
	Name string `gorm:"column:name;" json:"name"`
	Logo string `gorm:"column:logo" json:"logo"`
}

// TableName s
func (Goods) TableName() string {
	return "goods"
}

// TableName s
func (GoodsResType) TableName() string {
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

// BeforeCreate s
func (goodsPicture *GoodsPicture) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}

// TableName s
func (GoodsFirstType) TableName() string {
	return "goods_first_types"
}

// BeforeCreate s
func (goodsFirstType *GoodsFirstType) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}

// TableName s
func (GoodsSecondType) TableName() string {
	return "goods_second_types"
}

// BeforeCreate s
func (goodsSecondType *GoodsSecondType) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}

// TableName s
func (GoodsThirdType) TableName() string {
	return "goods_third_types"
}

// BeforeCreate s
func (goodsThirdType *GoodsThirdType) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}

// TableName s
func (GoodsBrand) TableName() string {
	return "goods_brands"
}

// TableName s
func (GoodsBrandResType) TableName() string {
	return "goods_brands"
}

// BeforeCreate s
func (goodsBrand *GoodsBrand) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}
