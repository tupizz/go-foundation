package main

import "fmt"

// apenas métodos podem ser definidos dentro de uma interface
type AtivacaoMatricula interface {
	Desativar()
	Ativar()
}

type Aluno struct {
	Nome  string
	Email string
	Ativo bool
}

func (c *Aluno) Desativar() {
	c.Ativo = false
}

func (c *Aluno) Ativar() {
	c.Ativo = true
}

// aqui que a interface entra em ação, eu dependo de uma interface e não do Aluno em si
func ToggleMatriculaEscola(ativacaoDto AtivacaoMatricula, estaMatriculado bool) {
	if estaMatriculado {
		ativacaoDto.Desativar()
	} else {
		ativacaoDto.Ativar()
	}
}

func main() {
	tadeu := Aluno{
		Nome:  "Tadeu",
		Email: "tadeu.tupiz@gmail.com",
		Ativo: true,
	}

	tadeu.Desativar()
	fmt.Println(tadeu.Ativo)
	tadeu.Ativar()
	fmt.Println(tadeu.Ativo)

	estaMatriculado := true
	ToggleMatriculaEscola(&tadeu, estaMatriculado)
	fmt.Println(tadeu.Ativo)

}
