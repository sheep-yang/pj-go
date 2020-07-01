package main

import (
	"myblog/controller"
	"myblog/dao/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库连接
	dataSourceName := "root:1QAZ-pl,@tcp(localhost:3306)/operation?parseTime=true"
	db.InitDB(dataSourceName)
	//设置路由信息
	//r := gin.Default()
	router := gin.New()
	router.GET("/", controller.GetIndex)
	db.SelectInfo()

}
