package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23

	// 值传递修改
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()
	fmt.Println(newValue.Type())
	fmt.Println(newValue.CanSet())

	newValue.SetFloat(3.14)
	fmt.Println(num)

	//如果不是指针地址
	value := reflect.ValueOf(num)
	//value.SetFloat(30) // error
	//fmt.Println(value.CanSet())
	value.Elem() //error

}
