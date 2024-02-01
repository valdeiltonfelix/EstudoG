package main

import "fmt"

// compondo um struct
type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

// criando uma struct
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Pessoa interface {
	Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	client1 := Cliente{
		"Valdeilton",
		34,
		true,
		Endereco{"Rua das carmelias", "s/n", "Ibotirama", "Ba"},
	}
	Desativacao(client1)
	// client1.Desativar()
	fmt.Println(client1)
}
