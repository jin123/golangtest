package main

import "fmt"

func main() {
	a := [...]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	//fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			fmt.Println("a[i]=", a[i])
			fmt.Println("a[j]=", a[j])
			fmt.Println("******************")
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
	fmt.Println(a)
}
