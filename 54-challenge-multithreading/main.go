package main

import (
	"54-challenge-multithreading/handler"
	"54-challenge-multithreading/request"
)

func main() {
	cep := "15091-330"

	apiCep := request.NewGetApiCep()
	viaCep := request.NewGetViaCep()

	handler.HandleGetCepData(cep, apiCep, viaCep)
}
