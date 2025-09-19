package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroNumConvenioController struct {
	repository *repositories.FiltroNumConvenioRepository
}

func NewFiltroNumConvenioController(repository *repositories.FiltroNumConvenioRepository) *FiltroNumConvenioController {
	return &FiltroNumConvenioController{repository: repository}
}

func (controller *FiltroNumConvenioController) GetFiltroNumConvenio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, fmt.Sprintf("parâmetro unidade gestora é obrigatório"), http.StatusInternalServerError)
		return
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, fmt.Sprintf("parâmetro ano é obrigatório"), http.StatusInternalServerError)
		return
	}
	numConvenio := r.URL.Query().Get("numConvenio")
	if numConvenio == "" {
		http.Error(w, fmt.Sprintf("parâmetro número de Convénio é obrigatório"), http.StatusInternalServerError)
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
	filtroConvenioRepository := repositories.NewFiltroNumConvenioRepository(db)
	filtroConvenioController, err := filtroConvenioRepository.GetFiltroNumConvenio(unidadeGestoraCodigo, ano, numConvenio)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtro de convénios", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroConvenioController)
}
