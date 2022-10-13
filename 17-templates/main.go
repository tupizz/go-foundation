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
	curso := Curso{"GO", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso {{.Nome}} - CargaHorária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
