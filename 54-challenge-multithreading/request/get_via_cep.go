package request

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetViaCepResponse struct {
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

type GetViaCep struct {
}

func NewGetViaCep() *GetViaCep {
	return &GetViaCep{}
}

func (g *GetViaCep) GetCep(cep string, channel chan Result) (*Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
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

	var getViaCepResponse GetViaCepResponse
	err = json.NewDecoder(response.Body).Decode(&getViaCepResponse)
	if err != nil {
		return nil, err
	}

	cepResult := CepResult{
		Cep:          getViaCepResponse.Cep,
		State:        getViaCepResponse.Uf,
		City:         getViaCepResponse.Localidade,
		Neighborhood: getViaCepResponse.Bairro,
		Street:       getViaCepResponse.Logradouro,
	}

	result := Result{
		Strategy:  "Get Via Cep",
		CepResult: cepResult,
	}

	channel <- Result{
		Strategy:  "Get Via Cep",
		CepResult: cepResult,
	}

	return &result, nil
}
