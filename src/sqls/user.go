package sqls

import (
	"errors"

	models "../models"
	utils "../utils"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sd
)

// RegisterUser s
func RegisterUser(user map[string]string) error {
	db, err := gorm.Open(utils.GetDataBaseConnection())
	if err != nil {
		// return error
		return err
	}
	defer db.Close()
	db.LogMode(true)
	count := 0
	db.Model(&models.User{}).Where("Email = ?", user["email"]).Count(&count)
	if count != 0 {
		return errors.New("this email is registed")
	}
	p, err := utils.Encode(user["password"])
	if err != nil {
		return errors.New("password is not safety")
	}
	u := models.User{ID: uuid.New().String(), Email: user["email"], Password: p}
	db.Create(&u)
	return nil
}

// ValidateUser s
func ValidateUser(user map[string]string) (*models.User, error) {
	db, err := gorm.Open(utils.GetDataBaseConnection())
	if err != nil {
		// return error
		return nil, err
	}
	defer db.Close()
	db.LogMode(true)

	p, err := utils.Encode(user["password"])
	if err != nil {
		return nil, errors.New("password is not safety")
	}

	u := models.User{}
	db.Where("Email = ? AND password = ?", user["email"], p).Find(&u)
	if u.ID == "" {
		return nil, errors.New("account is not right!")
	}
	return &u, nil
}
