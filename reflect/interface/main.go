package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 1.23
	// 接口类型变量得到反射类型对象
	value := reflect.ValueOf(x)

	// 反射对象得到值
	realValue := value.Interface().(float64)
	fmt.Println(realValue)

	// 反射类型对象转换为对应类型是强制转换
	pointer := reflect.ValueOf(&x)
	pointerInfo := pointer.Interface().(*float64)
	fmt.Println(pointerInfo)

}
