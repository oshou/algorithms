package main

import "fmt"

type Arr []int

func linearSearch(arr Arr, v int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func main() {
	arr := Arr{1, 3, 9, 2, 7, 3, 4, 7}
	v := 3
	fmt.Println(linearSearch(arr, v))
}
