package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", Home)
	http.ListenAndServe(":8282", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(w, Cursos{
		{"Java", 50},
		{"go", 70},
		{"Python", 120},
		{"Javascript", 200},
		{"Elexir", 150},
	})
	if err != nil {
		panic(err)
	}
}
