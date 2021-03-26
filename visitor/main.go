package main

import "fmt"

func main() {
	e := new(Element)
	e.Print(new(ProductionVisitor))
	e.Print(new(TestingVisitor))
}

type ProductionVisitor struct {
}

func (v ProductionVisitor) Visit() {
	fmt.Println("prod")
}

type TestingVisitor struct {
}

func (t TestingVisitor) Visit() {
	fmt.Println("test")
}

type IVisitor interface {
	Visit() // 访问者的访问方法
}

type Element struct{}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

func (el Element) Print(visitor IVisitor) {
	el.Accept(visitor)
}
