package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//fileName := "/Users/york/go/src/github.com/Fish-pro/myBlog/ioutil/a.txt"
	//data, err := ioutil.ReadFile(fileName)
	//HandlerError(err)
	//fmt.Println(string(data))

	// WriteFile
	//fileName := "/Users/york/go/src/github.com/Fish-pro/myBlog/ioutil/b.txt"
	//s1 := "hello word"
	//err := ioutil.WriteFile(fileName, []byte(s1), os.ModePerm)
	//HandlerError(err)

	// ReadAll()
	//s2 := "你好，我是陈泽"
	//r := strings.NewReader(s2)
	//data, err := ioutil.ReadAll(r)
	//HandlerError(err)
	//fmt.Println(string(data))

	// ReadDir() 读取子目录和子文件，只能读取一层
	//dir := "/Users/york/go/src/github.com/Fish-pro/"
	//fileInfos, err := ioutil.ReadDir(dir)
	//HandlerError(err)
	//for i := 0; i < len(fileInfos); i++ {
	//	//fmt.Println(fileInfos[i])
	//	fmt.Printf("第%d个文件，文件名：%s,是否是文件：%t\n", i, fileInfos[i].Name(), fileInfos[i].IsDir())
	//}

	// TempDir()
	dir, err := ioutil.TempDir("/Users/york/go/src/github.com/Fish-pro/myBlog/ioutil/", "test")
	HandlerError(err)
	fmt.Println(dir)
	defer os.Remove(dir)

	file, err := ioutil.TempFile(dir, "test")
	HandlerError(err)
	fmt.Println(file.Name())
	defer os.Remove(file.Name())
}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
