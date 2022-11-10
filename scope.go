package main

import "fmt"

// 声明全局变量
var d int

// 形参
func sum(x, y int) int {
	return x + y
}

func main() {
	defer fmt.Println("变量作用域")
	//局部变量
	var a, b, c int
	//初始化参数
	a = 10
	b = 20
	c = 30
	d = a + b
	fmt.Println(a, b, c, d)
	e := sum(a, b)
	fmt.Println(e)
}
