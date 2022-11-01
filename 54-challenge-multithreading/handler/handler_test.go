package handler

import (
	"54-challenge-multithreading/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type CepMock struct {
	mock.Mock
	Strategy string
}

func (t *CepMock) GetCep(cep string, channel chan request.Result) (*request.Result, error) {
	args := t.Called(cep, channel)
	channel <- request.Result{
		Strategy: t.Strategy,
	}
	return args.Get(0).(*request.Result), args.Error(1)
}

func TestHandleGetCepData_WhenGetCep1IsFaster(t *testing.T) {
	cepMock1 := &CepMock{
		Strategy: "GetCep1",
	}
	cepMock2 := &CepMock{
		Strategy: "GetCep2",
	}

	cepMock1.On("GetCep", "12345678", mock.Anything).Return(&request.Result{
		Strategy: "GetCep1",
	}, nil).After(500 * time.Millisecond)

	cepMock2.On("GetCep", "12345678", mock.Anything).Return(&request.Result{
		Strategy: "GetCep2",
	}, nil).After(700 * time.Millisecond)

	result := HandleGetCepData("12345678", cepMock1, cepMock2)
	assert.Equal(t, "GetCep1", result.Strategy)
}

func TestHandleGetCepData_WhenGetCep2IsFaster(t *testing.T) {
	cepMock1 := &CepMock{
		Strategy: "GetCep1",
	}
	cepMock2 := &CepMock{
		Strategy: "GetCep2",
	}

	cepMock1.On("GetCep", "12345678", mock.Anything).Return(&request.Result{
		Strategy: "GetCep1",
	}, nil).After(1500 * time.Millisecond)

	cepMock2.On("GetCep", "12345678", mock.Anything).Return(&request.Result{
		Strategy: "GetCep2",
	}, nil).After(700 * time.Millisecond)

	result := HandleGetCepData("12345678", cepMock1, cepMock2)
	assert.Equal(t, "GetCep2", result.Strategy)
}

func TestHandleGetCepData_WhenTimeout(t *testing.T) {
	cepMock1 := &CepMock{
		Strategy: "GetCep1",
	}
	cepMock2 := &CepMock{
		Strategy: "GetCep2",
	}

	cepMock1.On("GetCep", "12345678", mock.Anything).Return(&request.Result{
		Strategy: "GetCep1",
	}, nil).After(1500 * time.Millisecond)

	cepMock2.On("GetCep", "12345678", mock.Anything).Return(&request.Result{
		Strategy: "GetCep2",
	}, nil).After(1500 * time.Millisecond)

	result := HandleGetCepData("12345678", cepMock1, cepMock2)
	assert.Equal(t, "timeout", result.Strategy)
}
