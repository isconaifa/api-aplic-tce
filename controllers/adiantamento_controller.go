package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type AdiantamentoController struct {
	repository *repositories.AdiantamentoRepository
}

func NewAdiantamentoController(repository *repositories.AdiantamentoRepository) *AdiantamentoController {
	return &AdiantamentoController{repository: repository}
}

func (controller *AdiantamentoController) GetAdiantamentos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	adiantamentosRepository := repositories.NewAdiantamentoRepository(db)
	adiantamentos, err := adiantamentosRepository.GetAllAdiantamentos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonAdiantamentos, err := json.Marshal(adiantamentos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAdiantamentos)
}
