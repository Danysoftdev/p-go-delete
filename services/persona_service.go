package services

import (
	"errors"
	"strings"

	"github.com/danysoftdev/p-go-delete/models"
	"github.com/danysoftdev/p-go-delete/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

var Repo repositories.PersonaRepository

func SetPersonaRepository(r repositories.PersonaRepository) {
	Repo = r
}

func BuscarPersonaPorDocumento(doc string) (models.Persona, error) {
	if strings.TrimSpace(doc) == "" {
		return models.Persona{}, errors.New("el documento no puede estar vacío")
	}

	persona, err := Repo.ObtenerPersonaPorDocumento(doc)
	if err == mongo.ErrNoDocuments {
		return models.Persona{}, errors.New("persona no encontrada")
	}

	return persona, err
}

func BorrarPersona(documento string) error {
	if strings.TrimSpace(documento) == "" {
		return errors.New("el documento no puede estar vacío")
	}

	_, err := Repo.ObtenerPersonaPorDocumento(documento)
	if err == mongo.ErrNoDocuments {
		return errors.New("persona no encontrada")
	}

	return Repo.EliminarPersona(documento)
}
