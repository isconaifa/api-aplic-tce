package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroNumEmpenhoController struct {
	repository *repositories.FiltroNumEmpenhoRepository
}

func NewFiltroNumEmpenhoController(repository *repositories.FiltroNumEmpenhoRepository) *FiltroNumEmpenhoController {
	return &FiltroNumEmpenhoController{repository: repository}
}

func (controller *FiltroNumEmpenhoController) GetFiltroNumEmpenho(w http.ResponseWriter, r *http.Request) {
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
	numEmpenho := r.URL.Query().Get("numEmpenho")
	if numEmpenho == "" {
		http.Error(w, "Parâmetro 'numEmpenho' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroNumEmpenhoRepository := repositories.NewFiltroNumEmpenhoRepository(db)
	filtroNumEmpenho, err := filtroNumEmpenhoRepository.GetFiltroNumEmpenho(unidadeGestoraCodigo, ano, numEmpenho)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroNumEmpenho", http.StatusInternalServerError)
		return
	}
	jsonFiltroNumEmpenho, err := json.Marshal(filtroNumEmpenho)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroNumEmpenho)
}
