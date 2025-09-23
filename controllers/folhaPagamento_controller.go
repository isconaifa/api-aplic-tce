package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type FolhaPagamentoController struct {
	repository *repositories.FolhaPagamentoRepository
}

func NewFolhaPagamentoController(repository *repositories.FolhaPagamentoRepository) *FolhaPagamentoController {
	return &FolhaPagamentoController{repository: repository}
}

func (controller *FolhaPagamentoController) GetFolhaPagamento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, "Parâmetro 'unidadeGestoraCodigo' é obrigatório", http.StatusBadRequest)
		return
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, "Parâmetro 'ano' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco %v", err), http.StatusInternalServerError)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	folhaPagamentoRepository := repositories.NewFolhaPagamentoRepository(db)
	folhaPagamentoController, err := folhaPagamentoRepository.GetFolhaPagamento(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar folha de pagamento %v", err), http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(folhaPagamentoController)
}
