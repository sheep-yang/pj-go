package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 1024)

func main() {
	fmt.Println("this is a benging")
	go Producer(ch)
	go Cansumer(ch)
	select {}
}

//Producer 消息生产者
func Producer(ch chan<- string) {
	for {
		time.Sleep(time.Second)
		select {
		case ch <- "test chan":
		}
	}
}

//Cansumer 消费者
func Cansumer(ch <-chan string) {
	var count int
	for s := range ch {
		count++
		fmt.Println(s, count)
	}
}
