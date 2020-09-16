package main

import (
	"fmt"
)

type myInterface interface {
	myFunc(int)
	//	myFunc2()
}

type myStruct struct {
	i int
}

func (c myStruct) myFunc(v int) {
	c.i = v

	fmt.Println("test")
}

func main() {
	var myStruct1 interface{} = myStruct{}
	var myStruct2 interface{} = &myStruct{}
	//var myStruct2 interface{} = myStruct{}
	//use, _ := myStruct1.(myInterface)
	//use2, _ := myStruct2.(myInterface)
	fmt.Println(myStruct1, myStruct2)
	//use.myFunc(1)
	//use2.myFunc(2)

	//getType(myStruct1)
	///getType(myStruct2)
	t1 := myStruct1.(myStruct)
	t2 := myStruct2.(myInterface)
	fmt.Println(t1, t2)
	//getType(myStruct4)
}

func getType(arg interface{}) {
	switch arg.(type) {
	case int:
		fmt.Println("the type is int")
	case string:
		fmt.Println("the type is string")
	case myStruct: /*是自定义数据类型*/
		//使用.(数据类型)进行强转
		fmt.Println("this is myStruct")
	case myInterface: /*是定义的接口数据类型*/
		fmt.Println("this is myInterface")
	}
}
