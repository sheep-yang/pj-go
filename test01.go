package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	a := A{10, "4"}
	fmt.Println(a.test01(1))
	cpuNum := runtime.NumCPU()
	for i := 1; i < cpuNum; i++ {
		go func(i int) {
			fmt.Println("cpu num is %d", i)
		}(i)
	}
	time.Sleep(1 * 1e9)
}

type A struct {
	a1 int
	a2 string
}

func (a A) test01(age int) string {
	if a.a1 == age {
		fmt.Println("age is equal")
		return " "
	} else {
		return a.a2
	}
}
