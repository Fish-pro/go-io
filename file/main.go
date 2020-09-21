package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		fileInfo：文件信息
			interface:
				Name()

	*/
	fileInfo, err := os.Stat("/Users/york/go/src/github.com/Fish-pro/myBlog/file/a.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())

}
