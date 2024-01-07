package main

import "fmt"

func main() {
	fmt.Println(calc(1, 4))
}

func calc(a, b float64) (float64, float64, float64, float64) {
	soma := a + b
	subtra := a - b
	mult := a * b
	div := a / 2
	return soma, subtra, mult, div
}
