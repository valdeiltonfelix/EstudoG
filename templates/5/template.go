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
	template := []string{
		"static/header.html",
		"dinamic/contente.html",
		"static/footer.html",
	}
	http.HandleFunc("/", Home(template))
	http.ListenAndServe(":8182", nil)
}

func Home(contente []string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		t := template.Must(template.New("contente.html").ParseFiles(contente...))
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
}
