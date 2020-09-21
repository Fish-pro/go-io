package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "/Users/york/go/src/github.com/Fish-pro/myBlog/write/a.txt"
	// 1.打开文件
	//file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	HandlerError(err)
	// 3.关闭文件
	defer file.Close()
	// 2.写文件
	//bs := []byte{65, 66, 67, 68}
	bs := []byte{97, 98, 99, 100}
	//n, err := file.Write(bs)
	n, err := file.Write(bs[:2])
	HandlerError(err)
	fmt.Println(n)
	file.WriteString("\n")

	n1, err := file.WriteString("打发的")
	HandlerError(err)
	fmt.Println(n1)
	file.WriteString("\n")

	n2, err := file.Write([]byte("发发呆"))
	HandlerError(err)
	fmt.Println(n2)

}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
