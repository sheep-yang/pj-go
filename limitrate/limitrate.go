package limitrate

import (
	"fmt"
	"time"
)

var captis = 5
var ch = make(chan struct{}, captis)

// TestBucket  创建chan
func TestBucket() {
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

//GetBucket  获取chan
func GetBucket() {
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

//Test chuangjaing
func Test() {
	fmt.Println("yangqiang")
}
