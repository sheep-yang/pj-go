package session

import (
	"fmt"
	"redis_demo/db"
	"redis_demo/do_redis"
	"sync"
)

// Session 接口
type Session interface {
	Set(key, value interface{}) error // 设置Session
	Get(key interface{}) interface{}  // 获取Session
	Del(key interface{}) error        // 删除Session
	Save() error                      // 保存Session到redis
}

// Manager Session管理
type Manager struct {
}

//RedisSession 定义
type RedisSession struct {
	SessionID string
	Data      map[string]interface{}
	Rwlock    sync.RWMutex
}

// NewRedisSession 初始化结构体
func NewRedisSession(SessionID string) *RedisSession {
	return &RedisSession{
		SessionID: SessionID,
		Data:      make(map[string]interface{}, 24),
	}
}

//Set 设置session
func (s *RedisSession) Set(key string, value interface{}) (err error) {
	//先设置锁
	s.Rwlock.Lock()
	defer s.Rwlock.Unlock()
	s.Data[key] = value
	return
}

// Get 查询session
func (s *RedisSession) Get(key string) (value interface{}, err error) {
	s.Rwlock.Lock()
	defer s.Rwlock.Unlock()
	value, ok := s.Data[key]
	if !ok {
		fmt.Printf("this key %s is not exits\n", key)
		return
	}
	return
}

// Del 删除key
func (s *RedisSession) Del(key string) (err error) {
	s.Rwlock.Lock()
	defer s.Rwlock.Unlock()
	// 根据key删除RedisSession
	delete(s.Data, key)
	return
}

// Save 保存到REDIS
func (s *RedisSession) Save() (err error) {
	//保存之前先查询key是否存在,如果存在返回key存在，不存在就保存；
	data, _ := do_redis.SelectRedis(s.SessionID)
	fmt.Println(data)
	if data != " " {
		fmt.Println("this key is exits")
		return
	}
	fmt.Println("save is being")
	_, err = db.Conn.Do("Set", s.SessionID, s.Data)
	if err != nil {
		fmt.Println("redis.write err=", err)
		return
	}
	fmt.Println("保存到REDIS成功")
	return
}
