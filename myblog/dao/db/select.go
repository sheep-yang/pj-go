package db

import "fmt"

func SelectInfo() {
	defer DB.Close()
	fmt.Println("数据库连接成功")
	//	sqlStr := "select hostname from cicdinfo LIMIT 1;"
	//	var cicdinfo cicdInfo
	//	err := DB.Get(&cicdinfo, sqlStr)
	//if err != nil {
	//	fmt.Println("查询失败")
	//	return
	//	}
	//	fmt.Println(cicdinfo.Hostname)

}
