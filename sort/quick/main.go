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

	// pivotを設定
	pivot := arr[0]

	// pivotと大小比較して振り分け
	for _, v := range arr[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	// 2分した各要素に対して再帰実行
	left = quick(left)
	right = quick(right)

	// ソート(left -> pivot -> right)
	ans = append(left, pivot)
	ans = append(ans, right...)

	return ans
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := quick(unsorted)
	fmt.Println(sorted)
}
