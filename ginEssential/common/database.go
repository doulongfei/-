package common

import (
	"github.com/jinzhu/gorm"
	"oceanlearn.teach/ginessential/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	driverName := "mysql"
	db, err := gorm.Open(driverName, "root:123456@(127.0.0.1:3306)/ginessential?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	//创建数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
