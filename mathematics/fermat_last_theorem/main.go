package main

// Fermat's Last Theorem
//
// x**n + y**n  = z**n (n = 1,2)
// x**n + y**n != z**n (n >= 3)
// 0 < x < y < z
//
// ex)
// Input: 10, 1 => []
// Input: 10, 2 => [(3, 4, 5), (6, 8, 10)]
// Input: 10, 3 => []
// Input: 10, 4 => []

import (
	"fmt"
	"math"
	"time"
)

type Nums struct {
	x, y, z int
}

func fermatLastTheoremV1(maxNum, squareNum int) []Nums {
	result := []Nums{}
	sumMax := math.Pow(float64(maxNum-1), float64(squareNum)) + math.Pow(float64(maxNum), float64(squareNum))
	zMax := int(math.Pow(sumMax, 1.0/float64(squareNum)))
	for x := 1; x <= maxNum; x++ {
		for y := x + 1; y <= maxNum; y++ {
			for z := y + 1; z <= zMax; z++ {
				if math.Pow(float64(x), float64(squareNum))+math.Pow(float64(y), float64(squareNum)) == math.Pow(float64(z), float64(squareNum)) {
					result = append(result, Nums{x, y, z})
				}
			}
		}
	}
	return result
}

func fermatLastTheoremV2(maxNum, squareNum int) []Nums {
	result := []Nums{}
	for x := 1; x <= maxNum; x++ {
		for y := x + 1; y <= maxNum; y++ {
			sumPow := math.Pow(float64(x), float64(squareNum)) + math.Pow(float64(y), float64(squareNum))
			zFloat := math.Pow(sumPow, 1.0/float64(squareNum))
			if zFloat != float64(int(zFloat)) {
				continue
			}
			zInt := int(zFloat)
			zPow := math.Pow(float64(zInt), float64(squareNum))
			if sumPow == zPow {
				result = append(result, Nums{x, y, zInt})
			}
		}
	}
	return result
}

func main() {
	var (
		start, end time.Time
	)

	start = time.Now()
	fmt.Printf("v1_**1: %v\n", fermatLastTheoremV1(10, 1))
	fmt.Printf("v1_**2: %v\n", fermatLastTheoremV1(10, 2))
	fmt.Printf("v1_**3: %v\n", fermatLastTheoremV1(10, 3))
	fmt.Printf("v1_**4: %v\n", fermatLastTheoremV1(10, 4))
	fmt.Printf("v1_**5: %v\n", fermatLastTheoremV1(10, 5))
	fmt.Printf("v1_**5: %v\n", fermatLastTheoremV1(200, 5))
	end = time.Now()
	fmt.Printf("elapsed time:%f\n", end.Sub(start).Seconds())

	start = time.Now()
	fmt.Printf("v2_**1: %v\n", fermatLastTheoremV2(10, 1))
	fmt.Printf("v2_**2: %v\n", fermatLastTheoremV2(10, 2))
	fmt.Printf("v2_**3: %v\n", fermatLastTheoremV2(10, 3))
	fmt.Printf("v2_**4: %v\n", fermatLastTheoremV2(10, 4))
	fmt.Printf("v2_**5: %v\n", fermatLastTheoremV2(10, 5))
	fmt.Printf("v2_**5: %v\n", fermatLastTheoremV2(200, 5))
	end = time.Now()
	fmt.Printf("elapsed time:%f\n", end.Sub(start).Seconds())
}
