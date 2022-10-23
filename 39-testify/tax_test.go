package tax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result, err := CalculateTax(amount)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestCalculateTaxWhenTaxEqualsZero(t *testing.T) {
	amount := 0.0
	tax, err := CalculateTax(amount)

	assert.Error(t, err, "amount must be greater than zero")
	assert.Zero(t, tax)
}

func TestCalculateTaxWhenTaxGreaterThanThousands(t *testing.T) {
	amount := 1040.0
	tax, err := CalculateTax(amount)

	assert.Equal(t, tax, 10.0)
	assert.Nil(t, err)
}

func TestCalculateTaxWhenTaxGreaterThanTwentyThousands(t *testing.T) {
	amount := 31040.0
	tax, err := CalculateTax(amount)

	assert.Equal(t, tax, 20.0)
	assert.Nil(t, err)
}
