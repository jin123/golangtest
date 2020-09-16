package main

import (
	"fmt"
)

type Tester interface {
	Display()
	DisplayAppend(string)
	DisplayAppend2(string) string
}

type Tester2 interface {
	DisplayAppend(string)
}

type Test struct {
	s string
}

func (t *Test) Display() {
	fmt.Printf(" this this display %#v\n", t)
}

func (t *Test) DisplayAppend(s string) {
	t.s += s
	fmt.Printf("DisplayAppend:%v\n", t)
	//fmt.Printf("++++DisplayAppend+++++:%p\n", t)
}

func (t *Test) DisplayAppend2(s string) string {
	t.s += s
	fmt.Printf("DisplayAppend2:%p, %v\n", t, t)
	return t.s
}

func TestInterface(t Tester) {
	t.Display()
	t.DisplayAppend(" AND append two ")
	t.DisplayAppend2("  AND append three again")
}

func TestInterface2(t Tester2) {
	t.DisplayAppend("TestInterface2")
}

func main() {
	var test Test
	test.s = " first "
	// fmt.Printf("%p\n", &test)
	//  test.Display()
	// test.DisplayAppend(" raw")

	TestInterface(&test)
	//fmt.Println("test=", test)
	//TestInterface(test) //cannot use test (type Test) as type Tester in argument to TestInterface:Test does not implement Tester (Display method has pointer receiver)

	//TestInterface2(&test)
	//

	//TestInterface2(test)
}
