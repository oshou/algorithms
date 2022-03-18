package main

import "fmt"

func linearSearch(arr []int, v int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 5, 2, 3, 4}
	v := 9
	fmt.Println(linearSearch(arr, v))
}
