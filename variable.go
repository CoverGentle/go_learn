package main

import "fmt"

func main() {
	fmt.Println("变量声明")
	//变量声明
	//1.指定变量类型，声明后若不赋值，使用默认值
	var a int = 1
	fmt.Println(a)
	//2.根据值自行判断变量类型
	var b = "helloworld"
	fmt.Println(b)
	//3.省略var，注意:=左侧的变量不应该是已经声明过的，否则会报错
	c := true
	fmt.Println(c)

	//多变量声明
	fmt.Println("多变量声明")
	//类型相同多个变量
	var name1, name2, name3 string
	name1, name2, name3 = "小红", "小明", "小白"
	fmt.Println(name1, name2, name3)
	//自行推断
	var d1, d2, d3 = 1, "你好", true
	fmt.Println(d1, d2, d3)
	//类型不同的多个变量
	var (
		str1 int    = 100
		str2 string = "张三"
		str3 bool   = true
	)
	fmt.Println(str1, str2, str3)

	//值类型和引用类型
	var num int = 11
	fmt.Println(&num) //打印内存地址
}
