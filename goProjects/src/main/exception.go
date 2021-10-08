package main

import "fmt"

func main() {
	fmt.Println("异常测试")

	test12()

	fmt.Println("异常测试222")
	//数组的定义：
	var scours [5]int
	scours[0] = 1
	scours[1] = 2
	scours[2] = 3
	scours[3] = 4
	scours[4] = 5

	//for i := 0; i < len(scours); i++ {
	//	fmt.Printf("请录入第%d个学生的成绩", i+1)
	//	fmt.Scanln(&scours[i])
	//}
	fmt.Println(scours)
	var arr1 [3]int = [3]int{3, 6, 9}
	fmt.Println(arr1)

	var arr2 = [3]int{1, 4, 7}
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	var arr4 = [...]int{3: 33, 2: 22, 5: 55}
	fmt.Println(arr4)

	//	长度属于数组的一部分   数组的类型为[n]int

	var arr5 = [...]int{1, 2, 3, 4, 5}
	//通过指针  引用传递   可实际改变值 传递地址
	test13(&arr5)
	//值传递，传递过去的数组修改值后，无法返回修改后的值
	//test13(arr5)
	fmt.Println(arr5)

	//	二维数组的定义
	var arr6 [2][3]int16

	fmt.Println(arr6)
	fmt.Printf("arr的地址是：%p \n", &arr6)
	fmt.Printf("arr[0]的地址是：%p \n", &arr6[0])
	fmt.Printf("arr[0][0]的地址是：%p \n", &arr6[0][0])
	fmt.Printf("arr[1]的地址是：%p \n", &arr6[1])
	fmt.Printf("arr[1][0]的地址是：%p \n", &arr6[1][0])

	//	初始化操作
	var arr7 = [][]int{{1, 23, 4}, {2, 5, 6}}
	fmt.Println(arr7)

}

//指针接受地址
func test13(arr *[5]int) {
	//值传递
	//func test13(arr [5]int) {
	arr[0] = 99
	fmt.Println("方法中改变的值：", arr[0])
}

func test12() {
	//利用defer+recover来捕获错误
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("错误是：", err)
		}
	}()
	num1 := 10
	num2 := 2
	result := num1 / num2
	fmt.Println(result)
}
