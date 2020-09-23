package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("goroutine....i:", i)
		}
		// 打印结束，写通道表示结束
		ch1 <- true
		fmt.Println("goroutine end")
	}()

	data := <-ch1
	fmt.Println(data)
	fmt.Println("main end")
}
