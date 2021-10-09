package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user Is called...")
	fmt.Println("this  :%v\n", this)
}

func main() {
	users := User{1, "dou", 18}
	DoFiledAndMethod(users)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name())
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue:", inputValue)

	for i := 0; i < inputValue.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v ==%v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputValue.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v \n", m.Name, m.Type)

	}
}
