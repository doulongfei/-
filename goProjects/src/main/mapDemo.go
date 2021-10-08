package main

import (
	"fmt"
)

/*
map的定义 操作CRUD
*/
func main() {
	var a map[int]string
	//make函数的第二个参数可以不传  ，默认是1个
	a = make(map[int]string, 10)
	a[1] = "12"
	a[2] = "23"
	fmt.Println(a)
	//map使用前一定要使用make map的存储是无序的
	//	key是不可以重复的，如果有重复则会覆盖之前的value
	b := make(map[int]string)
	//增加
	b[4] = "124"
	b[5] = "12555554"
	b[6] = "53534"
	fmt.Println("b", b)
	//修改
	b[6] = "修改"
	fmt.Println("b", b)

	//删除  如果可以不存在则啥也不操作
	delete(b, 4)
	fmt.Println("b", b)
	//清空map集合①遍历集合②使用map:=make(map[T]T1) 重新make一个map

	//查找
	value, f := b[5]
	fmt.Println("对应的值：", value)
	fmt.Println("查找是否成功", f)
	c := map[int]string{
		8:  "1c24",
		45: "53c534",
	}
	fmt.Println(c)

	for k, v := range c {
		fmt.Printf("key:%v,  value:%v \n", k, v)
	}

	d := make(map[string]map[int]string)
	d["banji1"] = make(map[int]string)
	d["banji1"][1] = "lisi"
	d["banji1"][2] = "lisi2"
	d["banji1"][3] = "lisi3"

	d["班级2"] = make(map[int]string)
	d["班级2"][1] = "zhangsan"
	d["班级2"][2] = "zhangsan2"
	d["班级2"][3] = "zhangsan3"

	for k1, v1 := range d {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("学号：%v  ,姓名：%v \n", k2, v2)
		}
	}

}
