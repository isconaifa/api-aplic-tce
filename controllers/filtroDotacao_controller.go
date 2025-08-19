package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type FiltroDotacaoController struct {
	repository *repositories.FiltroDotacaoRepository
}

func NewFiltroDotacaoController(repository *repositories.FiltroDotacaoRepository) *FiltroDotacaoController {
	return &FiltroDotacaoController{repository: repository}
}

func (controller *FiltroDotacaoController) ObterFiltroDotacao(w http.ResponseWriter, r *http.Request) {
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
	dotacao := r.URL.Query().Get("dotacao")
	if dotacao == "" {
		http.Error(w, "Parâmetro 'dotacao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroDotacaoRepository := repositories.NewFiltroDotacaoRepository(db)
	filtroDotacao, err := filtroDotacaoRepository.ObterFiltroDotacao(unidadeGestoraCodigo, ano, dotacao)
	if err != nil {
		http.Error(w, "Erro ao buscar filtroDotacao", http.StatusInternalServerError)
		return
	}
	filtrodotacaoJson, err := json.Marshal(filtroDotacao)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(filtrodotacaoJson)
}
