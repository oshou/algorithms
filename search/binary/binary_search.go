package binary_search

func binarySearch(arr []int, v int) bool {
	i := 0
	j := len(arr)
	for i != j {
		var m = (i + j) / 2
		switch {
		case v == arr[m]:
			return true
		case v < arr[m]:
			j = m
		case arr[m] < v:
			i = m + 1
		}
	}
	return false
}
