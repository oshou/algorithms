package main

import "fmt"

var (
	left, right []int
)

func merge(arr []int) []int {
	var (
		left, right, ans []int
	)

	if len(arr) <= 1 {
		return arr
	}

	q := len(arr) / 2
	for i := 0; i < len(arr)-1; i++ {
		if i < q {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	// ソート
	left = merge(left)
	right = merge(right)

	// reverse
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ans = append(ans, left[0])
			left = left[1:]
		} else {
			ans = append(ans, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		ans = append(ans, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		ans = append(ans, right[0])
		right = right[1:]
	}

	return ans
}

func main() {
	unsorted := []int{1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52}
	sorted := merge(unsorted)
	fmt.Println(sorted)
}
