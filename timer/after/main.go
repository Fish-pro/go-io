package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch)

	fmt.Println(time.Now())
	timer2 := <-ch
	fmt.Println(timer2)
}
