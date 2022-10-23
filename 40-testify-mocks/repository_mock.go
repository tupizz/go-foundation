package tax

import "github.com/stretchr/testify/mock"

type TaxDbRepositoryMock struct {
	mock.Mock
}

func (t *TaxDbRepositoryMock) SaveTax(tax float64) error {
	args := t.Called(tax)
	return args.Error(0)
}
