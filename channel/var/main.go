package main

import "fmt"

func main() {
	var a chan int
	fmt.Printf("%T，%v\n", a, a)
	if a == nil {
		fmt.Println("a is nil")
		a = make(chan int)
		fmt.Println(a)
	}
	test1(a)

}
func test1(ch chan int) {
	fmt.Printf("%T, %v", ch, ch)
}
