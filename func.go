package main

import (
	"fmt"
	"math"
)

/**
func function_name( [parameter list] ) [return_types]{
   函数体
}
函数定义解析：

func：函数由 func 开始声明
function_name：函数名称，函数名和参数列表一起构成了函数签名。
parameter list]：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。
参数列表指定的是参数类型、顺序及参数个数。参数是可选的，也就是说函数也可以不包含参数。
return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
函数体：函数定义的代码集合。

*/

// 函数实例
func getNum(x, y int) int {
	var result int
	if x > y {
		result = x
	} else {
		result = y
	}
	return result
}

// 函数返回多个值
func swap(x, y string) (string, string) {
	return x, y
}

// 闭包
func getArea() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Circle struct {
	redis float64
}

func (r Circle) getA() float64 {
	return 3.14 * r.redis * r.redis
}

func main() {
	//defer语句
	defer fmt.Println("函数")
	//调用函数
	var num = getNum(10, 20)
	fmt.Println(num)
	//返回多个值
	str1, str2 := swap("hello", "world")
	fmt.Println(str1, str2)
	//函数作为值来传递
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(4))
	fmt.Println("-------------")

	//闭包
	getNum := getArea()
	for i := 0; i < 5; i++ {
		fmt.Println(getNum())
	}
	getNum1 := getArea()
	fmt.Println(getNum1())

	//方法
	var r Circle
	r.redis = 10
	fmt.Println("圆面积", r.getA())

}
