package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApiCep_GetCep(t *testing.T) {
	getCep := NewGetApiCep()
	channel := make(chan Result)
	go getCep.GetCep("15043-020", channel)
	cep := <-channel
	assert.Equal(t, "15043-020", cep.CepResult.Cep)
}
