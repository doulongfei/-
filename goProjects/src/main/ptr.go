package main
import(
	"fmt"
)
func main(){
	var age int =18
	fmt.Println(&age)

	var ptr *int =&age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：",&ptr)
	fmt.Println("ptr指向的数值为：",*ptr)
	//通过指针改变指向值
	var num int =10
	fmt.Println(num)
	//指针变量接受的一定是地址值
	var ptr1 *int =&num
	*ptr1=20
	fmt.Println(num)
	//指针变量的地址不可以不匹配 var num int =10和var ptr1 *int =&num必须都为int 要匹配

	// 基本数据类型（值类型），都有对应的指针类型，形式为*数据类型
	//int =》 *int  。。。

	


}