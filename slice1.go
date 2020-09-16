package main

import "fmt"

type test struct {
	a int
	b int
}

func main() {
	testRef()
}

func testSlice1() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr len", len(arr))
	fmt.Println("arr cap", cap(arr))
	var slice = arr[0:3]
	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)
	fmt.Println("slice len", len(slice))
	fmt.Println("slice cap", cap(slice))

	fmt.Println("arr2 len", len(arr))
	fmt.Println("arr2 cap", cap(arr))

	slice2 := make([]string, 5, 10)
	//slice2[0] = "zero"
	//slice2[1] = "one"
	//slice2[2] = "two"
	//slice2[3] = "three"
	//slice2[4] = "four"
	//slice2 = append(slice2, "five")
	SliceAppend(slice2, "six")
	fmt.Println("slice2 len", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))
	fmt.Println("slice2=", slice2)
	//fmt.Printf("修改前 参数p 指向的内存地址 = %p\n", slice2)
	//fmt.Printf("修改前 参数p 内存地址 = %p\n", &slice2)
	//test := test{}
	//fmt.Printf(" test 指向的内存地址 = %v\n", test)
	//fmt.Printf("test 内存地址 = %v\n", &test)
}

func testSlice2() {
	a := []string{"a", "b", "c", "d"}
	b := a
	c := a[0:3]

	d := []string{"d"}
	e := d
	d = []string{"g", "h"}

	a[1] = "f"

	fmt.Println(a) // [a f c d]
	fmt.Println(b) // [a f c d]
	fmt.Println(c) // [a f c]
	fmt.Println(d) // [g h]
	fmt.Println(e) // [d]
}

func testRef() {
	a := map[string]string{"a": "a"}
	newA := a
	for key, _ := range a {
		if key == "a" {
			a[key] = "A"
		}
	}
	a["1111"] = "111111"
	fmt.Println(a)    // map[a:A]
	fmt.Println(newA) // map[a:A]

	a = map[string]string{}
	fmt.Println(a)    // map[]
	fmt.Println(newA) // map[a:A]
}

func SliceAppend(s []string, str string) {
	s = append(s, str)
}
