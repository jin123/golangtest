package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

/*func query(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Kind ", k)
	fmt.Println("Type ", t)
	//v2 := v.Interface()
	//fmt.Println("Value ", v2.ordId)

}*/

func main() {

	formate()
	return
	/*o := order{
		ordId:      456,
		customerId: 56,
	}
	//fmt.Println(o.ordId)
	query(o)*/
}

func formate() {
	var a string = "{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
		    "list": [
			{
			    "hotel_id": "2000505",
			    "hotel_name": "汉庭上海东安路二店"
			}
		    ],
		    "count": 1
		}
	    }"
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	fmt.Println(reflect.TypeOf(valueOfA))
	// 获取interface{}类型的值, 通过类型断言转换
	var getA string = valueOfA.Interface().(string)
	// 获取64位的值, 强制类型转换为int类型
	//var getA2 int = int(valueOfA.Int())
	fmt.Println(getA, reflect.TypeOf(getA))
}
