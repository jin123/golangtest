package main

import (
	"fmt"
	"time"
)

func DelayPrint() {
	time.Sleep(250 * time.Millisecond)
	for i := 1; i <= 4; i++ {

		fmt.Println(i)
	}
}

func HelloWorld() {
	//time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hello world goroutine")
}

func main() {

	go DelayPrint() // 开启第一个goroutine
	go HelloWorld() // 开启第二个goroutine
	time.Sleep(2 * time.Second)
	fmt.Println("main function")

}
