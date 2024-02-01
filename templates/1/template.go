package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{Nome: "Curso de Java", CargaHoraria: 50}
	tmpl := template.New("CursoTemplate")
	tmpl, _ = tmpl.Parse("Curso:{{.Nome}} Cargahoraria:{{.CargaHoraria}} \n")
	err := tmpl.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
