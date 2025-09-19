package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

type ExercicioController struct {
	repository *repositories.ExercicioRepository
}

func NewExercicioController(repository *repositories.ExercicioRepository) *ExercicioController {
	return &ExercicioController{repository: repository}
}

func (controller *ExercicioController) GetExercicios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	exerciciosRepository := repositories.NewExercicioRepository(db)
	exercicios, err := exerciciosRepository.GetAllExercicios()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	_ = json.NewEncoder(w).Encode(exercicios)
}
