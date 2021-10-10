package main

import "fmt"

func sum11(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum11(a[:len(a)/2], c)
	go sum11(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
