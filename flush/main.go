package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "/Users/york/go/src/github.com/Fish-pro/myBlog/flush/a.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandlerError(err)

	w1 := bufio.NewWriter(file)
	//n, err := w1.WriteString("hello world")
	//HandlerError(err)
	//fmt.Println(n)
	//w1.Flush()

	for i := 1; i <= 1000; i++ {
		w1.WriteString(fmt.Sprintf("%d--hello", i))
	}
	w1.Flush()

}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
