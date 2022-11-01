package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetViaCep_GetCep(t *testing.T) {
	getCep := NewGetViaCep()
	channel := make(chan Result)
	go getCep.GetCep("15091-330", channel)
	cep := <-channel
	assert.Equal(t, "15091-330", cep.CepResult.Cep)
}
