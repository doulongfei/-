package main
import "fmt"
func main() {
	var count int =20
	if count<30{

		fmt.Println("fuibuqi")
	}
	if count1:=20;count1<30{
		fmt.Println("fuibuqi")

	}
	switch count {
		case 10,11,12:
			fmt.Println("switch的case后面可以有逗号分隔的多个条件。")
		case 9:
			fmt.Println("switch后面的条件可以使一个有返回结果的函数")
		case 8:
			fmt.Println("")
		case 7:
			fmt.Println("")
	}

	var a=1
	switch {
		case a==1:
			fmt.Println("aaa")
			fallthrough   //switch穿透（可执行满足条件的case和此case之后的一个case语句）
		case a==2:
			fmt.Println("bbb")
	}
	var sum int = 0
	for i := 0; i < count; i++ {
		sum+=i
	}
	fmt.Println(sum)
}