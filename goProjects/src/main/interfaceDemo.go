package main

import "fmt"

type Book struct {
	Autor string
}

func myFun(arg interface{}) {
	fmt.Println(arg)
	value, ok := arg.(string)
	if !ok {
		fmt.Println("不是string类型")
	} else {
		fmt.Println(value)
	}
}
func main() {

	book := Book{"dou"}
	myFun(book)

	myFun(123)

	myFun("doulongafkkal")

}
