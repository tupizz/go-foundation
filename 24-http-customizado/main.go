package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	request, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "application/json")
	response, err := c.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
