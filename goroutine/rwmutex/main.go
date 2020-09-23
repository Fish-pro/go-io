package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	//wg.Add(2)
	// 多个同时读取
	//go readData(1)
	//go readData(2)

	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main end........")

}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "fun2 start........")
	rwMutex.Lock()
	time.Sleep(3 * time.Second)
	fmt.Println(i, "fun2 running........")
	rwMutex.Unlock()
	fmt.Println(i, "fun2 end........")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "fun1 start........")
	rwMutex.RLock()
	time.Sleep(3 * time.Second)
	fmt.Println(i, "fun1 running........")
	rwMutex.RUnlock()
	fmt.Println(i, "fun1 end........")
}
