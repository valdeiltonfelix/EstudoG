package main

import "fmt"

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {

	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {

	var soma float64

	for _, v := range m {
		soma += v
	}

	return soma
}

func main() {
	////generics

	mp := map[string]int{"joao": 1000, "maria": 2000, "valdeilton": 1500}
	mp2 := map[string]float64{"joao": 1000.10, "maria": 200.20, "valdeilton": 100.20}

	fmt.Println(Soma(mp))
	fmt.Println(Soma(mp2))
	// fmt.Println(SomaFloat(mp2))
}
