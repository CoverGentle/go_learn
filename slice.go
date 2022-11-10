package main

import "fmt"

var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func main() {
	var slice0 []int = arr[2:5]
	fmt.Println(slice0, "slice0", cap(arr))
	var slice1 []int = arr[:8]
	fmt.Println(slice1, "slice1")
	var slice2 []int = arr[3:]
	fmt.Println(slice2, "slice2")
	var slice3 []int = arr[:]
	fmt.Println(slice3, "slice3")
	var slice4 []int = arr[:len(arr)-1]
	fmt.Println(slice4, "slice4")
	arr2 := [...]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	slice5 := arr2[3:5]
	slice6 := arr[3:5]
	fmt.Println(slice5, "slice5")
	fmt.Println(slice6, "slice6")
}
