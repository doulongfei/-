package main
import "fmt"
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
}