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

	t := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := t.Execute(os.Stdout, Cursos{
		Curso{"GO", 40},
		Curso{"JAVA", 90},
		Curso{"Node.js", 120},
	})
	if err != nil {
		panic(err)
	}
}
