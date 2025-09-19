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

type FiltroSubFuncaoController struct {
	repository *repositories.FiltroSubFuncaoRepository
}

func NewFiltroSubFuncaoController(repository *repositories.FiltroSubFuncaoRepository) *FiltroSubFuncaoController {
	return &FiltroSubFuncaoController{repository: repository}
}

func (controller *FiltroSubFuncaoController) GetAllFiltroSubFuncao(w http.ResponseWriter, r *http.Request) {
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
	codigoSubFuncao, err := strconv.Atoi(r.URL.Query().Get("codigoSubFuncao"))
	if codigoSubFuncao == 0 {
		http.Error(w, "Parâmetro 'codigoSubFuncao' é obrigatório", http.StatusBadRequest)
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
	filtroSubFuncaoRepository := repositories.NewFiltroSubFuncaoRepository(db)
	filtroSubFuncoes, err := filtroSubFuncaoRepository.GetAllFiltroSubFuncao(unidadeGestoraCodigo, ano, codigoSubFuncao)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroSubFuncoes", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroSubFuncoes)
}
