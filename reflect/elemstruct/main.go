package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s := Student{"york", 20, "云南大学"}
	p := &s
	fmt.Println(p.Name, (*p).Name, s.Name)

	value := reflect.ValueOf(p)
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name")
		f1.SetString("haha")

		f2 := newValue.FieldByName("Age")
		f2.SetInt(23)
	}
	fmt.Println(s)
}
