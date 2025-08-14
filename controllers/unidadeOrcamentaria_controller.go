package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type UnidadeOrcamentariaController struct {
	repository *repositories.UnidadeOrcamentariaRepository
}

func NewUnidadeOrcamentariaController(repository *repositories.UnidadeOrcamentariaRepository) *UnidadeOrcamentariaController {
	return &UnidadeOrcamentariaController{repository: repository}
}

func (controller *UnidadeOrcamentariaController) GetAllUnidadeOrcamentaria(w http.ResponseWriter, r *http.Request) {
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
	codigoOrgao := r.URL.Query().Get("codigoOrgao")
	if codigoOrgao == "" {
		http.Error(w, "Parâmetro 'codigoOrgao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	unidadeOrcamentariaRepository := repositories.NewUnidadeOrcamentariaRepository(db)
	unidadeOrcamentaria, err := unidadeOrcamentariaRepository.GetAllUnidadeOrcamentaria(unidadeGestoraCodigo, ano, codigoOrgao)
	if err != nil {
		http.Error(w, "Erro ao buscar Unidade Orcamentaria", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonUnidadeOrcamentaria, err := json.Marshal(unidadeOrcamentaria)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUnidadeOrcamentaria)
}
