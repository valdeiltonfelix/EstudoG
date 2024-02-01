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

	t := template.Must(template.New("CursosTemplate").Parse("Cursos:{{.Nome}}  Carga horaria: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
