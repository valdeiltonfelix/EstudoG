package main

import "fmt"

const a = "bicicleta"

type id int

var (
	b bool
	c int
	e float64
	f string
	d id = 12
)

func main() {

	var meuArray [4]int
	meuArray[0] = 1
	meuArray[3] = 12

	for i, v := range meuArray {
		fmt.Printf("valor do indece %d é o valor %d \n", i, v)
	}

	fmt.Println(meuArray)
	fmt.Println("Tamanho do array", len(meuArray))
}
