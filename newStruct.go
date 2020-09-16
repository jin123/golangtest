package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
}
type Bar struct {
}

func (b Bar) test() {
	fmt.Println("this is bar test")
}

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

func main() {

	regStruct["Foo"] = Foo{}
	regStruct["Bar"] = Bar{}
	str := "Bar"
	if regStruct[str] != nil {
		t := reflect.ValueOf(regStruct[str]).Type()
		v := reflect.New(t).Elem()
		v.test()
	}

}

func init() {
	regStruct = make(map[string]interface{})
	regStruct["Foo"] = Foo{}
	regStruct["Bar"] = Bar{}
}
