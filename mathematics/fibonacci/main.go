package main

import "fmt"

// メモ化
var (
	memo map[int]int
)

func init() {
	memo = make(map[int]int)
}

func fibonacci(n int) int {
	value, ok := memo[n]
	if ok {
		return value
	}

	if n == 1 || n == 2 {
		memo[n] = 1
		return 1
	} else {
		memo[n] = fibonacci(n-1) + fibonacci(n-2)
		return memo[n]
	}
}

func main() {
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
}
