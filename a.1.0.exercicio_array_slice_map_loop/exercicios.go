package main

import "fmt"

func main() {

	// i := 0
	//entedendo loops
	// for i < 100 {
	// 	println(i)
	// 	if i == 3 {
	// 		break
	// 	}
	// 	i++
	// }

	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		println(i)
	// 	}

	// }

	// j := 0
	// for {
	// 	println(" Enquanto tiver memoria continua a fazer ", j)
	// 	j++
	// 	if j == 10 {
	// 		break
	// 	}

	// }

	//ecreva uma matriz ou array na liguagem go
	//aqui é elementos fixos de inteiros
	var inteiros [5]int
	inteiros[0] = 6
	inteiros[3] = 10

	// informe o tamanho do array ou matriz
	//println("O tamanho é ", len(inteiros))
	//informe a capacidade
	//println("A capacidade é ", cap(inteiros))

	// criae uma matriz de matriz
	var matriz [2][3]int
	matriz[1][0] = 4
	// fmt.Println(matriz)

	//percorra todos os elementos da matriz
	// for i := range matriz {
	// 	fmt.Println(matriz[i])
	// 	for index, valor := range matriz[i] {
	// 		fmt.Println(index, valor)
	// 	}
	// }

	//slice.... consegue manipular os elementos em tempo de execução. seu tamanho pode aumentar o diminuir
	// Crie um elemento de inteiros com slice
	it := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//index := 3
	//it = append(it, 11, 12, 13)

	for i := 0; i < len(it); i++ {
		it = append(it[:i], it[len(it)-1:]...)

	}
	/// aqui é usada para remover um indece do slice
	// fmt.Println(it[:index], it[index:])
	// it = append(it[:index], it[index+1:]...)
	//fmt.Println(it)

	///definindo um mapa
	///Essa tecnica consiste em cria tipos associativos
	// essa forma é a literal
	mapa1 := map[string]int{}
	mapa1["valor1"] = 35
	mapa1["valor2"] = 45

	fmt.Println(mapa1["valor1"])
	//Nao se pode inicializar um map com o valor zerado
	var mapa2 map[string]string
	//mapa2["nome1"] = "valdeilton"
	// mapa2["nome"] = "Valdeilton"
	fmt.Println(mapa2)

	//utilizando a funcao make
	map3 := make(map[string]string)
	map3["nome"] = "Valdeilton"
	map3["idade"] = "53"
	fmt.Println(map3)
	//se a chave nao existir vai retorna false,, sintaxe abaixo permite que pegamos o valor e a existencia
	valor, seexiste := map3["cpf"]

	fmt.Println(valor, seexiste)

}
