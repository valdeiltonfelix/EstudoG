package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BuscaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		result, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
		errorHandler(err, "Erro ao tentantar busca da url")
		defer result.Body.Close()

		result2, err := io.ReadAll(result.Body)
		errorHandler(err, "Erro ao tentantar ler os dados")

		var data BuscaCep

		err = json.Unmarshal(result2, &data)
		errorHandler(err, "Erro ao tentantar converte json")

		file, err := os.Create("cidade.txt")
		errorHandler(err, "Erro ao tentantar grava no arquivo")

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, LOCALIDADE %s UF %s \n", data.Cep, data.Localidade, data.Uf))
		errorHandler(err, "Erro ao tentantar grava no arquivo")
		fmt.Println("Arquivo criado com sucesso !! cidade.txt ")

	}
}

func errorHandler(err error, menssage string) {

	if err != nil {
		panic(menssage)
	}

}
