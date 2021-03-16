package main

import (
	"fmt"
	"log_agent/conf"
	"log_agent/es"
	"log_agent/kafaka"
	"log_agent/logtailf"
	"sync"

	"gopkg.in/ini.v1"
)

//Wg 定义等待组
var Wg sync.WaitGroup

func main() {
	//1.初始化配置文件
	
	//2.动态从etcd获取配置信息
	

	//3.初始化kafaka连接
	

	//4.初始化logtail
	
	//5.初始化es
	
	//6.发送消息到kafaka
	
	//7.读取kafaka的消息，并写入es

}
