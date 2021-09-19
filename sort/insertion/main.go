package main

import "fmt"

func insertion(arr []int) []int {
	sorted = []int
	for i := 1; i <= len(arr)-1; i++ {
		insertIndex := i
		for j := insertIndex-1; j >=0; j-- {
			if arr[i] < arr[j]{
				arr[j+1]=arr[j]
				insertIndex-=1
			}
		}
	}
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := selection(unsorted)
	fmt.Println(sorted)
}
