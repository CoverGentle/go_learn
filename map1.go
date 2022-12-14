package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	defer fmt.Println("map1")
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scourMap = make(map[string]int,200)
	for i := 0; i < 10; i++ {
		key :=fmt.Sprintf("stu%02d",i) //生成stu开头的字符串
		value := rand.Intn(100) //生成0~99的随机整数
		scourMap[key] =value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string,0,200)
	for k := range scourMap {
		keys = append(keys,k)
	}
	fmt.Println(keys,"keys")
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key,scourMap[key])
	}


	//元素为map类型的切片
	var mapSlice = make([]map[string]string,3)
	fmt.Println(mapSlice)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
