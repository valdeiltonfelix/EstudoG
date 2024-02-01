package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"n"`
	Saldo  float64 `json:"s"`
}

func main() {

	conta := Conta{Numero: 12323, Saldo: 3045.00}
	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		fmt.Println(err)
	}

	jsonPuro := []byte(`{"n":1000,"s":10}`)
	var conta3 Conta

	if json.Unmarshal(jsonPuro, &conta3) != nil {
		fmt.Println("Erro ao tenta encoder")
	}
	fmt.Println(conta3)
}
