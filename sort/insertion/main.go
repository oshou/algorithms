package main

import "fmt"

func minIndex(arr []int) int {
	var (
		min    int
		minIdx int
	)
	for i, _ := range arr {
		if i == 0 {
			minIdx = 0
			min = arr[minIdx]
		} else {
			if arr[i] < min {
				minIdx = i
				min = arr[minIdx]
			}
		}
	}
	return minIdx
}

func insertion(arr []int) []int {
	for i, _ := range arr {
		idx := minIndex(arr[i : len(arr)-1])
		arr[i], arr[i+idx] = arr[i+idx], arr[i]
	}
	return arr
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := insertion(unsorted)
	fmt.Println(sorted)
}
