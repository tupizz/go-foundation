package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// as tags estão entre crase e consitem em uma anotação de propriedades
type Conta struct {
	Numero int `json:"code"`
	Saldo  int `json:"amount"`
}

func ConverterEmJSON(valor any) string {
	// usamos Marshal para transformar em JSON uma struct de dados
	res, err := json.Marshal(valor)

	if err != nil {
		panic(err)
	}

	return string(res)
}

func EncodeStructNoTerminal(valor any) {
	// Usamos encoder para ao invés de armazenar em uma variável cuspir em algum lugar:
	// seja no terminal, seja em um webserver ou até mesmo um arquivo
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(valor)
}

func TransformaJsonEmStruct(bytes []byte, pointer *Conta) {
	err := json.Unmarshal(bytes, &pointer)
	if err != nil {
		panic(err)
	}
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  100,
	}

	ConverterEmJSON(conta)
	EncodeStructNoTerminal(conta)

	var contaFromJson Conta
	jsonPuro := []byte(`{ 
							"code": 2, 
							"amount": 200 
						}`)
	TransformaJsonEmStruct(jsonPuro, &contaFromJson)
	fmt.Println(contaFromJson.Saldo)
}
