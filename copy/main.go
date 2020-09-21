package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	srcFile := "/Users/york/go/src/github.com/Fish-pro/myBlog/copy/a.txt"
	targetFile := "/Users/york/go/src/github.com/Fish-pro/myBlog/copy/b.txt"

	//n, err := CopyFile(srcFile, targetFile)
	//HandlerError(err)
	//fmt.Println(n)

	//n1, err := CopyFile1(srcFile, targetFile)
	//HandlerError(err)
	//fmt.Println(n1)

	n2, err := CopyFile2(srcFile, targetFile)
	HandlerError(err)
	fmt.Println(n2)

	fmt.Println("main end")
}

func CopyFile2(srcFile string, targetFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(targetFile, bs, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

func CopyFile1(srcFile string, targetFile string) (int64, error) {
	src, err := os.Open(srcFile)
	HandlerError(err)
	defer src.Close()

	target, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandlerError(err)
	defer target.Close()

	return io.Copy(target, src)
}

func CopyFile(srcFile string, targetFile string) (int, error) {
	src, err := os.Open(srcFile)
	HandlerError(err)
	defer src.Close()

	target, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandlerError(err)
	defer target.Close()

	bs := make([]byte, 1024, 1024)
	total := 0
	for {
		n, err := src.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("end ")
			break
		} else if err != nil {
			return total, err
		}
		total += n
		_, err = target.Write(bs[:n])
		if err != nil {
			return total, err
		}
	}
	return total, nil

}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
