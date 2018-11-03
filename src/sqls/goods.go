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
	db.Model(&models.Goods{}).Where("sku = ?", goods["sku"]).Count(&count)
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

// QueryGoodsList s
func QueryGoodsList(filter map[string]string) ([]models.Goods, error) {
	db, err := gorm.Open(utils.GetDataBaseConnection())
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)

	ls := []models.Goods{}
	var limit int64
	if filter["limit"] != "" {
		l, err := strconv.ParseInt(filter["limit"], 10, 32)
		limit = l
		if err != nil {
			return nil, err
		}
	}
	if limit < 10 {
		limit = 10
	}
	var offset int64
	if filter["offset"] != "" {
		o, err := strconv.ParseInt(filter["offset"], 10, 32)
		offset = o
		if err != nil {
			return nil, err
		}
	}
	sort := filter["sort"]
	if sort == "" {
		sort = "updated_at desc"
	} else {
		if sort[0:1] == "-" {
			sort = sort[1:len(sort)]
			sort += " asc"
		} else if sort[0:1] == "+" {
			sort = sort[1:len(sort)]
			sort += " aesc"
		} else {
			sort += " desc"
		}
	}

	db.Order(sort).Limit(limit).Offset(offset).Find(&ls)
	return ls, nil
}
