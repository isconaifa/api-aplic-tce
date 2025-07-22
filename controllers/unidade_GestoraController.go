package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type UnidadeGestoraController struct {
	repository *repositories.UnidadeGestoraRepository
}

func NewUnidadeGestoraController(repository *repositories.UnidadeGestoraRepository) *UnidadeGestoraController {
	return &UnidadeGestoraController{repository: repository}
}

func (controller *UnidadeGestoraController) GetUnidadesGestoras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	unidadesGestorasRepository := repositories.NewUnidadeGestoraRepository(db)
	unidadesGestoras, err := unidadesGestorasRepository.GetAllUnidadesGestoras()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonUnidadesGestoras, err := json.Marshal(unidadesGestoras)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUnidadesGestoras)
}
