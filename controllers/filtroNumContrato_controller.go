package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroNumContratoController struct {
	repository *repositories.FiltroNumContratoRepository
}

func NewFiltroNumContratoController(repository *repositories.FiltroNumContratoRepository) *FiltroNumContratoController {
	return &FiltroNumContratoController{repository: repository}
}

func (controller *FiltroNumContratoController) GetFiltroNumContrato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, fmt.Sprintf("Parametro unidadeGestoraCodigo obrigatório"), http.StatusInternalServerError)
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, fmt.Sprintf("Parametro ano obrigatório"), http.StatusInternalServerError)
	}
	numContrato := r.URL.Query().Get("numContrato")
	if numContrato == "" {
		http.Error(w, fmt.Sprintf("Parametro numContrato obrigatório"), http.StatusInternalServerError)
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	numContratoRepo := repositories.NewFiltroNumContratoRepository(db)
	numContratoController, err := numContratoRepo.GetFiltroNumContrato(unidadeGestoraCodigo, ano, numContrato)
	if err != nil {
		http.Error(w, "Erro ao buscar ao banco", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(numContratoController)

}
