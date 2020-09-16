package main

import (
	"encoding/json"
	"fmt"
)

/**
golang 的json用法https://blog.csdn.net/wade3015/article/details/88401969
*/

type Stu struct {
	Name  string `json:"name"`
	Age   int
	HIgh  bool
	Sex   string `json:"sex"`
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

type Server struct {
	ServerName string `json:"serverName,string"`
	ServerIP   string `json:"serverIP,omitempty"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}
type Request struct {
	Data   []interface{} `json:"data"`
	Errmsg string        `json:"errmsg"`
	Errno  int           `json:"errno"`
}
type Response struct {
	Errmsg string        `json:"errmsg"`
	Data   []interface{} `json:"data"`
	Errno  int           `json:"errno"`
}

func main() {
	JsonToMapDemo()
	//beJson()
}

//json转化成结构体和map
func simpJson() {

}

//转换成json
func beJson() {

	//实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "张三",
		Age:  18,
		HIgh: true,
		Sex:  "男",
	}

	//指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	//Marshal失败时err!=nil
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	beJ := string(jsonStu)
	fmt.Println(beJ)
	return
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Guangzhou_Base", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_Base", ServerIP: "127.0.02"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_Base1", ServerIP: "192.0.02"})
	//s.Servers = append(s.Servers, stu)
	//fmt.Printf("%+v",stu)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}

	j := string(b)
	fmt.Println(j)

}
