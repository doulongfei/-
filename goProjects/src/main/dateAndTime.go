package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("测试")
	fmt.Printf("%v,对应的类型为：%T", now, now)
	fmt.Println(now.YearDay())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	str := fmt.Sprintln("接受Springt-- 可用于变量接受")
	fmt.Println(str)
	//格式化输出时间中的字符串必须写2006年01月02日 15:04:05  中间的分隔符可以自定义写
	timeStr := now.Format("2006年01月02日 15:04:05")
	fmt.Println(timeStr)

	//new分配内存 用于基本数据类型，struct
	num := new(int)
	fmt.Printf("num的类型：%T,num的值：%v,num的地址：%v,num指针指向的值：%v", num, num, &num, *num)

}
func name(num1, num2 int) int {
	return 90
}
