package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// tamanho, err2 := f.WriteString("VALDEILTON DE SOUZA FELIX")
	tamanho, err2 := f.Write([]byte("VALDEILTON DE SOUZA FELIX"))

	if err2 != nil {
		panic(err2)
	}
	fmt.Println("O tamanho do arquivo é  ", tamanho)

	///
	f.Close()
	///leitura

	arq, err3 := os.ReadFile("arquivo.txt")

	if err3 != nil {
		panic(err3)
	}

	fmt.Println(string(arq))

	arq2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arq2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
