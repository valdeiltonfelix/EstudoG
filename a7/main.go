package main

import (
	"fmt"
)

func main() {
	salario := map[string]float64{}
	salario["Valdeilton"] = 2500
	salario["Jose"] = 2600
	salario["Maria"] = 2700
	for nome, salario := range salario {
		fmt.Printf("%v salario de %f \n", nome, salario)
	}

	inteiros := make(map[string]int)
	inteiros["numero1"] = 2
	inteiros["numero2"] = 3
	fmt.Println(inteiros["numero2"])

}
