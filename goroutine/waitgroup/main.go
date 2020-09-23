package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go fun1()
	go fun2()

	fmt.Println("开始阻塞等待")
	wg.Wait()
	fmt.Println("main end")
}

func fun1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func1------", i)
	}
	wg.Done() // counter-- wg.Add(-1)
}

func fun2() {
	defer wg.Done()
	for j := 1; j <= 10; j++ {
		fmt.Println("func2------", j)
	}
}
