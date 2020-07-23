package limitrate

import (
	"testing"
	"time"
)

func TestGetBucket(t *testing.T) {
	go TestBucket()
	for i := 0; i < 5; i++ {
		go GetBucket()
	}
	time.Sleep(3 * time.Second)
}
