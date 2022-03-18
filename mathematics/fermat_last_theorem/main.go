package main

import (
	"fmt"
	"math"
)

type Nums struct {
	x, y, z int
}

func fermatLastTheoremV1(max, square int) []Nums {
	ans := []Nums{}
	for i := 1; i <= max; i++ {
		for j := 2; j <= max; j++ {
			sumPow := math.Pow(float64(i), float64(square)) + math.Pow(float64(j), float64(square))
			zFloat := math.Pow(sumPow, 1.0/float64(square))
			zInt := int(zFloat)

			if zFloat != float64(zInt) {
				continue
			}
			zPow := math.Pow(float64(zInt), float64(square))
			if sumPow == zPow {
				ans = append(ans, Nums{i, j, zInt})
			}
		}
	}
	return ans
}

func main() {
	var max, square int
	fmt.Scan(&max, &square)
	fmt.Println(fermatLastTheoremV1(max, square))
}
