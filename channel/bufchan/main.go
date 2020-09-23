package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		非缓冲通道：make(chan type)
			一次发送，一次接受，都是阻塞的
		缓冲通道：make(chan type,size)
			发送数据：缓冲去满了才会阻塞
			接受数据：缓冲区空了才回阻塞
	*/
	// 无缓冲区
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	//ch1 <- 100 //阻塞

	// 有缓冲区
	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 100
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 200
	ch2 <- 300
	ch2 <- 400
	ch2 <- 500
	fmt.Println(len(ch2), cap(ch2))
	//ch2 <- 600 //阻塞-->死锁

	fmt.Println("--------------------")
	ch3 := make(chan string, 4)
	go sendData(ch3)

	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("read end")
			break
		}
		fmt.Println(v)
	}
	fmt.Println("main end")

}

func sendData(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "data-" + strconv.Itoa(i)
		fmt.Printf("goroutine--send:%d\n", i)
	}
	close(ch)
}
