package main

import (
	"fmt"
	"reflect"
)

type Orange struct {
	size   int    `kitty:"size"`
	Weight int    `kitty:"wgh"`
	From   string `kitty:"source"`
}

func (this Orange) GetWeight() int {
	fmt.Println("get weight")
	return this.Weight
}

func main() {

	orange := Orange{1, 18, "Shanghai"}

	refValue := reflect.ValueOf(orange) // value

	refType := reflect.TypeOf(orange) // type

	fmt.Println("orange refValue:", refValue)
	fmt.Println("orange refType:", refType)

	orangeKind := refValue.Kind() // basic type
	fmt.Println("orange Kind:", orangeKind)

	fieldCount := refValue.NumField() // field count
	fmt.Println("fieldCount:", fieldCount)

	/*for i := 0; i < fieldCount; i++ {
		fieldType := refType.Field(i)          // field type
		fieldValue := refValue.Field(i)        // field vlaue
		fieldTag := fieldType.Tag.Get("kitty") // field tag

		fmt.Println("fieldTag:", fieldTag)
		fmt.Println("field type:", fieldType.Type)
		fmt.Println("fieldValue:", fieldValue)

	}*/

	// method
	result := refValue.Method(0).Call(nil)
	fmt.Println(result)
	fmt.Println("get type result", reflect.TypeOf(result))
	fmt.Println("method result:", result[0])
}
