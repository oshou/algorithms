package main

import "fmt"

func selection(arr []int) []int {
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j <= len(arr)-1; j++ {
			// iから末尾までで最小要素を探す
			if j == i {
				minIndex = j
			}
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		// 現在位置の値と最小値をswitch
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := selection(unsorted)
	fmt.Println(sorted)
}
