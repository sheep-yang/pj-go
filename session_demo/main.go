package main

import (
	"fmt"
	"redis_demo/db"
	"redis_demo/do_redis"
)

func main() {
	//初始化redis连接信息
	db.InitRedis()
	SessionID := "ddddd"
	//newsession := session.NewRedisSession(SessionID)
	//newsession.Set(SessionID, "150")
	//newsession.Save()
	//	v, _ := newsession.Get("yangqiang")
	//	fmt.Println(v)
	data, err := do_redis.SelectRedis(SessionID)
	fmt.Println(data, err)

}
