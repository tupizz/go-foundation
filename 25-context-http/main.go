package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	//ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://api.genderize.io?name=alex", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
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
