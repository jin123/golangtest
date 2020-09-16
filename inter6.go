package main

import (
	"fmt"
	"reflect"
)

// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package.  The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	//Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	//Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	//Swap(i, j int)
}

//结构体转字典
func StructToMap(obj []interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data

}

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data []interface{}) {
	//data.(([]interface{}))
	//fmt.Println("data_type=", reflect.TypeOf(data))
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	//n := data.Len()
	//maxDepth := 0
	//objs := StructToMap(data)
	//fmt.Println("objs=", objs)
	//ret := make([]interface{}, 0)
	ret := StructToMap(data)
	fmt.Println("ret=", ret)
	/*for i := 0; i < n; i++ {
		fmt.Println("each_data=", data.(map)[i])
		for j := i + 1; j < n; j++ {
			if data.Less(j, j+1) {
				data.Swap(i, j)
			}
		}
	}*/

	/*for i := n; i > 0; i >>= 1 {
		maxDepth++
	}*/
	//maxDepth *= 2
	//quickSort(data, 0, n, maxDepth)
}

type Person struct {
	Name string
	Age  int
}

/*func (p Person) String() string {
	//fmt.Println("to striing")
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}*/

// ByAge implements sort.Interface for []Person based on
// the Age field.
//type ByAge []Person //自定义

func main() {
	/*people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
		{"jyx", 3},
		{"jrm", 30},
	}*/
	p1 := Person{"Bob1", 12}
	fmt.Sprintf("p1=%+v", p1)
	p2 := Person{"Bob2", 31}
	p3 := Person{"Bob3", 1}
	p4 := Person{"Bob4", 4}
	p := []Person{p1, p2, p3, p4}
	Sort(p)
	//fmt.Println(reflect.TypeOf(p))
	//fmt.Println("people11111=", people)
	//fmt.Println(reflect.TypeOf(people))
	//fmt.Println(reflect.TypeOf(people))
	//fmt.Println(reflect.TypeOf(ByAge(people)))
	//Sort(ByAge(people))
	//fmt.Println(p)
}
