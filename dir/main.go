package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileDir := "/Users/york/go/src/github.com/Fish-pro/myBlog/"

	ReadAllfile(fileDir, 0)

}

func ReadAllfile(srcFile string, level int) {
	s := "|——"
	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileList, _ := ioutil.ReadDir(srcFile)
	for _, obj := range fileList {
		fmt.Printf("%s%s\n", s, obj.Name())
		if obj.IsDir() {
			ReadAllfile(obj.Name(), level+1)
		} else {

		}
	}
}

func HandlerError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
