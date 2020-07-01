package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

//InitDB初始化数据库连接信息
func InitDB(dataSourceName string) {
	var err error
	//dataSourceName = "root:1QAZ-pl,@tcp(localhost:3306)/operation?parseTime=true"
	DB, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库open failed")
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("mysql 连接成功")
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
}
