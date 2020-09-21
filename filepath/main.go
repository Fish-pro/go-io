package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	/*
		filePath
		打开文件：将当前对象和文件之间建立连接
	*/
	filePath1 := "/Users/york/go/src/github.com/Fish-pro/myBlog/filepath/a.txt"
	filePath2 := "a.txt"

	fmt.Println(filepath.IsAbs(filePath1))
	fmt.Println(filepath.IsAbs(filePath2))
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2))

	fmt.Println("获取父目录：", path.Join(filePath1, ".."))

	//err := os.Mkdir("/Users/york/go/src/github.com/Fish-pro/myBlog/filepath/bb", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("文件创建成功")
	//err := os.MkdirAll("/Users/york/go/src/github.com/Fish-pro/myBlog/filepath/cc/dd/ee", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("文件创建成功")

	//file1, err := os.Create("/Users/york/go/src/github.com/Fish-pro/myBlog/filepath/a.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(file1)

	//file2, err := os.Create("c.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(file2)

	//file3, err := os.Open("c.txt")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(file3)

	//file4, err := os.OpenFile("c.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(file4)
	//defer func() {
	//	file4.Close()
	//	fmt.Println("文件关闭成功")
	//}()

	//err := os.Remove(filePath1)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("ok")

	//err := os.Remove("/Users/york/go/src/github.com/Fish-pro/myBlog/filepath/cc/dd/ee")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println("ok")

	err := os.RemoveAll("/Users/york/go/src/github.com/Fish-pro/myBlog/filepath/cc")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("ok")

}
