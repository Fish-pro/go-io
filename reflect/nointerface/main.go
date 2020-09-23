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
	fmt.Printf("name:%s,age:%d,sex:%s", p.Name, p.Age, p.Sex)
}

func GetMsg(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("type:", getType.Name())
	fmt.Println("kind:", getType.Kind())

	getValue := reflect.ValueOf(input)
	p := getValue.Interface().(Person)
	p.PrintInfo()

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		fmt.Println(field.Name, field.Type, getValue.Field(i).Interface()) // 获取第一个数值
	}

	fmt.Println("---------------")
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Println(method.Name, method.Type)
	}
}

func main() {
	p1 := Person{"york", 20, "man"}
	GetMsg(p1)
}
