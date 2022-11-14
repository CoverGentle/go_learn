package main

import "fmt"
func main()  {
	defer fmt.Println("定义切片")
	var slice1 []int
	fmt.Println(slice1)
	var slice2 []int = make([]int, 4,6)
	fmt.Println(slice2)
	slice3 :=make([]int,5,8)
	slice3 =  append(slice3, 8)
	slice3 = append(slice3,9,19,19,20,38,39,80,88,99,00)
	fmt.Println(slice3)
	//切片初始化
	slice4 := []int{1,2,3}
	fmt.Println(slice4)
	s1 := slice4[:]
	fmt.Println(s1)
	s2 :=slice4[0:len(slice4)-1]
	fmt.Println(s2)
	s3 :=slice4[0:]
	fmt.Println(s3)
	s4 :=slice4[:len(slice4)-0]
	fmt.Println(s4)
	var numbers = make([]int,3,5)
	printNumber(numbers)

}
func printNumber(x []int)  {
	var words []int
	if words == nil {
		fmt.Println("切片是空的")
		fmt.Println(words,len(words),cap(words))
	}

	fmt.Println("切片的len And cap",len(x),cap(x),x)
}
