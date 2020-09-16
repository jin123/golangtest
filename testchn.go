package main

import (
	"fmt"
	"time"
)

// 定义goroutine 1
func Echo(out chan<- string) { // 定义输出通道类型
	fmt.Println("1")
	//time.Sleep(1 * time.Second)
	out <- "咖啡色的羊驼"
	close(out)
}

// 定义goroutine 2
func Receive(in chan<- string, out <-chan string) { // 定义输出通道类型和输入类型
	fmt.Println("2")
	time.Sleep(1 * time.Second)
	temp := <-out // 阻塞等待echo的通道的返回
	in <- temp
	close(in)
}

func main() {
	echo := make(chan string)
	receive := make(chan string)

	go Echo(echo)
	go Receive(receive, echo)

	getStr := <-receive // 接收goroutine 2的返回

	fmt.Println(getStr)
}
