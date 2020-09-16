package main

import "fmt"


//整形
type III int
func (i III) print() {
    fmt.Println("this is int !!!! ",i)
}

//函数
func Test() {
	fmt.Println("this is test")
}

type FunT func()

func (f FunT) echo() {
	fmt.Println("this is echo ")
}
func (f FunT) Print() {
	//	fmt.Println(f)
	f.echo()
}
func main() {
	var f FunT
	//fmt.Println("this is f:", f)
	f = func(){
		fmt.Println(" iam func")
	}
	f()
	//fmt.Println(f)
	f.Print()

	var f2 III =9
	f2.print()
}
