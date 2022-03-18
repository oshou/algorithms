package main

import "fmt"

var (
	memo map[int]int
)

func init() {
	memo = make(map[int]int)
}

func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}

func fib2(n int) int {
	value, ok := memo[n]
	if ok {
		return value
	}

	if n == 1 || n == 2 {
		memo[n] = 1
		return 1
	} else {
		memo[n] = fib2(n-1) + fib2(n-2)
		return memo[n]
	}
}

func main() {
	fmt.Println(fib1(10))
	fmt.Println(fib1(100))
	fmt.Println(fib1(1000))

	fmt.Println(fib2(10))
	fmt.Println(fib2(100))
	fmt.Println(fib2(1000))
}
