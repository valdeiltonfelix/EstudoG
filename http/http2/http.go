package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8085", nil)

}

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

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		w.WriteHeader(http.StatusNotFound)
		//w.Write([]byte("Pagina não encontrada"))
		return
	}
	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := Cep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func Cep(cep string) (*BuscaCep, error) {
	resposta, erro := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if erro != nil {
		return nil, erro
	}
	defer resposta.Body.Close()
	body, erro := io.ReadAll(resposta.Body) //ioutil.ReadAll(resposta.Body)

	if erro != nil {
		return nil, erro
	}

	var mapear BuscaCep

	erro = json.Unmarshal(body, &mapear)

	if erro != nil {
		return nil, erro
	}

	return &mapear, nil
}
