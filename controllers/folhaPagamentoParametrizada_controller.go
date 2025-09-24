package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FolhaPagamentoParametrizadaController struct {
	repository *repositories.FolhaPagamentoParametrizadaRepository
}

func NewFolhaPagamentoParametrizadaController(repository *repositories.FolhaPagamentoParametrizadaRepository) *FolhaPagamentoParametrizadaController {
	return &FolhaPagamentoParametrizadaController{repository: repository}
}

func (controller *FolhaPagamentoParametrizadaController) GetFolhaPagamentoParametrizada(w http.ResponseWriter, r *http.Request) {
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
	codigoOrgao := r.URL.Query().Get("codigoOrgao")
	codigoUnidadeOrcamentaria := r.URL.Query().Get("codigoUnidadeOrcamentaria")
	mesReferencia := r.URL.Query().Get("mesReferencia")

	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	folhasPagamentoRepository := repositories.NewFolhaPagamentoParametrizadaRepository(db)
	folhaPagamento, err := folhasPagamentoRepository.GetFolhaPagamentoParametrizada(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, mesReferencia)
	if err != nil {
		http.Error(w, "Erro ao processar sua solicitação", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(folhaPagamento); err != nil {
		log.Printf("Erro ao escrever a resposta JSON: %v", err)
	}
}
