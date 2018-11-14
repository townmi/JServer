package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Goods s
type Goods struct {
	ID           string     `gorm:"column:id;type:varchar(50);primary_key"`
	GoodsName    string     `gorm:"column:name;"`
	GoodsSKU     string     `gorm:"column:sku"`
	GoodsDesc    string     `gorm:"column:desc"`
	GoodsPicture string     `gorm:"column:picture"`
	GoodsPrice   float64    `gorm:"column:price"`
	CreatedAt    int64      `gorm:"column:created_at;"`
	UpdatedAt    int64      `gorm:"column:updated_at;"`
	GoodsTypeID  string     `gorm:"column:type_id;"`
	GoodsType    GoodsType  `gorm:"ForeignKey:type_id;"`
	GoodsBrandID string     `gorm:"column:brand_id;"`
	GoodsBrand   GoodsBrand `gorm:"ForeignKey:brand_id;"`
}

// GoodsResType s
type GoodsResType struct {
	ID           string     `gorm:"column:id;" json:"id"`
	GoodsName    string     `gorm:"column:name;" json:"name"`
	GoodsSKU     string     `gorm:"column:sku;" json:"sku"`
	GoodsDesc    string     `gorm:"column:desc;" json:"desc"`
	GoodsPicture string     `gorm:"column:picture;" json:"picture"`
	GoodsPrice   float64    `gorm:"column:price;" json:"price"`
	GoodsTypeID  string     `gorm:"column:type_id;" json:"typeId"`
	GoodsType    GoodsType  `gorm:"ForeignKey:type_id;" json:"type"`
	GoodsBrandID string     `gorm:"column:brandId;"`
	GoodsBrand   GoodsBrand `gorm:"ForeignKey:brand;"`
}

// GoodsPicture s
type GoodsPicture struct {
	ID        int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	GoodsID   string `gorm:"column:goods_id;type:varchar(50)"`
	URL       string `gorm:"column:url"`
	CreatedAt int64  `gorm:"column:created_at;"`
	UpdatedAt int64  `gorm:"column:updated_at;"`
}

// GoodsType s
type GoodsType struct {
	ID        string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Name      string `gorm:"column:name;" json:"name"`
	ParentID  string `gorm:"column:parent_id;" json:"parentId"`
	Picture   string `gorm:"column:picture" json:"picture"`
	CreatedAt int64  `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;" json:"updatedAt"`
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
func (GoodsType) TableName() string {
	return "goods_types"
}

// BeforeCreate s
func (goodsType *GoodsType) BeforeCreate(scope *gorm.Scope) error {
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
