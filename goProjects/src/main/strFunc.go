package main

import (

	"fmt"
	"strings"
	"strconv"
)
func main() {
	var str string ="hello"

	fmt.Println(len(str))

	for i, v := range str {
		fmt.Println("索引：%d,值：%d",i,v)
	}

	for _,value:=range str{
		fmt.Println("值：%d",value)
	}


	// 字符串转整数
	n,err:=strconv.Atoi("66")
	fmt.Println(n,err)

	str1:=strconv.Itoa(23434)
	fmt.Println(str1)

	// 统计一个字符串中有几个指定的字串

	count:=strings.Count("golangandjavaga","ga")

	fmt.Println(count)

	// 不区分大小写的字符串比较
	flag:=strings.EqualFold("hello","HELLO")

	fmt.Println(flag)

	// 区分大小写的字符串比较

	fmt.Println("hello"=="HELLo")
	//返回字串在字符串第一次出现的索引值，如果没有返回-1
	index:=strings.Index("javaAndgolang","l")

	fmt.Println(index)
}