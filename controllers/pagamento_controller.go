package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type PagamentoController struct {
	repository *repositories.PagamentoRepository
}

func NewPagamentoController(repository *repositories.PagamentoRepository) *PagamentoController {
	return &PagamentoController{repository: repository}
}

func (controller *PagamentoController) GetPagamentos(w http.ResponseWriter, r *http.Request) {
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
	numEmpenho := r.URL.Query().Get("numEmpenho")
	if numEmpenho == "" {
		http.Error(w, "Parâmetro 'numEmpenho' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	pagamentoRepository := repositories.NewPagamentoRepository(db)
	pagamentos, err := pagamentoRepository.GetPagamentos(unidadeGestoraCodigo, ano, numEmpenho)
	println(pagamentos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar pagamentos: %v", err), http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(pagamentos)
}
