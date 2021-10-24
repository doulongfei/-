package main

import "fmt"

func fib(n int) int {
	memo := make([]int, n+1)
	return dfs(n, memo)
	return 0
}

func dfs(n int, memo []int) int {
	if n <= 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = (dfs(n-1, memo) + dfs(n-2, memo)) % 1000000007
	return memo[n]
}
func main() {
	num := fib(5)
	fmt.Println(num)
}
