package sqls

import (
	"errors"
	"strconv"

	models "../models"
	utils "../utils"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// CreateGoods s
func CreateGoods(goods map[string]string) error {
	db, err := gorm.Open(utils.GetDataBaseConnection())
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	count := 0
	db.Model(&models.Goods{}).Where("goods_sku = ?", goods["sku"]).Count(&count)
	if count != 0 {
		return errors.New("this goods is registed")
	}
	price, err := strconv.ParseFloat(goods["price"], 64)
	if err != nil {
		return err
	}
	g := models.Goods{
		ID:           uuid.New().String(),
		GoodsName:    goods["name"],
		GoodsSKU:     goods["sku"],
		GoodsDesc:    goods["desc"],
		GoodsPicture: goods["picture"],
		GoodsPrice:   price,
	}
	db.Create(&g)
	return nil
}
