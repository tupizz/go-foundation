package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	assert.Equal(t, expected, result)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxDbRepositoryMock{}

	// criamos comportamentos ao mock
	repository.On("SaveTax", 10.0).Return(nil)

	err := CalculateTaxAndSave(1000.0, repository) // tax will be 10 in this case
	assert.Nil(t, err)
	repository.AssertExpectations(t) // devemos chamar para garantir que todas as expectativas foram atendidas
	repository.AssertCalled(t, "SaveTax", 10.0)
	repository.AssertNumberOfCalls(t, "SaveTax", 1)
}

func TestCalculateTaxAndSaveShouldThrow(t *testing.T) {
	repository := &TaxDbRepositoryMock{}

	// criamos comportamentos ao mock
	repository.On("SaveTax", 0.0).Return(errors.New("erros saving tax"))

	err := CalculateTaxAndSave(-5000.0, repository) // tax will be 0 in this case
	assert.Error(t, err, "error saving tax")
	repository.AssertExpectations(t)           // assert that all expectations were met, deve ser chamado
	repository.AssertCalled(t, "SaveTax", 0.0) // assert that a specific method was called, deve ser chamado
	repository.AssertNumberOfCalls(t, "SaveTax", 1)
}
