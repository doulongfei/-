package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"oceanlearn.teach/ginessential/common"
)

func main() {
	db := common.InitDb()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	//r.POST("/api/auth/register", controller.Register)
	panic(r.Run())
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	//fmt.Println("hhhhh")
}
