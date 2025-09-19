package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type AcaoController struct {
	repository *repositories.AcaoRepository
}

func NewAcaoController(repository *repositories.AcaoRepository) *AcaoController {
	return &AcaoController{repository: repository}
}

func (acaoController *AcaoController) GetAllAcoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, "Parâmetro 'unidadeGestoraCodigo' é obrigatório", http.StatusBadRequest)
		return
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, "Parâmetro 'exercicio' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	acoesRepository := repositories.NewAcaoRepository(db)
	acoes, err := acoesRepository.GetAllAcoes(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar ações: %v", err), http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(acoes)
}
