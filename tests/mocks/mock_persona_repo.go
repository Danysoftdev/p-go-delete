package mocks

import (
	"github.com/danysoftdev/p-go-delete/models"
	"github.com/stretchr/testify/mock"
)

// MockPersonaRepo implementa la interfaz PersonaRepository para pruebas
type MockPersonaRepo struct {
	mock.Mock
}

func (m *MockPersonaRepo) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	args := m.Called(doc)
	return args.Get(0).(models.Persona), args.Error(1)
}

func (m *MockPersonaRepo) EliminarPersona(doc string) error {
	args := m.Called(doc)
	return args.Error(0)
}
