package main

import "fmt"

func selection(arr []int) []int {
	minIndex := -1
	size := len(arr) - 1
	for i := 0; i < size; i++ {
		for j := i; j <= size; j++ {
			if minIndex < 0 {
				minIndex = j
			}
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		minIndex = -1
	}
	return arr
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := selection(unsorted)
	fmt.Println(sorted)
}
