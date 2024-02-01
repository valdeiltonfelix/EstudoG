package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
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
