package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	srcFile := "/Users/york/go/src/github.com/Fish-pro/myBlog/temp/a.jpg"
	targetFile := "/Users/york/go/src/github.com/Fish-pro/myBlog/temp/b.jpg"

	lastName := srcFile[strings.LastIndex(srcFile, "/")+1:]
	tempFile := lastName + "temp.txt"
	fmt.Println(lastName)
	fmt.Println(tempFile)

	file1, err := os.Open(srcFile)
	HandlerError(err)
	file2, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandlerError(err)
	file3, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandlerError(err)

	defer func() {
		file1.Close()
		file2.Close()
		file3.Close()
	}()

	// 1.读取临时文件中的数据
	bs := make([]byte, 100, 100)
	n, err := file2.Read(bs)
	//HandlerError(err)
	countStr := string(bs[:n])
	count, err := strconv.ParseInt(countStr, 10, 64)
	//HandlerError(err)
	fmt.Println(count)

	// 2.根据count值设置偏移量
	file1.Seek(count, io.SeekStart)
	file3.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	n1 := -1
	n2 := -1
	total := int(count)

	// 3.复制文件

	for {
		n1, err = file1.Read(data)
		if n1 == 0 || err == io.EOF {
			fmt.Println("over:", total)
			break
		}
		n2, err = file3.Write(data[:n1])

		// 记录复制的数据量
		total += n2
		file2.Seek(0, io.SeekStart)
		file2.WriteString(strconv.Itoa(total))

		//fmt.Println("total:%d\n", total)
		//if total >= 8000 {
		//	panic("self error")
		//}
	}

}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
