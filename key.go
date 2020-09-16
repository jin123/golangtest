package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	// var SecretKey = map[string]string{
	// 	"test":  "123456",
	// 	"crius": "3582c68cec4f9e75e7ff90868494eff0",
	// 	"sleep": "27df7ac0203c8070d551108666e358ec",
	// }
	t := "{'company_id':1125899908710318,'order_id':1125899914520626,'member_id':1125899909300282}"
	srcSign, _ := MD5(t + "&key=" + "27df7ac0203c8070d551108666e358ec")

	fmt.Println("sleep" + "_" + srcSign)
}

func MD5(str string) (string, []byte) {
	hash := md5.New()
	hash.Write([]byte(str))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), sum
}
