package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"oceanlearn.teach/ginessential/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	driverName := "mysql"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	_ = fmt.Sprint("%s:%s@(%s:%s)/%s?charset=%smb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	//db, err := gorm.Open(driverName, args)
	db, err := gorm.Open(driverName, "root:123456@(127.0.0.1:3306)"+
		"/ginessential?charset=utf8mb4&parseTime=True&loc=Local")
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
