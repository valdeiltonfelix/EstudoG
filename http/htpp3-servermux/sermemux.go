package main

import (
	"net/http"
)

type blog struct {
	title string
}

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Word!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Meu blog"})
	http.ListenAndServe(":8081", mux)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "arquivo.html")
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
