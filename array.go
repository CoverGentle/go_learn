package main

import "fmt"

func main() {
	defer fmt.Println("数组")
	//声明数组
	var balance [10]float32
	fmt.Println(balance)
	//初始化数组
	var balance1 = []float32{1, 2, 3, 4, 5, 6}
	fmt.Println(balance1[2])
}
