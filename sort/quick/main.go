package main

import "fmt"

func quick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var (
		ans   []int
		left  []int
		right []int
	)

	// Pivot基準に大小で2分割
	pivot := arr[0]
	for _, v := range arr[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	// 分割した要素でそれぞれ再帰実行
	left = quick(left)
	right = quick(right)

	// 順序制御
	ans = append(left, pivot)
	ans = append(ans, right...)

	return ans
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := quick(unsorted)
	fmt.Println(sorted)
}
