package main

import "fmt"

//struct结构体拥有类的特性，go封装继承多态  拥有面向对象的特性

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var t Teacher
	fmt.Println(t)
	t.Age = 18
	t.Name = "dou"
	t.School = "兰州财经大学"
	fmt.Println(t)

	var t1 Teacher = Teacher{
		Name:   "zhang",
		Age:    18,
		School: "lanzhoucaijingdaxue",
	}
	fmt.Print(t1)

	var t3 *Teacher = new(Teacher)
	//	t3是指针，指向的地址
	(*t3).School = "xuxiao"
	(*t3).Name = "fangfang"
	t3.Age = 20
	fmt.Println("t3", t3)
	fmt.Println("*t3", *t3)

	var t4 *Teacher = &Teacher{
		Name: "fei",
		Age:  18,
	}
	t4.School = "lanzhou财经大学"
	fmt.Println(t4)
}
