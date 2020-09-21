package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "/Users/york/go/src/github.com/Fish-pro/myBlog/bufio/a.txt"
	file, err := os.Open(fileName)
	HandlerError(err)
	defer file.Close()

	// 创建reader对象
	//br := bufio.NewReader(file)
	// 1.Read() 高效读取
	//bs := make([]byte, 1024)
	//n, err := br.Read(bs)
	//HandlerError(err)
	//fmt.Println(n)
	//fmt.Println(string(bs[:n]))

	// 2.ReadLine()
	//data, flag, err := br.ReadLine()
	//HandlerError(err)
	//fmt.Println(string(data))
	//fmt.Println(flag)

	// 3.ReadString()
	//s, err := br.ReadString('\n')
	//HandlerError(err)
	//fmt.Println(s)
	//s1, err := br.ReadString('\n')
	//HandlerError(err)
	//fmt.Println(s1)
	//for {
	//	s, err := br.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Println("end")
	//		break
	//	}
	//	fmt.Println(s)
	//}

	// ReadBytes()
	//data, err := br.ReadBytes('\n')
	//HandlerError(err)
	//fmt.Println(string(data))

	// Scanner
	//s1 := ""
	//fmt.Scanln(&s1)
	//fmt.Println(s1)
	b1 := bufio.NewReader(os.Stdin)
	data, err := b1.ReadString('\n')
	HandlerError(err)
	fmt.Println(data)

}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
