package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroSomenteAnuladosController struct {
	repository *repositories.FiltroSomenteAnuladosRepository
}

func NewFiltroSomenteAnuladosController(repository *repositories.FiltroSomenteAnuladosRepository) *FiltroSomenteAnuladosController {
	return &FiltroSomenteAnuladosController{repository: repository}
}

func (controller *FiltroSomenteAnuladosController) GetAllFiltroSomenteAnulados(w http.ResponseWriter, r *http.Request) {
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
	dataInicioStr := r.URL.Query().Get("dataInicioStr")
	if dataInicioStr == "" {
		http.Error(w, "Parâmetro 'dataInicioStr' é obrigatório", http.StatusBadRequest)
		return
	}
	dataFimStr := r.URL.Query().Get("dataFimStr")
	if dataFimStr == "" {
		http.Error(w, "Parâmetro 'dataFimStr' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroSomenteAnuladosRepository := repositories.NewFiltroSomenteAnuladosRepository(db)
	filtroSomenteAnulados, err := filtroSomenteAnuladosRepository.GetAllFiltroSomenteAnulados(unidadeGestoraCodigo, ano, dataInicioStr, dataFimStr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Erro ao buscar filtroSomenteAnulados: %v", err), http.StatusInternalServerError)
		return
	}
	jsonFiltroSomenteAnulados, err := json.Marshal(filtroSomenteAnulados)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroSomenteAnulados)
}
