package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		单项通道
	*/
	ch1 := make(chan int) // 双向
	//ch2 := make(chan <- int) //单向，只能写
	//ch3 := make(<- chan int) //单向，只能读

	//ch1 <- 10
	//data := ch1
	//fmt.Println(data)

	//ch2 <- 10
	//data := <- ch2

	//ch3 <- 10
	//data := <- ch3
	//fmt.Println(data)
	wg.Add(2)
	go fun1(ch1)
	go fun2(ch1)

	//data := <- ch1
	//fmt.Println(data)

	wg.Wait()
	fmt.Println("main end")
}

func fun1(ch chan<- int) {
	defer wg.Done()
	ch <- 100
	fmt.Println("fun1 end")
}

func fun2(ch <-chan int) {
	defer wg.Done()
	data := <-ch
	fmt.Println(data)
	fmt.Println("fun2 end")
}
