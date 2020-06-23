package main

import (
	"fmt"
	"log_agent/conf"
	"log_agent/kafaka"
	"log_agent/logtailf"
	"sync"

	"gopkg.in/ini.v1"
)

var Wg sync.WaitGroup

func main() {
	//1.初始化配置文件
	logconf := new(conf.LogConf)
	err := ini.MapTo(logconf, "./conf/conf.ini")
	if err != nil {
		fmt.Println("failed to load ini,err:", err)
	}

	//2.初始化kafaka连接
	kafaka.Init(logconf.Kafaka.Ipaddr, logconf.Kafaka.Topic)
	//3.初始化logtail
	logtailf.Initlog("./app.log")
	Wg.Add(1)
	go logtailf.Readlog()
	go kafaka.SendMsg()
	Wg.Wait()

}
