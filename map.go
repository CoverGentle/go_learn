package main

import "fmt"
func main()  {
	defer fmt.Println("集合")
	//定义map
	//var countryCapitalMap map[string]string
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japn"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("captal of",country,"is",countryCapitalMap[country])
	}
	//判断某个键是否存在 value, ok := map[key]
	v,ok :=countryCapitalMap["Chines"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("查无此项",v)
	}

	//map遍历
	for k, v := range countryCapitalMap {
		fmt.Println("k",k,"v",v)
	}
	for key := range countryCapitalMap {
		fmt.Println("key",key)
	}

	//delete(map,key) 删除map中的键值对
	delete(countryCapitalMap,"France")
	for k1, v1 := range countryCapitalMap {
		fmt.Println(k1,v1)
	}



}
