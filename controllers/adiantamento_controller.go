package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type AdiantamentoController struct {
	repository *repositories.AdiantamentoRepository
}

func NewAdiantamentoController(repository *repositories.AdiantamentoRepository) *AdiantamentoController {
	return &AdiantamentoController{repository: repository}
}

func (controller *AdiantamentoController) GetAdiantamento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, "Parâmetro 'unidadeGestoraCodigo' é obrigatório", http.StatusBadRequest)
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, "Parâmetro 'ano' é obrigatório", http.StatusBadRequest)
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	adiantamentoRepository := repositories.NewAdiantamentoRepository(db)
	adiantamentos, err := adiantamentoRepository.GetAdiantamento(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar adiantamentos: %v", err), http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(adiantamentos)
}
