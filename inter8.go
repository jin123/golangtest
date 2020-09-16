package main

import "fmt"

type annimal interface {
	play()
	eat()
	sleep()
}

type cat interface {
	annimal
}

type test1 struct {
	cat
}

func (t test1) eat() {
	fmt.Println("i can eat")
}

func main() {
	var cat annimal = test1{}
	cat.eat()
}
