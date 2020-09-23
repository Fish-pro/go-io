package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		通过range获取通道的值
	*/
	ch1 := make(chan int)
	go sendData(ch1)

	for v := range ch1 { // v <- ch1
		fmt.Println("v-----", v)
	}
	fmt.Println("main end")
}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
	close(ch) // 如果不关闭，则会发生死锁
}
