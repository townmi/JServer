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
	db, err := gorm.Open(GetDataBaseConnection())
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

	isGoodsPictureExist := db.HasTable("goods_pictures")
	if !isGoodsPictureExist {
		db.CreateTable(&models.GoodsPicture{})
		fmt.Println("goods_pictures table created")
	}

	isGoodsFirstTypeExist := db.HasTable("goods_first_types")
	if !isGoodsFirstTypeExist {
		db.CreateTable(&models.GoodsFirstType{})
		fmt.Println("goods_first_types table created")
	}

	isGoodsSecondTypeExist := db.HasTable("goods_second_types")
	if !isGoodsSecondTypeExist {
		db.CreateTable(&models.GoodsSecondType{})
		fmt.Println("goods_second_types table created")
	}

	isGoodsThirdTypeExist := db.HasTable("goods_third_types")
	if !isGoodsThirdTypeExist {
		db.CreateTable(&models.GoodsThirdType{})
		fmt.Println("goods_third_types table created")
	}

	isGoodsBrandExist := db.HasTable("goods_brands")
	if !isGoodsBrandExist {
		db.CreateTable(&models.GoodsBrand{})
		fmt.Println("goods_brands table created")
	}
}
