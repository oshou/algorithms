package main

import "fmt"

func main() {
	scores := []int{90, 75, 50, 80, 60}

	size := len(scores)
	for i := 0; i < size; i++ {
		fmt.Println(scores[i])
	}
}
