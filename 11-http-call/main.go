package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.tadeutupinamba.com.br")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(res))
}
