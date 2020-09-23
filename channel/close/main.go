package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		close(ch):关闭通道
			goroutine 写入通道10个数据，每次写一个，阻塞一次
		    main 读取数据，每次读取一个阻塞一次
	*/
	ch1 := make(chan int)
	go sendData(ch1)

	for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch1
		if !ok {
			fmt.Println("read end:", ok, v)
			break
		}
		fmt.Println("read------", v, ok)
	}
	fmt.Println("main end")
}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
