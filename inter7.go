package main

import (
	"fmt"
)

func main() {
	user := &User{name: "Chris"}
	user.ISubUser = &NormalUser{}
	user.sayHi()
	user.ISubUser = &ArtisticUser{}
	user.sayHi()
}

type ISubUser interface {
	sayType()
	speakChinese()
}

//******************************

type User struct {
	name string
	ISubUser
}

func (u *User) sayHi() {
	u.sayName()
	u.sayType()
	u.speakChinese()
}

func (u *User) sayName() {
	fmt.Printf("I am %s.", u.name)
}

//************************

type NormalUser struct {
}

func (n *NormalUser) sayType() {
	fmt.Println("I am a normal user.")
}
func (n *NormalUser) speakChinese() {
	fmt.Println("normal user can speak chinese")
}

//*********************

type ArtisticUser struct {
}

func (a *ArtisticUser) sayType() {
	fmt.Println("I am an artistic user.")
}
func (a *ArtisticUser) speakChinese() {
	fmt.Println("artistic user can speak chinese")
}
