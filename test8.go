
package main

import (
	"fmt"
	"time"
	"runtime"
)
func init(){
	fmt.Println("init")
}
func main(){
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
go sayHello()
go sayWorld()
time.Sleep(1 * time.Second)
}

func sayHello(){
fmt.Print("hello-")
runtime.Gosched()
}


func sayWorld(){
	//runtime.Gosched()
	fmt.Println("world !!!!")
	
}