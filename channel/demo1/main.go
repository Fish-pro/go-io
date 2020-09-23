package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("goroutine......")
		//time.Sleep(3*time.Second)
		data := <-ch1
		fmt.Println(data)
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 10
	fmt.Println("main end")

	// 死锁
	ch := make(chan int)
	ch <- 100

}
