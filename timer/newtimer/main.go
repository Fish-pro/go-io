package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		创建一个计时器
	*/
	//timer := time.NewTimer(3*time.Second)
	//fmt.Printf("%T\n",timer)
	//fmt.Println(time.Now())
	//
	//ch := timer.C
	//fmt.Println(<-ch)

	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 strat")
	}()

	time.Sleep(3 * time.Second)
	flag := timer2.Stop()
	if flag {
		fmt.Println("timer 2 stop")
	}

}
