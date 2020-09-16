package main

import "fmt"

/**
使用值接收者和使用指针接受者实现接口的区别？
1. 使用值接受者实现接口，结构体类型和结构体指针类型的变量都能存
2. 指针接收者实现接口只能存结构体指针类型的变量
*/
type animal interface {
	eat()
}

type Dog struct {
	Name string
	Age  int
}

func (d Dog) eat() {
	fmt.Println(d.Name, d.Age)
}

func main() {
	var a animal
	dPoint := &Dog{
		Name: "susan1",
		Age:  12,
	}
	dStruct := Dog{
		Name: "susan2",
		Age:  13,
	}
	a = dPoint
	a.eat()
	a = dStruct
	a.eat()
	//fmt.Println("dStruct_a=", a)
}
