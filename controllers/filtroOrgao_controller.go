package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroOrgaoController struct {
	repository *repositories.FiltroOrgaoRepository
}

func NewFiltroOrgaoController(repository *repositories.FiltroOrgaoRepository) *FiltroOrgaoController {
	return &FiltroOrgaoController{repository: repository}
}
func (controller *FiltroOrgaoController) GetAllFiltroOrgao(w http.ResponseWriter, r *http.Request) {
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
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	filtroOrgaoRepository := repositories.NewFiltroOrgaoRepository(db)
	filtroOrgaos, err := filtroOrgaoRepository.GetAllFiltroOrgao(unidadeGestoraCodigo, ano, codigoOrgao)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroOrgaos", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroOrgaos)
}
