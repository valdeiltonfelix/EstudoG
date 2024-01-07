package main

import "fmt"

func main() {
	fmt.Println(calc(1, 2, 5, 6, 7))
}

func calc(numeros ...float64) float64 {
	numtotal := 0.0
	for _, numero := range numeros {
		numtotal += numero
	}

	return numtotal
}
