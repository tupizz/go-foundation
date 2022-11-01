package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetViaCep_GetCep(t *testing.T) {
	getCep := NewGetViaCep()
	channel := make(chan Result)
	go getCep.GetCep("15043-020", channel)
	cep := <-channel
	assert.Equal(t, "15043-020", cep.CepResult.Cep)
}
