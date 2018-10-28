package models

// City s
type City struct {
	ID         string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	CityName   string `gorm:"column:city_name;type:varchar(100);" json:"name"`
	CityCode   string `gorm:"column:city_code;type:varchar(40);" json:"code"`
	ProvinceID string `gorm:"column:province_id;type:varchar(50);" json:"provinceId"`
}

// Province s
type Province struct {
	ID           string `gorm:"column:id;type:varchar(50);primary_key"`
	ProvinceName string `gorm:"column:province_name;type:varchar(100);"`
	ProvinceCode string `gorm:"column:province_code;type:varchar(40);"`
}

// Area s
type Area struct {
	ID       string `gorm:"column:id;type:varchar(50);primary_key"`
	AreaName string `gorm:"column:area_name;type:varchar(100);"`
	AreaCode string `gorm:"column:area_code;type:varchar(40);"`
	CityID   string `gorm:"column:city_id;type:varchar(50);"`
}
