package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type CompetenciaController struct {
	repository *repositories.CompetenciaRepository
}

func NewCompetenciaController(repository *repositories.CompetenciaRepository) *CompetenciaController {
	return &CompetenciaController{repository: repository}
}

func (controller *CompetenciaController) GetCompetencias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	competenciasRepository := repositories.NewCompetenciaRepository(db)
	competencias, err := competenciasRepository.GetAllCompetencias()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonCompetencias, err := json.Marshal(competencias)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCompetencias)
}
