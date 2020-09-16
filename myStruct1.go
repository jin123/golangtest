package main

import "fmt"

type MyTest1 struct {
	name string
	age  int
}

func (m MyTest1) say() {
	fmt.Println("say hello~~~~")
}
func (m *MyTest1) SetName(name string) {
	m.name = name
	//fmt.Println("say hello~~~~")
}

func (m *MyTest1) SetAge(age int) {
	m.age = age
	//fmt.Println("say hello~~~~")
}

func main() {
	a := MyTest1{"jack", 10}
	b := &MyTest1{"tom", 20}
	a.SetName("11111")
	fmt.Println("a=", a)
	b.SetName("222222")
	fmt.Println("b=", *b)
}
