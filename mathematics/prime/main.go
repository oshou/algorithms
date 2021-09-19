package main

import "fmt"

func getPrime(n int) []int {
	var ans []int
	arr := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = true
	}
	for i := 2; i*i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			if j != 1 {
				arr[i*j] = false
			}
		}
	}
	for i := 1; i <= n; i++ {
		if arr[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(getPrime(10000))
}
