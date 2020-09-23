package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10

var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go saleTicket("售票口1")
	go saleTicket("售票口2")
	go saleTicket("售票口3")
	go saleTicket("售票口4")
	//time.Sleep(5*time.Second)

	fmt.Println("wait start")
	wg.Wait()
	fmt.Println("main end")
}

func saleTicket(name string) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		// 上锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出票一张--->", ticket)
			ticket--
		} else {
			// 条件不满足，也要解锁
			mutex.Unlock()
			fmt.Println("售罄")
			break
		}
		mutex.Unlock()
	}
}
