package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	select {
	case <-ch1:
		fmt.Println("case1")
	case <-ch2:
		fmt.Println("case2")
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
		//default:
		//	fmt.Println("default")
	}

	fmt.Println("main end")
}
