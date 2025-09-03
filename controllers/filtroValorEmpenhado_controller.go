package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type FiltroValorEmpenhadoController struct {
	repository *repositories.FiltroValorEmpenhadoRepository
}

func NewFiltroValorEmpenhadoController(repository *repositories.FiltroValorEmpenhadoRepository) *FiltroValorEmpenhadoController {
	return &FiltroValorEmpenhadoController{repository: repository}
}

func (controller *FiltroValorEmpenhadoController) GetAllFiltroValorEmpenhado(w http.ResponseWriter, r *http.Request) {
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
	valorInicialStr := r.URL.Query().Get("valorInicialStr")
	if valorInicialStr == "" {
		http.Error(w, "Parâmetro 'valorInicialStr' é obrigatório", http.StatusBadRequest)
		return
	}
	valorFinalStr := r.URL.Query().Get("valorFinalStr")
	if valorFinalStr == "" {
		http.Error(w, "Parâmetro 'valorFinalStr' é obrigatório", http.StatusBadRequest)
		return
	}
	if valorFinalStr < valorInicialStr {
		http.Error(w, "Valor final deve ser maior que valor inicial", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroValorEmpenhadoRepository := repositories.NewFiltroValorEmpenhadoRepository(db)
	filtroValorEmpenhados, err := filtroValorEmpenhadoRepository.GetAllFiltroValorEmpenhado(unidadeGestoraCodigo, ano, valorInicialStr, valorFinalStr)
	if err != nil {
		http.Error(w, "Erro ao buscar programas", http.StatusInternalServerError)
		return
	}
	conversaoJSON, err := json.Marshal(filtroValorEmpenhados)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(conversaoJSON)
}
