package main

import "fmt"

func main() {
	///closure em go
	total := func() float64 {
		return calc(1, 2, 5, 6, 7)
	}

	fmt.Println(total)
}

func calc(numeros ...float64) float64 {
	numtotal := 0.0
	for _, numero := range numeros {
		numtotal += numero
	}

	return numtotal
}
