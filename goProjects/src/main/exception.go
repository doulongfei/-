package main

import "fmt"

func main() {
	fmt.Println("异常测试")

	test12()

	fmt.Println("异常测试222")

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
