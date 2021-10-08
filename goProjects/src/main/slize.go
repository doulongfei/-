package main

import "fmt"

func main() {
	//切面的 数组连续片段的引用
	var intarr = [6]int{1, 2, 3, 4, 5, 6}

	var slice []int= intarr[1:3]
	fmt.Print(slice)
	slice1:=intarr[1:3]
	fmt.Println("intarr",intarr)

	fmt.Println("slice1",slice1)
	fmt.Println("slice的元素个数：",len(slice1))

	fmt.Println("slice的容量：",cap(slice1))



}
