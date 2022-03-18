package main

func binarySearch(arr []int, v int) bool {
	i := 0
	j := len(arr) - 1
	for i >= j {
		var m = (i + j) / 2
		switch {
		case v == arr[m]:
			return true
		case v < arr[m]:
			j = m
		case v > arr[m]:
			i = m + 1
		}
	}
	return false
}
