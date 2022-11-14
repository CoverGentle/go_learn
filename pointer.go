package main

import "fmt"

func main() {
	defer fmt.Println("指针")
	var a int = 20
	fmt.Println("变量的地址",&a)


	var b int=20  //声明的实际变量
	var ip *int  //定义指针变量
	ip = &b  //指针变量的储存地址
	fmt.Println("b变量的地址",&b)
	fmt.Println("ip变量的储存地址",ip)
	fmt.Println("ip变量的值",*ip)

	var ptr *int
	fmt.Println(ptr)
	fmt.Println(ptr)
}
