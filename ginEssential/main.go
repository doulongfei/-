package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDb()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		log.Println(name, telephone, password)

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于6位"})
			return
		}
		log.Println(name, telephone, password)

		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)

		if isTelephone(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		log.Println(name, telephone, password)
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})
	panic(r.Run())
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	//fmt.Println("hhhhh")
}

func isTelephone(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

/*
生成10位的随机字符串
*/
func RandomString(i int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, i)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func InitDb() *gorm.DB {
	driverName := "mysql"
	db, err := gorm.Open(driverName, "root:123456@(127.0.0.1:3306)/ginessential?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	//创建数据表
	db.AutoMigrate(&User{})
	return db
}
