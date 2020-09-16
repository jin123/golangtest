package main

import (
	"fmt"
)

/**
经常会见到: p . *p , &p 三个符号
p是一个指针变量的名字，表示此指针变量指向的内存地址，
如果使用%p来输出的话，它将是一个16进制数。
而*p表示此指针指向的内存地址中存放的内容，一般是一个和指针类型一致的变量或者常量。
而我们知道，&是取地址运算符，&p就是取指针p的地址。等会，怎么又来了个地址，它到底和p有什么区别？区别在于，指针p同时也是个变量，既然是变量，编译器肯定要为其分配内存地址，就像程序中定义了一个int型的变量i，编译器要为其分配一块内存空间一样。而&p就表示编译器为变量p分配的内存地址，而因为p是一个指针变量，这种特殊的身份注定了它要指向另外一个内存地址，程序员按照程序的需要让它指向一个内存地址，这个它指向的内存地址就用p表示。而且，p指向的地址中的内容就用*p表示
————————————————
版权声明：本文为CSDN博主「liudashuang2017」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/liudashuang2017/article/details/79950267
*/
type Point struct {
	X int
	Y int
}

func (p Point) Print() {
	fmt.Println(p.X, p.Y)
}
func (p *Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}

type interface1 interface {
	display()
	displayAppend(string)
	displayAppend2(string) string
	displayAppend3(string)
}

type interface2 interface {
	displayAppend(string)
}

type Test struct {
	s string
}

//--------------Test struct-------------
func (t Test) display() {

	fmt.Printf("this is display:%p, %#v\n", t, t)
}

func (t *Test) displayAppend(s string) {
	fmt.Println(t)
	t.s += s
	fmt.Printf("this is displayAppend :%p, %#v\n", t, t)
}

func (t *Test) displayAppend2(s string) string {
	//fmt.Println("t.s2:", t.s)
	t.s += s
	////fmt.Println(t)
	fmt.Printf("this is DisplayAppend2 : %p, %#v\n", &t, t)

	return t.s
}

func (t Test) displayAppend3(s string) {
	//fmt.Println("t.s3:", t.s)
	t.s += s
	fmt.Printf("this is DisplayAppend3  %p, %#v\n", &t, t)
	//return t.s
}

func TestInterface(t interface1) {
	t.display()
	t.displayAppend(" AND   second  1111111  ")
	//fmt.Printf("this is after  %p, %#v\n", t, t)
	//t.displayAppend2("  AND  second again")

	//t.displayAppend3(" AND   third")
}

func main() {

	test := &Test{s: "first"}
	TestInterface(test)
	//fmt.Printf("this is :%p, %#v\n", *test, test)
	var test2 interface{} = &Test{s: "first"}

	//var test3 interface{} = Test{s: "first"}
	//test.displayAppend3(" this is 3 ")
	app, ok := test2.(interface1)

	//app2, ok2 := test3.(Test)
	fmt.Println(app, ok)
	//fmt.Println(app2, ok2)
	fmt.Println("调用display")
	app.display()
	//fmt.Println(app, test2)
	//fmt.Println(test.(interface1))
	/*fmt.Println("下面是其他的数据：")

	p := Point{1, 1}
	ptr := &p
	p.ScaleBy(2)
	//fmt.Println(p)
	p.Print()
	p.ScaleBy(2)
	p.Print()
	ptr.Print()*/
	/*ptr := &p
	    p.Print()   //1. 正确
		ptr.Print() //2. 正确

		fmt.Printf("ptr: %+v",ptr)
	    p.ScaleBy(2)      //3. 正确
	    ptr.ScaleBy(2)    //4. 正确
	    Point{1,1}.Print()    //5. 正确
	    (&Point{1,1}).Print() //6. 正确
	    (&Point{1,1}).ScaleBy( 2) //7. 正确*/
}
