package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 10

func main() {
	go saleTicket("售票口1")
	go saleTicket("售票口2")
	go saleTicket("售票口3")
	go saleTicket("售票口4")
	time.Sleep(3 * time.Second)
}

func saleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出票一张--->", ticket)
			ticket--
		} else {
			fmt.Println("售罄")
			break
		}
	}
}
