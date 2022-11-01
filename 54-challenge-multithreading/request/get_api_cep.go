package request

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetApiCepResponse struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type GetApiCep struct {
}

func NewGetApiCep() *GetApiCep {
	return &GetApiCep{}
}

func (g *GetApiCep) GetCep(cep string, channel chan Result) (*Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", "https://cdn.apicep.com/file/apicep/"+cep+".json", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	var getViaCepResponse GetApiCepResponse
	err = json.NewDecoder(response.Body).Decode(&getViaCepResponse)
	if err != nil {
		return nil, err
	}

	cepResult := CepResult{
		Cep:          getViaCepResponse.Code,
		State:        getViaCepResponse.State,
		City:         getViaCepResponse.City,
		Neighborhood: getViaCepResponse.District,
		Street:       getViaCepResponse.Address,
	}

	result := Result{
		Strategy:  "Get Api Cep",
		CepResult: cepResult,
	}

	channel <- Result{
		Strategy:  "Get Api Cep",
		CepResult: cepResult,
	}

	return &result, nil
}
