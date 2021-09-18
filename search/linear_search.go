package linear_search

func linearSearch(arr []int, v int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}
