package main

import "fmt"


type value interface {
	Show()
}

type test_value struct{}

// func (value *test_value) Show() { fmt.Printf("test \n") }



type Phone interface {
	call()
	//call2()
}

type NestPreparer interface {
	NestPrepare()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {
}

func (iPhone ApplePhone) call() {
	fmt.Println("I am Apple Phone, I can call you!")
}

func main() {
	var phone Phone
	phone = NokiaPhone{}
	fmt.Println(phone)
	///lePhone{}

	//phone.call()

	//var _ value = new(test_value) // 使用类似方法处理，编译的时候会报错



}
