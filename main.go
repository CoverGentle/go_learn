package main

import "fmt"

var (
	a int
	b bool
	c string
)

func GetData() (int, int) {
	return 10, 20
}

// 结构体
type Books struct {
	title  string
	author string
}

func main() {
	//str := "nihao"
	//m, _ := GetData()
	//_, n := GetData()
	//var a int = 32
	//fmt.Println(a, b, c, str, "helloworld")
	//fmt.Println(m, n)

	////数组
	//var balance [10]float32
	//fmt.Println(balance)
	//var strArr = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//fmt.Println(strArr, len(strArr))

	////	结构体
	//var book Books
	//book.title = "曹操"
	//book.author = "易中天"
	//fmt.Println(book)
	//fmt.Println(Books{"活着", "余华"})

	//切片
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	//2.:=
	s2 := []int{}
	//3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	//4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	//5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr[1:4] //前包后不包
	fmt.Println(s6)
}
