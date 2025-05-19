package repositories

import "github.com/danysoftdev/p-go-delete/models"

type PersonaRepository interface {
	
	ObtenerPersonaPorDocumento(documento string) (models.Persona, error)
	EliminarPersona(documento string) error
}
