package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroSomenteLiquidadosController struct {
	repository *repositories.FiltroSomenteLiquidadosRepository
}

func NewFiltroSomenteLiquidadosController(repository *repositories.FiltroSomenteLiquidadosRepository) *FiltroSomenteLiquidadosController {
	return &FiltroSomenteLiquidadosController{
		repository: repository,
	}
}

func (controller *FiltroSomenteLiquidadosController) GetAllFiltroSomenteLiquidados(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroSomenteLiquidadosRepository := repositories.NewFiltroSomenteLiquidadosRepository(db)
	filtroSomenteLiquidados, err := filtroSomenteLiquidadosRepository.GetAllFiltroSomenteLiquidados(unidadeGestoraCodigo, ano, dataInicioStr, dataFimStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar filtroSomenteAnulados: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtroSomenteLiquidados)
}
