package main

import (
	"fmt"
	"time"
)

var captis = 5
var ch = make(chan struct{}, captis)

func main() {
	fmt.Println("this is a benging")
	go testBucket()
	for i := 0; i < 5; i++ {
		go getBucket()
	}
	select {}
}
func testBucket() {
	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case timeNow := <-ticker.C:
			select {
			case ch <- struct{}{}:
				//	default:
			}
			fmt.Println("this is a nums", len(ch), timeNow)
		}
	}

}

func getBucket() {
	for {
		time.Sleep(5 * time.Millisecond)
		select {
		case <-ch:
			fmt.Println("get ch")
		default:
			fmt.Println("no get ch")
		}
	}
}
