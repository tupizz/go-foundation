package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{ "name": "tadeu" }`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
