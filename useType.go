package main

import "fmt"

/**
理一下思路：
使用type定义函数 func(int)
定义 Call 接口，Call中有一个函数 call(int)
在main()中调用one(2, callback)，在one()中调用two()，传入two()函数前，对callback函数实现了类型转换，从普通函数转换成type定义的函数。
在 two() 中调用传入的 c 因为 c 实现了 Call 接口，所以可以调用 call() 函数，最终调用了我们传入的 callback() 函数

作者：紫葡萄0
链接：https://www.jianshu.com/p/431abe0d2ed5
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func main() {
	one(" chinese", callback)
	//one(" english", callback2)
}

//main中调用的函数
func one(say string, f func(string)) {
	//f(say)
	//fmt.Println(f)
	two(say, typeFunc(f))
}

//one()中调用的函数
func two(say string, c Call) {

	c.call(say)
	c.call2(say)

}

//定义的type函数
type typeFunc func(string)

//fun实现的Call接口的call()函数

func (f typeFunc) call(say string) {
	fmt.Println("this is call")
	f(say) //被注释了
}
func (f typeFunc) call2(say string) {
	fmt.Println("this is call2")
	f(say)
}

//接口
type Call interface {
	//fmt.Println(int)
	call2(string)
	call(string)
}

//需要传递函数
func callback(say string) {
	fmt.Printf(" I am say %s \n", say)
	//	fmt.Println(i)
}
func callback2(say string) {
	fmt.Printf(" I am say22222 %s \n", say)
	//	fmt.Println(i)
}
