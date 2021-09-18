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
	arr := []int{80, 65, 70, 90, 30, 25, 35}
	sorted := bubble(arr)
	fmt.Println(sorted)
}
