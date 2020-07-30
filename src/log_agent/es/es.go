package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
	err    error
)

// Init es初始化连接信息
func Init(esurl string) {
	client, err = elastic.NewClient(elastic.SetURL(esurl))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")
}

// PutEs 数据写入es
func PutEs(data string) {
	put1, err := client.Index().Index("user").Type("go").BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Println("写入es失败", err)
		return
	}
	fmt.Println("es 写入成功")
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
