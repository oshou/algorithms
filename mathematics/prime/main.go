package main

import "fmt"

func prime(n int) []int {
	arr := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = true
	}

	for i := 2; i*i <= n; i++ {
		for j := 2; i*j <= n; j++ {
			arr[i*j] = false
		}
	}

	var ans []int
	for i := 1; i < n; i++ {
		if arr[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(prime(n))
}
