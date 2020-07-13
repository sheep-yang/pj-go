package do_redis

import (
	"fmt"
	"redis_demo/db"

	"github.com/garyburd/redigo/redis"
)

// SelectRedis 根据key查询对应的value
func SelectRedis(key string) (data string, err error) {
	data, err = redis.String(db.Conn.Do("get", key))
	if err != nil {
		fmt.Println("key is not exits:", err)
		return
	}
	return
	//	fmt.Println(data)
}

func insertRedis() {
	_, err := db.Conn.Do("Set", "name", "admin123")
	if err != nil {
		fmt.Println("redis.write err=", err)
		return
	}
	fmt.Println("写入成功")
}
