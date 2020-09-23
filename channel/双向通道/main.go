package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)

	data := <-ch1
	fmt.Println(data)

	ch1 <- "my name is york"
	<-done
	fmt.Println("main end")

}

func sendData(ch chan string, done chan bool) {
	ch <- "我是陈泽"

	mainData := <-ch
	fmt.Println(mainData)
	done <- true
}
