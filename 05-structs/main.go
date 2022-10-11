package main

import "fmt"

type Endereco struct {
	Rua    string
	Estado string
	Cidade string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// we bind a function to a given struct
func (c *Client) Desativar() {
	c.Ativo = false
}

func main() {
	tadeu := Client{
		Nome:  "Tadeu",
		Idade: 28,
		Ativo: true,
		Endereco: Endereco{
			Rua:    "Rua maria loureiro",
			Estado: "MG",
			Cidade: "Araxá",
		},
	}

	fmt.Println(tadeu.Rua)
	fmt.Println(tadeu.Endereco.Rua)

	fmt.Println(tadeu.Ativo)
	tadeu.Desativar()
	fmt.Println(tadeu.Ativo)
}
