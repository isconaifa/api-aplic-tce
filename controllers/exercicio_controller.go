package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
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
	exerciciosRepository := repositories.NewExercicioRepository(db)
	exercicios, err := exerciciosRepository.GetAllExercicios()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonExercicios, err := json.Marshal(exercicios)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonExercicios)
}
