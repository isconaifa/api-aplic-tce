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

type FiltroFuncaoController struct {
	repository *repositories.FiltroFuncaoRepository
}

func NewFiltroFuncaoController(repository *repositories.FiltroFuncaoRepository) *FiltroFuncaoController {
	return &FiltroFuncaoController{repository: repository}
}

func (controller *FiltroFuncaoController) GetAllFiltroFuncao(w http.ResponseWriter, r *http.Request) {
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
	codigoFuncao, err := strconv.Atoi(r.URL.Query().Get("codigoFuncao"))
	if codigoFuncao == 0 {
		http.Error(w, "Parâmetro 'codigoFuncao' é obrigatório", http.StatusBadRequest)
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
	filtroFuncaoRepository := repositories.NewFiltroFuncaoRepository(db)
	filtroFuncoes, err := filtroFuncaoRepository.GetAllFiltroFuncao(unidadeGestoraCodigo, ano, codigoFuncao)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroFuncoes", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroFuncoes)
}
