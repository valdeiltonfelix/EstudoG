package main

import "fmt"

func main() {

	// cont1 := iseq()
	// fmt.Println(cont1())
	// fmt.Println(cont1())
	// fmt.Println(cont1())
	// fmt.Println(cont1())
	// fmt.Println(cont1())
	// fmt.Println(cont1())

	// cont2 := iseq()
	// fmt.Println(cont2())
	// fmt.Println(cont2())

	// status1 := status()
	// fmt.Println(status1())
	// fmt.Println(status1())
	// fmt.Println(status1())

	rreturn := retorna()

	fmt.Println(rreturn(2, 5))
	fmt.Println(rreturn(2, 20))

}

// funcoes pode ter varios retornos
func soma(a, b float64) (float64, float64, float64, float64) {
	soma := a + b
	sub := a - b
	mult := a * b
	div := a / b

	return soma, sub, mult, div
}

// funções pode ter clouseros
func retorna() func(i, x int) int {

	return func(x, i int) int {
		return x * i
	}
}

func iseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func status() func() string {

	var index int = 0
	orderStatus := map[int]string{
		1: "PENDÊNCIA",
		2: "FAZENDO",
		3: "TERMINADO",
	}

	return func() string {
		fmt.Println(index)
		if index > 3 {
			index = 0
		}

		index++
		return orderStatus[index]
	}

}
