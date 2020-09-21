package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1.打开文件
	fileName1 := "/Users/york/go/src/github.com/Fish-pro/myBlog/io/a.txt"
	file, err := os.Open(fileName1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 3.关闭文件
	defer file.Close()

	// 2.读取文件
	bs := make([]byte, 4, 4)
	/*n, err := file.Read(bs)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	n1, err := file.Read(bs)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(n1)
	fmt.Println(bs)
	fmt.Println(string(bs))

	n2, err := file.Read(bs)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(n2)
	fmt.Println(bs)
	fmt.Println(string(bs))

	n3, err := file.Read(bs)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(n3)
	fmt.Println(bs)
	fmt.Println(string(bs))
	*/

	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("end")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
