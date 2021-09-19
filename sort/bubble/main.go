package main

import (
	"fmt"
)

func bubble(arr []int) []int {
	lastIndex := len(arr) - 1
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex-i; j++ {
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
