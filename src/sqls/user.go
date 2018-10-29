package sqls

import (
	models "../models"
	utils "../utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sd
)

// RegisterUser s
func RegisterUser(user map[string]string) (bool, error) {
	db, err := gorm.Open("mysql", "root:abcd1234@tcp(10.211.55.4:3306)/JHome?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// return error
		return false, nil
	}
	defer db.Close()
	count := 0
	db.Where("Email = ?", user["email"]).Count(&count)
	if count != 0 {
		return false, nil
	}
	p, err := utils.Encode(user["password"])
	if err != nil {
		return false, err
	}
	u := models.User{Email: user["email"], Password: p}
	db.Create(&u)
	return true, nil
}
