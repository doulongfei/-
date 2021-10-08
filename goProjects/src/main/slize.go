package main

import "fmt"

func main() {
	//切片的 数组连续片段的引用
	//切片的数据结构分为：①底层数组的指针②切片的长度③切片的容量
	var intarr = [6]int{1, 2, 3, 4, 5, 6}

	slice := intarr[1:3]
	fmt.Print(slice)
	slice1 := intarr[1:3]
	fmt.Println("intarr", intarr)

	fmt.Println("slice1", slice1)
	fmt.Println("slice的元素个数：", len(slice1))

	fmt.Println("slice的容量：", cap(slice1))

	fmt.Printf("数组下标为1的地址:%p", &intarr[1])
	fmt.Printf("切片下标为0的地址:%p", &slice1[0])

	slice1[1] = 99
	fmt.Println(intarr, "数组")
	fmt.Println(slice1, "切片")

	fmt.Println("===============")
	//	切片的定义  make底层创建一个数组，对外不可见，只能通过make定义的间接的去操作
	slice11 := make([]int, 4, 20)
	fmt.Println(slice11)

	fmt.Println("长度：", len(slice11))
	fmt.Println("容量", cap(slice11))
	slice11[0] = 66
	slice11[1] = 88
	fmt.Println(slice11)

	//	方式三
	slice111 := []int{1, 4, 7, 8, 3}
	fmt.Println(slice111)
	fmt.Println("长度：", len(slice111))
	fmt.Println("容量", cap(slice111))
	//	切片的遍历  for循环
	//	切片追加 原理：创建新数组 然后赋值
	slice1112 := slice111[1:3]
	ints := append(slice1112, 88, 99)
	fmt.Println(ints)
	//	切片追加个切片
	ints = append(ints, slice1112...) //...必须写
	fmt.Println(ints)

	//切片的拷贝
	var a = []int{1, 4, 7, 3, 6, 9}
	var b = make([]int, 10)
	copy(b, a)
	fmt.Println(b)

}
