package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Valdeilton felix"
	fmt.Printf("O cliente %v  andou \n", c.nome)
}
func main() {
	cliente := Cliente{"Valdeilton"}
	//fmt.Println(cliente)
	cliente.andou()
	fmt.Println(cliente)
}
