package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	input := []byte("hello world")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println("encodeString=", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}

	destring := string(decodeBytes)
	fmt.Println("decode=", destring)
	new_input := []byte(destring)
	fmt.Println("re encode=", base64.StdEncoding.EncodeToString(new_input))
}
