package etcd

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	config clientv3.Config
	err    error
	client *clientv3.Client
)

//初始化etcd 连接信息
func Init() {
	//配置
	config = clientv3.Config{
		Endpoints:   []string{"132.232.11.147:2382"},
		DialTimeout: time.Second * 5,
	}
	//连接
	client, err = clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client = client
	fmt.Println("etcd 连接成功")

}

//从etcd获取key信息
func Getkey(key string) {
	kv := clientv3.NewKV(client)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	resp, err := kv.Get(ctx, key) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil {
		panic(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
