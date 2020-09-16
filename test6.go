package main

import (
	"fmt"
)

func testChann() {
	fmt.Println("start")
	str := make(chan string, 1)
	str <- "111"
	<-str
	/*go func() {
		fmt.Println("write")
		time.Sleep(4 * time.Second)

		str <- "1111"
		str <- "22222"
	}()
	fmt.Println(<-str)
	fmt.Println(<-str)*/
	//go in(str)
	//fmt.Println("test")
	//go out(str)

	fmt.Println("end")

}
func main() {
	testChann()

}
