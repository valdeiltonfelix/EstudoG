package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCep)
	http.ListenAndServe(":8085", nil)

}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, vou acessou meu servidor"))
}
