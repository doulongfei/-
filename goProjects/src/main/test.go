package main //必须有包
import (
	"fmt"
	"unsafe"
)
//全局变量，定义在函数外
// var n=1
// var n2
var(

	name4="dou"
	sex="nan"
	age=193453
)

func main(){
	name:="doulongfei"
	var name1="zhangfangfang"
	var age1 int8=127
	//支持多变量同时声明
	fmt.Println("hello golang\n",name,"\n",name1)
	fmt.Println(name4,sex,age)
	fmt.Println(age1)
	fmt.Println(unsafe.Sizeof(age1))
	//%T变量类型
	fmt.Printf("sex的类型：%T\n",sex)
	var flag=true
	fmt.Println(flag)
	fmt.Println("=================================")
	fmt.Println(`================-----+++++++`)
	//字符串拼接  加号必须在行尾，     go语言的特性是：一行只能一条语句。
	var str="sfa"+"fafa"+"sfsf"+
	"sfasf"+"sfaf"+
	"sfsf"
	var str1 string
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println("类型转换：============")
	var swap =12
	var swapInt=string(swap)
	fmt.Println(swapInt)






}