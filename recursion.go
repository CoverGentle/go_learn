package main

import "fmt"
func main()  {
	fmt.Println("递归函数")
	//recursion()
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
}

func recursion()  {
	//var i  = 0
	//i++
	//fmt.Println("我调我自己",i)
	//recursion()

}

func Factorial(x int)(result int)  {
	if x == 0 {
		result =1
	}else{
		result = x*Factorial(x-1)
	}

	return
}
