package main

import "fmt"
func main()  {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	var name string
	fmt.Printf("请输入姓名")
	fmt.Scanln(&name)
}