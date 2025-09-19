package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FiltroDetalheFonteController struct {
	Repository *repositories.FiiltroDetalheFonteRepository
}

func NewFiltroDetalheFonteController(repository *repositories.FiiltroDetalheFonteRepository) *FiltroDetalheFonteController {
	return &FiltroDetalheFonteController{Repository: repository}
}

func (controller *FiltroDetalheFonteController) GetAllFiiltroDetalheFonte(w http.ResponseWriter, r *http.Request) {
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
	codigoDestinacaoRecurso, err := strconv.Atoi(r.URL.Query().Get("codigoDestinacaoRecurso"))
	if err != nil {
		http.Error(w, "ocorreu um erro", http.StatusBadRequest)
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
	filtroDetalheFonteRepository := repositories.NewFiiltroDetalheFonteRepository(db)
	filtroDetalheFontes, err := filtroDetalheFonteRepository.GetAllFiiltroDetalheFonte(unidadeGestoraCodigo, ano, codigoDestinacaoRecurso)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroDetalheFontes", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroDetalheFontes)
}
