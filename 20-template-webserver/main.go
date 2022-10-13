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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		}

		t := template.Must(template.New("content.html").ParseFiles(templates...))

		model := Cursos{
			Curso{"GO", 40},
			Curso{"JAVA", 90},
			Curso{"Node.js", 120},
		}

		err := t.Execute(w, model)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
