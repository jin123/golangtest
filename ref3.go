package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func SetName(p person, name string) {
	p.Name = name
}

func SetNameByPointer(p *person, name string) {
	p.Name = name
}

func (p person) SetPersonName(name string) {
	p.Name = name
}

func (p *person) SetPersonNameByPointer(name string) {
	p.Name = name
}

func main() {
	p := new(person)
	p.Name = "Tom"
	p.Age = 12
	fmt.Println(p) // &{Tom 12}

	//SetName(*p, "Jerry")
	//fmt.Println(p) // &{Tom 12}

	SetNameByPointer(p, "Jerry")
	fmt.Println(p) // &{Jerry 12}

	p.SetPersonName("Tom")
	fmt.Println(p) // &{Jerry 12}

	p.SetPersonNameByPointer("Tom")
	fmt.Println(p) // &{Tom 12}
}
