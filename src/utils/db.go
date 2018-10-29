package utils

import (
	"fmt"

	models "../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sd
)

// CheckDataBase s
func CheckDataBase() {
	// CREATE DATABASE IF NOT EXISTS JHome DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
	db, err := gorm.Open("mysql", "root:abcd1234@tcp(192.168.80.134:3306)/JHome?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	isCityTableExist := db.HasTable("cities")
	if !isCityTableExist {
		db.CreateTable(&models.City{})
		fmt.Println("cities table created")
	}
	isProvinceExist := db.HasTable("provinces")
	if !isProvinceExist {
		db.CreateTable(&models.Province{})
		fmt.Println("provinces table created")
	}
	isAreaExist := db.HasTable("areas")
	if !isAreaExist {
		db.CreateTable(&models.Area{})
		fmt.Println("areas table created")
	}
	isUserExist := db.HasTable("users")
	if !isUserExist {
		db.CreateTable(&models.User{})
		fmt.Println("users table created")
	}
	isGoodsExist := db.HasTable("goods")
	if !isGoodsExist {
		db.CreateTable(&models.Goods{})
		fmt.Println("goods table created")
	}
}
