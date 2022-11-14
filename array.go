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


	//访问数组元素
	var n [10]int
	//var i ,j int
	for i := 0; i < 10; i++ {
		n[i] = i+100
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j,n[j])
	}
}
