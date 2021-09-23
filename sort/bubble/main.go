package main

import (
	"fmt"
)

// 隣り合う要素で大小比較、大きな値を最後尾へ追いやる
// O(n^2): 二重for分のため
func bubble(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := bubble(unsorted)
	fmt.Println(sorted)
}
