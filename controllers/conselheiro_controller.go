package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

type ConselheiroController struct {
	repository *repositories.ConselheiroRepository
}

func NewConselheiroController(repository *repositories.ConselheiroRepository) *ConselheiroController {
	return &ConselheiroController{repository: repository}
}

func (controller *ConselheiroController) GetConselheiros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	conselheirosRepository := repositories.NewConselheiroRepository(db)
	exercicio, err := strconv.Atoi(r.URL.Query().Get("exercicio"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	conselheiros, err := conselheirosRepository.GetAllConselheiros(exercicio)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonConselheiros, err := json.Marshal(conselheiros)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonConselheiros)

}
