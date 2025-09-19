package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
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
		http.Error(w, fmt.Sprintf("ocorreu ao conectar ao banco %v:", err), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	competenciasRepository := repositories.NewCompetenciaRepository(db)
	competencias, err := competenciasRepository.GetAllCompetencias()
	if err != nil {
		http.Error(w, fmt.Sprintf("ocorreu ao buscar competencias %v:", err), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(competencias)
}
