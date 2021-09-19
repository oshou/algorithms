package main

import "fmt"

func fizzbuzz(n int) {
	var s string
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			s = "FizzBuzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		default:
			s = fmt.Sprint(i)
		}
		fmt.Println(s)
	}
}

func main() {
	var n int
	fmt.Scan("n: ", &n)
	fizzbuzz(15)
}
