package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	s = s[:110000029722]
	fmt.Println(cap(s), len(s))
}
