package main

import (
	"fmt"
	"time"
)

func main() {
	go printInt()
	for i := 1; i <= 100; i++ {
		fmt.Println("main-----A")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main----------end")
}

func printInt() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("int------%d\n", i)
	}
}
