package main

import "fmt"

//type Books struct {
//	Title string
//	Author string
//	Subject string
//	ID int
//}
//
////结构体作为函数参数
//func printBook(book Books)  {
//	fmt.Println(book.ID)
//	fmt.Println(book.Title)
//	fmt.Println(book.Author)
//	fmt.Println(book.Subject)
//
//}
//func main()  {
//	defer fmt.Println("结构体")
//	var book1 Books
//	book1.ID = 1
//	book1.Title = "go"
//	book1.Author = "哈哈"
//	book1.Subject = "go语言教程"
//
//	//fmt.Println(book1.ID)
//	//fmt.Println(book1.Title)
//	//fmt.Println(book1.Author)
//	//fmt.Println(book1.Subject)
//
//	//结构提作为函数参数
//	printBook(book1)
//
//
//}

type People struct {
	name string
	age int
	hobby string
}

func printPeople(people *People)  {
	fmt.Println("指针地址",&people)
	fmt.Println(people.name)
	fmt.Println(people.age)
	fmt.Println(people.hobby)
}

func main()  {
	fmt.Println("结构体指针")
	var man People
	var woman People
	man.name = "男的"
	man.age = 20
	man.hobby = "篮球"

	woman.name = "女的"
	woman.age = 20
	woman.hobby = "足球"

	printPeople(&man)
	printPeople(&woman)
}

