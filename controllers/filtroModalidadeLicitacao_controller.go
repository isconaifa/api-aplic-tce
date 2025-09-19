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

type FiltroModalidadeLicitacaoController struct {
	repository *repositories.FiltroModalidadeLicitacaoRepository
}

func NewFiltroModalidadeLicitacaoController(repository *repositories.FiltroModalidadeLicitacaoRepository) *FiltroModalidadeLicitacaoController {
	return &FiltroModalidadeLicitacaoController{repository: repository}
}

func (controller *FiltroModalidadeLicitacaoController) GetAllFiltroModalidadeLicitacao(w http.ResponseWriter, r *http.Request) {
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
	codigoModalidadeLicitacao, err := strconv.Atoi(r.URL.Query().Get("codigoModalidadeLicitacao"))
	if err != nil {
		http.Error(w, fmt.Sprintf("O parâmetro 'codigoModalidadeLicitacao' é obrigatório: %v", err), http.StatusBadRequest)
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
	filtroModalidadeLicitacaoRepository := repositories.NewFiltroModalidadeLicitacaoRepository(db)
	filtroModalidadeLicitacaos, err := filtroModalidadeLicitacaoRepository.GetAllFilterModalidadeLicitacao(unidadeGestoraCodigo, ano, codigoModalidadeLicitacao)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar filtroModalidadeLicitacaos: %v", err), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroModalidadeLicitacaos)
}
