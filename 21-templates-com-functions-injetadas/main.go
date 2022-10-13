package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func FormataHoras(horas int) string {
	return fmt.Sprintf("%d horas", horas)
}

func main() {
	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"ToUpper":      ToUpper,
		"FormataHoras": FormataHoras,
	})
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t = template.Must(t.ParseFiles(templates...))

	model := Cursos{
		Curso{"go", 40},
		Curso{"java", 90},
		Curso{"node.js", 120},
	}

	err := t.Execute(os.Stdout, model)

	if err != nil {
		panic(err)
	}
}
