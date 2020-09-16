package main

import (
	"fmt"
	"reflect"
)

// 薪资计算器接口
type SalaryCalculator interface {
	CalculateSalary() int
}

// 普通挖掘机员工
type Contract struct {
	empId    string
	basicpay int
}

// 有蓝翔技校证的员工
type Permanent struct {
	empId    string
	basicpay int
	jj       int // 奖金
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

//结构体转字典
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		//fmt.Println(obj1.Field(i).Name)
		//fmt.Println("***********")
		//fmt.Println(obj2.Field(i))
		fmt.Println("field_name=", obj1.Field(i).Name)
		fmt.Println("values=", obj2.Field(i))
		data[obj1.Field(i).Name] = obj2.Field(i)
	}
	return data

}

// 总开支
func totalExpense(s []SalaryCalculator) int {
	expense := 0
	maps := make([]interface{}, 0)
	for _, v := range s {
		expense = expense + v.CalculateSalary()
		datas := StructToMap(v)
		maps = append(maps, datas)
	}
	fmt.Println("maps=", maps)
	return expense
	//fmt.Printf("总开支 $%d", expense)
}

func main() {
	pemp1 := Permanent{"id1", 1, 1}
	pemp2 := Permanent{"id2", 1, 1}
	cemp1 := Contract{"id3", 1}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}

	//fmt.Println("employees=", employees, reflect.TypeOf(employees))
	total := totalExpense(employees)
	fmt.Println("total=", total)
}
