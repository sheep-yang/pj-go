package db

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn

//InitRedis 初始化redis
func InitRedis() {
	var err error
	Conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("redis conn sucess")
}
