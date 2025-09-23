package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
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
	if codigoOrgao == "" {
		http.Error(w, "Parâmetro 'codigoOrgao' é obrigatório", http.StatusBadRequest)
		return
	}
	codigoUnidadeOrcamentaria := r.URL.Query().Get("codigoUnidadeOrcamentaria")
	if codigoUnidadeOrcamentaria == "" {
		http.Error(w, "Parâmetro 'codigoUnidadeOrcamentaria' é obrigatório", http.StatusBadRequest)
		return
	}
	mesReferencia := r.URL.Query().Get("mesReferencia")
	if mesReferencia == "" {
		http.Error(w, "Parâmetro 'mesReferencia' é obrigatório", http.StatusBadRequest)
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	folhaPagamentoParametrizadaRepository := repositories.NewFolhaPagamentoParametrizadaRepository(db)
	folhaPagamentoParametrizadas, err := folhaPagamentoParametrizadaRepository.GetFolhaPagamentoParametrizada(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, mesReferencia)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar folhaPagamentoParametrizadas: %v", err), http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(folhaPagamentoParametrizadas)
}
