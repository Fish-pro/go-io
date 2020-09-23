package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("say:", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("name:%s,age:%d,sex:%s\n", p.Name, p.Age, p.Sex)
}

func (p Person) Test(i, j int, k string) {
	fmt.Println("test", i, j, k)
}

func main() {
	p := Person{"york", 20, "man"}
	value := reflect.ValueOf(p)
	fmt.Println(value.Kind(), value.Type())

	methodValue1 := value.MethodByName("PrintInfo")
	fmt.Println(methodValue1.Kind(), methodValue1.Type())

	// 没有参数调用
	methodValue1.Call(nil)
	args := make([]reflect.Value, 0)
	methodValue1.Call(args)

	methodValue2 := value.MethodByName("Say")
	fmt.Println(methodValue2.Kind(), methodValue2.Type())
	args1 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args1)

	methodValue3 := value.MethodByName("Test")
	fmt.Println(methodValue3.Kind(), methodValue3.Type())
	args2 := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20), reflect.ValueOf("hello world")}
	methodValue3.Call(args2)

}
