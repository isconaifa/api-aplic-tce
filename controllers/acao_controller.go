package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
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
	exercicio := r.URL.Query().Get("exercicio")
	if exercicio == "" {
		http.Error(w, "Parâmetro 'exercicio' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	acoesRepository := repositories.NewAcaoRepository(db)
	acoes, err := acoesRepository.GetAllAcoes(unidadeGestoraCodigo, exercicio)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonAcoes, err := json.Marshal(acoes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAcoes)
}
