package handler

import (
	"54-challenge-multithreading/request"
	"encoding/json"
	"fmt"
	"time"
)

func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func HandleGetCepData(cep string, getCepApiOne, getCepApiTwo request.GetCep) request.Result {
	chApiOne := make(chan request.Result)
	chApiTwo := make(chan request.Result)

	go getCepApiOne.GetCep(cep, chApiOne)
	go getCepApiTwo.GetCep(cep, chApiTwo)

	var result request.Result

	select {
	case responseApiOne := <-chApiOne:
		result = responseApiOne
		prettyPrint(responseApiOne)
	case responseApiTwo := <-chApiTwo:
		result = responseApiTwo
		prettyPrint(responseApiTwo)
	case <-time.After(time.Second * 1):
		result = request.Result{Strategy: "timeout"}
		fmt.Println("timeout")
	}

	return result
}
