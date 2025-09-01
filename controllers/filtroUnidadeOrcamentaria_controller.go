package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type FiltroUnidadeOrcamentariaController struct {
	repository *repositories.FiltroUnidadeOrcamentariaRepository
}

func NewFiltroUnidadeOrcamentariaController(repository *repositories.FiltroUnidadeOrcamentariaRepository) *FiltroUnidadeOrcamentariaController {
	return &FiltroUnidadeOrcamentariaController{repository: repository}
}

func (controller *FiltroUnidadeOrcamentariaController) GetAllFiltroUnidadeOrcamentaria(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroUnidadeOrcamentariaRepository := repositories.NewFiltroUnidadeOrcamentariaRepository(db)
	filtroUnidadeOrcamentaria, err := filtroUnidadeOrcamentariaRepository.GetAllFiltroUnidadeOrcamentaria(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria)
	if err != nil {
		http.Error(w, "Erro ao buscar filtroUnidadeOrcamentaria", http.StatusInternalServerError)
		return
	}
	jsonFiltroUnidadeOrcamentaria, err := json.Marshal(filtroUnidadeOrcamentaria)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroUnidadeOrcamentaria)
}
