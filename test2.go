package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{1, 2, 3, 4}
		fmt.Println("s:",s)
	// 将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	// 保存结构体的地址备用（偏移量为 0）
	up0 := uintptr(p)
	// 将通用指针转换为 byte 型指针
	pa := (*byte)(p)
	//fmt.Println("*pb:",*pb)
	// 给转换后的指针赋值
	*pa = 10
	// 结构体内容跟着改变
	fmt.Println("sa:",s)

	// 偏移到第 2 个字段
	up := up0 + unsafe.Offsetof(s.b)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb := (*byte)(p)
	// 给转换后的指针赋值
	*pb = 20
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 3 个字段
	up = up0 + unsafe.Offsetof(s.c)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pc := (*byte)(p)
	// 给转换后的指针赋值
	*pc = 30
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 4 个字段
	up = up0 + unsafe.Offsetof(s.d)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 int64 型指针
	pd := (*int64)(p)
	// 给转换后的指针赋值
	*pd = 40
	// 结构体内容跟着改变
	fmt.Println(s)
}
