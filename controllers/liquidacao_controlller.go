package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type LiquidacaoController struct {
	repository *repositories.LiquidacaoRepository
}

func NewLiquidacaoController(repository *repositories.LiquidacaoRepository) *LiquidacaoController {
	return &LiquidacaoController{repository: repository}
}

func (controller *LiquidacaoController) GetLiquidacoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, fmt.Sprintf("Parâmetro 'unidadeGestoraCodigo' é obrigatório"), http.StatusBadRequest)
		return
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, fmt.Sprintf("Parâmetro 'ano' é obrigatório"), http.StatusBadRequest)
		return
	}
	codigoOrgao := r.URL.Query().Get("codigoOrgao")
	if codigoOrgao == "" {
		http.Error(w, fmt.Sprintf("Parâmetro 'código de Orgão' é obrigatório"), http.StatusBadRequest)
		return
	}
	codigoUnidadeOrcamentaria := r.URL.Query().Get("codigoUnidadeOrcamentaria")
	if codigoUnidadeOrcamentaria == "" {
		http.Error(w, fmt.Sprintf("Parâmetro 'código de Unidade Orçamentária' é obrigatório"), http.StatusBadRequest)
		return
	}
	numEmpenho := r.URL.Query().Get("numEmpenho")
	if numEmpenho == "" {
		http.Error(w, fmt.Sprintf("Parâmetro 'número de Empenho' é obrigatório "), http.StatusBadRequest)
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
	liquidacaoRepository := repositories.NewLiquidacaoRepository(db)
	liquidacoes, err := liquidacaoRepository.GetAllLiquidacao(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, numEmpenho)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Erro ao buscar liquidacoes: %v", err), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(liquidacoes)
}
