package main

import "fmt"


//interface 是一种类型 interface是一组method的集合
type phone interface {
	call()
}

type Nokia struct {
	
}

func (nokia Nokia)call()  {
	fmt.Println("i am nokia,i can call you")
}

type Iphone struct {
	
}

func (iphone Iphone) call() {
	fmt.Println("i am iphone,i can call you")
}


//定义一个Sayer接口
type Sayer interface {
	say()
}

type dog struct {
	
}

type cat struct {
	
}

func (d dog) say()  {
	fmt.Println("汪汪汪")
}

func (c cat) say()  {
	fmt.Println("喵喵喵")
}
func main()  {
	fmt.Println("接口")
	var phone phone
	phone = new(Nokia)
	phone.call()
	phone = new(Iphone)
	phone.call()

	var sayer Sayer
	a :=cat{} //实例化一个cat
	b :=dog{} //实例化一个dog
	sayer = a
	sayer.say()
	sayer = b
	sayer.say()


}
