package main
import "fmt"

func sum(num1 int,num2 int) int { //如果没有返回值，则参数后面的返回值类型不写，即可。类似于void方法
	return num1+num2
}

func sum1(num1 int,num2 int) (int ,int) { //如果没有返回值，则参数后面的返回值类型不写，即可。类似于void方法
	return num1+num2,num1*num2
}
func main() {
	var str string="hello golang 你好"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n",str[i])
	}

	for i, v := range str {
		fmt.Printf("索引为：%d,具体的值为：%c \n", i,v)
	}

	//双重循环   内部循环可结束外面的循环
	label:    //此语句的作用用于第二层循环可在适当的时候结束第一层循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i:  %v,j:  %v \n", i,j)
			if i==2&&j==2 {
				break label
			}
		}
	}
	fmt.Println("hello golang 你好1")
	fmt.Println("hello golang 你好2")
	fmt.Println("hello golang 你好3")
	goto helloLabel
	fmt.Println("hello golang 你好4")
	fmt.Println("hello golang 你好5")
	fmt.Println("hello golang 你好6")
	helloLabel:
	fmt.Println("hello golang 你好7")
	fmt.Println("hello golang 你好8")
	fmt.Println("hello golang 你好9")
	
	res:=sum(20,30)
	fmt.Println(res)
	res1,res2:=sum1(20,30)
	res3,_:=sum1(20,30)    //_  :表示忽略第二个值
	fmt.Println(res1,res2)
	fmt.Println(res3)


	//go语言中没有函数重载

	num:=10
	fmt.Println(num)
	test(&num)
	fmt.Println(num)

	//函数也是一种数据类型，可以赋值给一个变量，可以通过改变量，调用函数

	a := test1

	fmt.Printf("a的类型是：%T \n", a)
	a(10)

	test2(1)
	test2(1,2,3,4,5)

	test3(1,3.14,test1)

}
func test(num *int)  {
	*num=30
	fmt.Println(num)
	fmt.Println("传递过来的值：",*num)
}
func test1(num int)  {
	fmt.Println(num)
}
//可变类型参数
func test2(args...int)  {
	for _, v := range args {
		fmt.Println(v)
	}
}
//给func(int) 起了一个别名，后面可以直接使用
type myFunc func(int)
func test4(num1 int,num2 float32,testFunc myFunc)  {
	fmt.Println("调用函数类型的参数")
	fmt.Println(testFunc)
}
func test3(num1 int,num2 float32,testFunc func(int))  {
	fmt.Println("调用函数类型的参数")
	fmt.Println(testFunc)
}
//求两个数的和、差
func test5(num1 int,num2 int) (int ,int) {
	res1:=num1+num2
	res2:=num1-num2
	return res1,res2
}
//给返回值命名
func test56(num1 int,num2 int) (sum int ,sub int) {
	sum:=num1+num2
	sub:=num1-num2
	return
}