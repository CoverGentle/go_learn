package main

import "fmt"
/*
循环语句range
for key, value := range oldMap {
    newMap[key] = value
}
*/
func main()  {
	defer fmt.Println("Range关键字")
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums :=[]int{2,3,4}
	sum :=0
	for _,num := range nums{
		sum+=num
	}
	fmt.Println("sum",sum)
	fmt.Println("====================================")

	// 忽略 2nd value，支持 string/array/slice/map。
	s := "ABC"
	for i := range s {
		fmt.Println("",s[i])
	}
	// 忽略 index。
	for _, c := range s {
		println(c)
	}

}
