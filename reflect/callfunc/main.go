package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Println(value.Kind(), value.Type())

	value1 := reflect.ValueOf(fun2)
	fmt.Println(value1.Kind(), value1.Type())

	value2 := reflect.ValueOf(fun3)
	fmt.Println(value2.Kind(), value2.Type())

	// 通过反射调用函数
	value.Call(nil)
	value1.Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf("test")})

	resultValue := value2.Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf("test")})
	fmt.Printf("%T\n", resultValue)
	fmt.Println(len(resultValue))
	fmt.Println(resultValue[0].Kind(), resultValue[0].Type())

	s := resultValue[0].Interface().(string)
	fmt.Println(s)
}

func fun1() {
	fmt.Println("fun1")
}

func fun2(i int, s string) {
	fmt.Println(i, s)
}

func fun3(i int, j string) string {
	fmt.Println("fun2")
	return j + strconv.Itoa(i)
}
