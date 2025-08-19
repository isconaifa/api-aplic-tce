package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroNumLicitacoesController struct {
	repository *repositories.FiltroNumLicitacaoRepository
}

func NewFiltroNumLicitacoesController(repository *repositories.FiltroNumLicitacaoRepository) *FiltroNumLicitacoesController {
	return &FiltroNumLicitacoesController{repository: repository}
}

func (controller *FiltroNumLicitacoesController) GetAllFiltroNumLicitacoes(w http.ResponseWriter, r *http.Request) {
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
	numLicitacao := r.URL.Query().Get("numLicitacao")
	if numLicitacao == "" {
		http.Error(w, "Parâmetro 'numLicitacao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroNumLicitacaoRepository := repositories.NewFiltroNumLicitacaoRepository(db)
	filtroNumLicitacoes, err := filtroNumLicitacaoRepository.GetAllFiltroNumLicitacao(unidadeGestoraCodigo, ano, numLicitacao)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroNumLicitacoes", http.StatusInternalServerError)
		return
	}
	jsonFiltroNumLicitacoes, err := json.Marshal(filtroNumLicitacoes)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroNumLicitacoes)
}
