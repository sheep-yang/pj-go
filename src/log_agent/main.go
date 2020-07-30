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
	logconf := new(conf.LogConf)
	err := ini.MapTo(logconf, "./conf/conf.ini")
	if err != nil {
		fmt.Println("failed to load ini,err:", err)
	}
	//2.动态从etcd获取配置信息
	//etcd.Init(logconf.Etcdinfo.Etcdipaddr)

	//3.初始化kafaka连接
	kafaka.Init(logconf.Kafaka.Ipaddr, logconf.Kafaka.Topic)

	//4.初始化logtail
	logtailf.Initlog(logconf.Path)
	//5.初始化es
	es.Init(logconf.Esinfo.Esurl)
	//6.发送消息到kafaka
	Wg.Add(1)
	go logtailf.Readlog()
	go kafaka.SendMsg()
	go kafaka.KafakaToEs()
	Wg.Wait()
	//7.读取kafaka的消息，并写入es

}
