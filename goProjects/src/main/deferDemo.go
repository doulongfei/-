package main

import "fmt"
func add(num1 int ,num2 int ) int {
	// defer关键字，不会立即执行defer后面的语句，而是将defer后面的语句压入一个栈中，然后继续执行后面的语句
	// 在函数执行完毕后，从栈中取出语句开始执行，按照栈的先进后出的规则执行。
	defer fmt.Println("num1",num1)

	defer fmt.Println("num2",num2)

	var sum int =num1+num2
	fmt.Println("sum=",sum)
	return sum

	/*
	sum= 50
	num2 30
	num1 20
	50
	*/
}
func main() {
	fmt.Println(add(20,30))
	fmt.Println(add1(20,30))
}
func add1(num1 int ,num2 int ) int {
	// defer关键字，不会立即执行defer后面的语句，而是将defer后面的语句压入一个栈中，然后继续执行后面的语句
	// 在函数执行完毕后，从栈中取出语句开始执行，按照栈的先进后出的规则执行。
	defer fmt.Println("num1",num1)

	defer fmt.Println("num2",num2)

	num1+=10
	num2+=10
	var sum int =num1+num2
	fmt.Println("sum=",sum)
	return sum

	/*
	sum= 50
	num2 30
	num1 20
	50
	*/
}