package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	// 获取逻辑cpu数量
	fmt.Println("cpu------->", runtime.NumCPU())

	// 设置逻辑cpu执行go程序的最大数量 [1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

}

func main() {
	// 获取goroot目录
	fmt.Println("goroot----->", runtime.GOROOT())

	// 获取操作系统
	fmt.Println("os/platform---->", runtime.GOOS)

	//goshed
	//go func() {
	//	for i:=1;i<=5;i++{
	//		fmt.Println("goroutine...")
	//	}
	//}()
	//
	//for i:=1;i<=4;i++{
	//	// 让出cpu执行时间片
	//	runtime.Gosched()
	//	fmt.Println("main...")
	//}

	// 启动一个goroutine
	go func() {
		fmt.Println("start............")
		fun()
		fmt.Println("end............")
	}()

	time.Sleep(3 * time.Second)

}

func fun() {
	defer fmt.Println("defer........")
	//终止函数
	//return
	runtime.Goexit()
	fmt.Println("run................")
}
