package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FiltroAcaoController struct {
	repository *repositories.FiltroAcaoRepository
}

func NewFiltroAcaoController(repository *repositories.FiltroAcaoRepository) *FiltroAcaoController {
	return &FiltroAcaoController{repository: repository}
}
func (controller *FiltroAcaoController) GetAllFiltroAcao(w http.ResponseWriter, r *http.Request) {
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
	acaoCodigo, err := strconv.Atoi(r.URL.Query().Get("acaoCodigo"))
	if acaoCodigo == 0 {
		http.Error(w, "Parâmetro 'codigoFuncao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroAcaoRepository := repositories.NewFiltroAcaoRepository(db)
	filtroAcaos, err := filtroAcaoRepository.GetAllFiltroAcao(unidadeGestoraCodigo, ano, acaoCodigo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroAcaos", http.StatusInternalServerError)
		return
	}
	jsonFiltroAcaos, err := json.Marshal(filtroAcaos)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroAcaos)
}
