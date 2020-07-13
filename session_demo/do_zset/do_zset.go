package do_zset

import (
	"fmt"
	"redis_demo/db"

	"github.com/garyburd/redigo/redis"
)

func ZsetRedis() {
	//写入ZSET 数据
	_, err := db.Conn.Do("ZADD", "score", 80, "dazhaozhao", 85, "xiaoming")
	if err != nil {
		panic(err)
	}
	// 获取成员个数
	result, err := db.Conn.Do("ZCARD", "score")
	if err != nil {
		panic(err)
	}
	fmt.Printf("成员个数是:%d\n", result)

	//取出 升序
	scoreMap, err := redis.StringMap(db.Conn.Do("ZREVRANGE", "score", 0, 2, "withscores"))
	for name := range scoreMap {
		fmt.Println(name, scoreMap[name])
	}

	//取出 降序
	scoreMap, err = redis.StringMap(db.Conn.Do("ZRANGE", "score", 0, 1, "withscores"))
	for name := range scoreMap {
		fmt.Println(name, scoreMap[name])
	}
	//取出 dazhaozhao的分数
	score, err := redis.Int(db.Conn.Do("ZSCORE", "score", "dazhaozhao"))
	if err != nil {
		panic(err)
	}
	fmt.Println(score)

}
