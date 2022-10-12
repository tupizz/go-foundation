package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type BuscaCepResultado struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		fmt.Println(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))

		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao ler resposta: %v\n", err)
		}

		var data BuscaCepResultado
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao fazer parse da resposta: %v\n", err)
		}

		// If the file doesn't exist, create it, or append to the file
		//
		// 0644 =>
		// 		| owning user  | group | other (default) |
		//      |--------------|-------|-----------------|
		//		| read & write | read  | read            |
		//
		// rwx oct    meaning
		// 001 01   = execute
		// 010 02   = write
		// 011 03   = write & execute
		// 100 04   = read
		// 101 05   = read & execute
		// 110 06   = read & write
		// 111 07   = read & write & execute

		file, err := os.OpenFile("cidade.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		if _, err := file.Write([]byte(fmt.Sprintf("-----------------\nCEP: %s,\nLocalidade: %s,\nUF: %s\n", data.Cep, data.Localidade, data.Uf))); err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)
	}
}
