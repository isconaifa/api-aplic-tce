package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

type FuncaoController struct {
	repository *repositories.FuncaoRepository
}

func NewFuncaoController(repository *repositories.FuncaoRepository) *FuncaoController {
	return &FuncaoController{repository: repository}
}

func (controller *FuncaoController) GetFuncoes(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	funcoesRepository := repositories.NewFuncaoRepository(db)
	funcoes, err := funcoesRepository.GetAllFuncoes(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, "Erro ao buscar funcoes", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(funcoes)
}
