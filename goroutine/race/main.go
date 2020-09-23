package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	go func() {
		a = 2
		fmt.Println("goroutine....", a)
	}()

	a = 3
	time.Sleep(1 * time.Second)
	fmt.Println("main...", a)
}
