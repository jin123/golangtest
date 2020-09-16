package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	chann := make(chan string, 1)

	<-chann
	go func() {
		chann <- "11111"
	}()

	//test3()
}

func test1() {
	var buffer [512]byte

	n, err := os.Stdin.Read(buffer[:])
	if err != nil {

		fmt.Println("read error:", err)
		return

	}

	fmt.Println("count:", n, ", msg:", string(buffer[:]))
}

func test2() {

	reader := bufio.NewReader(os.Stdin)

	result, err := reader.ReadString('\n')
	if err != nil {

		fmt.Println("read error:", err)
	}

	fmt.Println("result:", result)
}

func test3() {

}
