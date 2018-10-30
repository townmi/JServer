package sqls

import (
	"fmt"

	models "../models"
	utils "../utils"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sd
)

// RegisterUser s
func RegisterUser(user map[string]string) (bool, error) {
	db, err := gorm.Open(utils.GetDataBaseConnection())
	if err != nil {
		// return error
		return false, nil
	}
	defer db.Close()
	db.LogMode(true)
	count := 0
	db.Model(&models.User{}).Where("Email = ?", user["email"]).Count(&count)
	fmt.Println(count)
	if count != 0 {
		return false, nil
	}
	p, err := utils.Encode(user["password"])
	if err != nil {
		return false, err
	}
	fmt.Println(p)
	u := models.User{ID: uuid.New().String(), Email: user["email"], Password: p}
	db.Create(&u)
	return true, nil
}
