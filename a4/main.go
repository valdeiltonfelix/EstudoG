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
	println(d)
	fmt.Printf("O tipo de E é %T\n", d)
}
