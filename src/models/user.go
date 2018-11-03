package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User models
type User struct {
	ID           string `gorm:"column:id;type:varchar(50);primary_key;not null;"`
	Mobile       string `gorm:"column:mobile;type:varchar(20);"`
	Email        string `gorm:"column:email;type:varchar(100);"`
	Level        int    `gorm:"column:level;type:int(10);"`
	UserName     string `gorm:"column:username;type:varchar(100);"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);"`
	DeviceID     string `gorm:"column:device_id;type:varchar(100);"`
	CityID       string `gorm:"column:city_id;type:varchar(50);"`
	WechatID     string `gorm:"column:wechat_id;type:varchar(50);"`
	CreatedAt    int64  `gorm:"column:created_at;"`
	UpdatedAt    int64  `gorm:"column:updated_at;"`
	BindMobileAt int64  `gorm:"column:bind_mobile_at;"`
	Password     string `gorm:"column:password;"`
	Points       int64  `gorm:"column:points"`
}

// BeforeCreate s
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix()*1000)
	scope.SetColumn("updated_at", time.Now().Unix()*1000)
	return nil
}

// TableName s
func (User) TableName() string {
	return "users"
}
