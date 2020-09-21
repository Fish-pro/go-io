package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "/Users/york/go/src/github.com/Fish-pro/myBlog/seek/a.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	HandlerError(err)
	defer file.Close()

	bs := []byte{0}

	//_, err = file.Read(bs) // a
	//HandlerError(err)

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(1, 0)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(1, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd)
	file.WriteString("ABC")

}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
