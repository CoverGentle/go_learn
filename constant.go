package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("常量-----const identifier [type] = value")
	//多个相同类型的声明
	const str1, str2, str3 = "常量1", "常量2", "常量3"
	fmt.Println(str1, str2, str3)
	//应用举例
	const width int = 10
	const height int = 20
	var area int
	area = width * height
	fmt.Println("面积为 :", area)

	//常量用作枚举
	const (
		unknown = 0
		female  = 1
		male    = 2
	)
	fmt.Println(unknown, female, male)

	//常量可以用len(),cap(),unsafe.Sizeof()常量计算表达式的值，常量表达式中，函数必须是内置函数
	const (
		a = "abcdefg"
		b = len(a)
		c = unsafe.Sizeof(a) //计算占用的字节数
	)
	fmt.Println(a, b, c)

	//特殊常量iota
	//iota可以用作枚举值
	const (
		a1 = iota //0
		b1        //1
		c1        //2
		d  = "ha" //独立值，iota += 1
		e         //"ha"   iota += 1
		f  = 100  //iota +=1
		g         //100  iota +=1
		h  = iota //7,恢复计数
		i         //8
	)
	fmt.Println(a1, b1, c1, d, e, f, g, h, i)

	const (
		n = 100
		m
	)
	fmt.Println(n, m)

}
