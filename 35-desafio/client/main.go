package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type QuotationResponse struct {
	Bid string `json:"bid"`
}

func getQuotationFromLocal() (error, *QuotationResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return err, nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err, nil
	}

	defer res.Body.Close()
	var response QuotationResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err, nil
	}

	return nil, &response
}

func appendQuotationInFile(quotation string) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if _, err = file.Write([]byte(fmt.Sprintf("Dólar: %v, Horário: %v\n", quotation, currentTime))); err != nil {
		return err
	}

	return nil
}

func main() {
	err, quotation := getQuotationFromLocal()
	if err != nil {
		panic(err)
	}

	err = appendQuotationInFile(quotation.Bid)
	if err != nil {
		panic(err)
	}
}
