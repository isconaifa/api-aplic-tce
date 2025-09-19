package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
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
		http.Error(w, fmt.Sprintf("ocorreu um erro ao buscar conselheiro %v:", err), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	conselheirosRepository := repositories.NewConselheiroRepository(db)
	exercicio, err := strconv.Atoi(r.URL.Query().Get("exercicio"))
	if err != nil {
		http.Error(w, fmt.Sprintf("ocorreu um erro ao buscar conselheiro %v:", err), http.StatusBadRequest)
		return
	}
	conselheiros, err := conselheirosRepository.GetAllConselheiros(exercicio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar conselheiro %v:", err), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(conselheiros)
}
