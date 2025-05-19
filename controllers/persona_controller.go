package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/danysoftdev/p-go-delete/services"

	"github.com/gorilla/mux"
)

func EliminarPersona(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	documento := params["documento"]

	err := services.BorrarPersona(documento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Persona eliminada exitosamente"})
}
