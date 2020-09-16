package main

import (
	"fmt"
)

func main() {
	/*input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("你输入的是：", input.Text())
	*/
	var str string
	fmt.Scanln(&str)
	fmt.Printf("INPUT :%s\n", str)
}
