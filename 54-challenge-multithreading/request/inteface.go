package request

type CepResult struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type Result struct {
	Strategy  string    `json:"strategy"`
	CepResult CepResult `json:"cep_result"`
}

type GetCep interface {
	GetCep(cep string, channel chan Result) (*Result, error)
}
