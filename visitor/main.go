package main

import "fmt"

func main() {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	err := fn(info, nil)
	return err
}

type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%sn", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

// other visitor
type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%sn", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

// log visitor
type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}
