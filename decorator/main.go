package main

import (
	"fmt"
	"time"
)

func main() {
	decorator(add)(10, 20)
}

func decorator(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("start...")
		sTime := time.Now().Unix()
		res := f(a, b)
		eTime := time.Now().Unix()
		fmt.Printf("duration:%d\n", eTime-sTime)
		fmt.Println("end...")
		return res
	}
}

func add(a, b int) int {
	time.Sleep(3 * time.Second)
	return a + b
}
