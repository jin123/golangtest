package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		fmt.Println("111111")
	}()

	go func() {
		fmt.Println("22222")
	}()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
