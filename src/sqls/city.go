package sqls

import (
	models "../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sd
)

// QueryCityList s
func QueryCityList() []models.City {
	ls := []models.City{}
	db, err := gorm.Open("mysql", "root:abcd1234@tcp(10.211.55.4:3306)/JHome?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.Limit(10).Offset(0).Find(&ls)
	return ls
}
