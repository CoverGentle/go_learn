package main

import "fmt"

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

func main()  {
	fmt.Println("接口")
	var phone phone
	phone = new(Nokia)
	phone.call()
	phone = new(Iphone)
	phone.call()
}
