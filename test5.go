package main

import "fmt"

func main() {
	callme(myFunc)
}

func myFunc() {
	fmt.Println("我被别人反调用")
}

func callme(f func()) {
	callme2(typeFunc(f))
}

func callme2(c Call) {
	c.call()
}

//定义的type函数
type typeFunc func()

//fun实现的Call接口的call()函数
func (f typeFunc) call() {
	f()
	fmt.Println("typeFunc call")
}

//接口
type Call interface {
	call()
	//call1()
	//myFunc()
}
