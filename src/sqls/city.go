package sqls

import (
	models "../models"
	utils "../utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sd
)

// QueryCityList s
func QueryCityList() []models.City {
	ls := []models.City{}
	db, err := gorm.Open(utils.GetDataBaseConnection())
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	db.Limit(10).Offset(0).Find(&ls)
	return ls
}
