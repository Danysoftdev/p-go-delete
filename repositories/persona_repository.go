package repositories

import (
	"context"
	"time"

	"github.com/danysoftdev/p-go-delete/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Permite inyectar la colección desde fuera (ideal para pruebas)
func SetCollection(c *mongo.Collection) {
	collection = c
}

// ObtenerPersonaPorDocumento busca una persona por su Documento
func ObtenerPersonaPorDocumento(documento string) (models.Persona, error) {
	var persona models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"documento": documento}).Decode(&persona)
	return persona, err
}

// EliminarPersona elimina una persona por su Documento
func EliminarPersona(documento string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"documento": documento})
	return err
}

type RealPersonaRepository struct{}

func (r RealPersonaRepository) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	return ObtenerPersonaPorDocumento(doc)
}

func (r RealPersonaRepository) EliminarPersona(doc string) error {
	return EliminarPersona(doc)
}