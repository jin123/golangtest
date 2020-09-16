// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i string
	//	k int64
	j int64
}
type Coordinate struct {
	X, Y float32
}

/**
机构体
1 %v    值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
2 %#v    值的Go语法表示
3 %T    值的类型的Go语法表示
4 %%    百分号
布尔值：
5 %t    单词true或false
整数：
复制代码
1 %b    表示为二进制
2 %c    该值对应的unicode码值
3 %d    表示为十进制
4 %o    表示为八进制
5 %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
6 %x    表示为十六进制，使用a-f
7 %X    表示为十六进制，使用A-F
8 %U    表示为Unicode格式：U+1234，等价于"U+%04X"
*/

//打印坐标
func (coo *Coordinate) GetCoordinate() {
	fmt.Printf("(%.2f,%.2f)\n", coo.X, coo.Y)
	return
}
func main() {
	/*a := make([]int, 10, 10)
	fmt.Println(a)
	b := [...]string{"USA", "China", "India", "Germany", "France"}
	c := b // a copy of a is assigned to b
	c[0] = "Singapore"
	fmt.Println("b is ", b)
	fmt.Println("c is ", c)
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	var b []int = a[3:] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	fmt.Println(cap(b))
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素,容量为 4 个元素
	newSlice := slice[1:3:3]
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60，会改变底层数组中的元素
	newSlice = append(newSlice, 60)
	fmt.Println(slice, newSlice);
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50, 60, 70}
	// 创建一个新切片
	// 其长度与容量相同
	newSlice := slice[1:3:5] // 注意这里
	//fmt.Println(newSlice)
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60，会改变底层数组中的元素
	newSlice = append(newSlice, 80)
	newSlice = append(newSlice, 90)
	fmt.Println(slice)
	fmt.Println(newSlice)
	// newSlice 的底层数组已经不是 slice 了，这个改变不会影响 slice
	//newSlice[0] = 0
	//fmt.Println(slice, newSlice, cap(newSlice))*/
	//p0 := Coordinate{1, 2}
	//给该结构体p2变量分配内存，它返回指向已分配内存的指针
	/*p1 := Coordinate{X: 1, Y: 2}
	p2 := new(Coordinate)
	p2.X = 1
	p2.Y = 2
	p3 := &Coordinate{X: 1, Y: 2}
	p4 := &Coordinate{1, 2}
	//fmt.Println("-------输出p0-------")
	//fmt.Printf("%v\n%T\n", p0, p0)
	//fmt.Println("-------输出p0-1-------")
	//fmt.Printf("%+v\n%T\n", p0, p0)
	fmt.Println("-------输出p1-------")
	fmt.Printf("%+v\n%T\n", p1, p1)
	fmt.Println("-------输出p2-------")
	fmt.Printf("%+v\n%T\n", p2, p2)
	fmt.Println("-------输出p3-------")
	fmt.Printf("%+v\n%T\n", p3, p3)
	fmt.Println("-------输出p4-------")
	fmt.Printf("%+v\n%T\n", p4, p4)
	fmt.Printf("%b", 5)*/
	/*s := struct {
		a int64
		b int64
		c int64
		d int64
	}{0, 0, 0, 0}
	//	fmt.Printf("%+v:", s)
	// 将结构体指针转换为通用指针
	//p := unsafe.Pointer(&s)

	//up1 := unsafe.Offsetof(s.d)
	fmt.Println("up a :", unsafe.Offsetof(s.a))
	fmt.Println("up b :", unsafe.Offsetof(s.b))
	fmt.Println("up c :", unsafe.Offsetof(s.c))
	fmt.Println("up d :", unsafe.Offsetof(s.d))*/
	n := Num{i: "abc", j: 653662099}

	nPointer := unsafe.Pointer(&n) //它表示任意类型且可寻址的指针值，可以在不同的指针类型之间进行转换
	/**
	  修改 n.i 值：i 为第一个成员变量。因此不需要进行偏移量计算，直接取出指针后转换为 Pointer，再强制转换为字符串类型的指针值即可

	  修改 n.j 值：j 为第二个成员变量。需要进行偏移量计算，才可以对其内存地址进行修改。在进行了偏移运算后，当前地址已经指向第二个成员变量。接着重复转换赋值即可

	  作者：EDDYCJY
	  链接：https://www.jianshu.com/p/fb25dce48317
	  来源：简书
	  著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	*/
	//fmt.Println(n.j)
	niPointer := (*string)(unsafe.Pointer(nPointer)) //获取前面结构体字符串的部分
	//ptr := uintptr(nPointer)
	//fmt.Println("nPointer:", nPointer)
	//fmt.Println("ptr:", ptr)
	//fmt.Println("njPointer:", niPointer)
	//fmt.Println("第一次：", niPointer)
	*niPointer = "煎鱼"
	//fmt.Println("第二次：", *niPointer)
	//fmt.Println(njPointer)
	//unsafe.Offsetof：吧指针移动到指定位置
	//fmt.Println("test:", unsafe.Offsetof(n.j))
	//uintptr：uintptr 是 Go 的内置类型。返回无符号整数，可存储一个完整的地址。后续常用于指针运算

	/*fmt.Println("nPointer:", nPointer)
	fmt.Println("uintptr:", uintptr(nPointer))
	fmt.Println("unsafe", unsafe.Pointer(uintptr(nPointer)))*/

	/**
	*结构体的成员变量在内存存储上是一段连续的内存
	*结构体的初始地址就是第一个成员变量的内存地址
	*基于结构体的成员地址去计算偏移量。就能够得出其他成员变量的内存地址
	*修改 n.i 值：i 为第一个成员变量。因此不需要进行偏移量计算，直接取出指针后转换为 Pointer，再强制转换为字符串类型的指针值即可

	*修改 n.j 值：j 为第二个成员变量。需要进行偏移量计算，才可以对其内存地址进行修改。在进行了偏移运算后，当前地址已经指向第二个成员变量。接着重复转换赋值即可
	 */
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))

	//fmt.Println("njPointer:", unsafe.Pointer(uintptr(nPointer)))
	//njPointer2 := (unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	//fmt.Println("njPointer2:", njPointer2)
	//*njPointer2 = 3
	*njPointer = 2
	//fmt.Println("*:", njPointer)
	//fmt.Println(njPointer)
	fmt.Println("n.i: ", n.i)
	fmt.Println("n.j", n.j)

}
