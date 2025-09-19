package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroNumConcursoController struct {
	repository *repositories.FiltroNumConcursoRepository
}

func NewFiltroNumConcursoController(repository *repositories.FiltroNumConcursoRepository) *FiltroNumConcursoController {
	return &FiltroNumConcursoController{repository: repository}
}

func (controller *FiltroNumConcursoController) GetFiltroNumConcurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, fmt.Sprintf("Parâmetro unidadeGestoraCodigo obrigatório"), http.StatusInternalServerError)
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, fmt.Sprintf("Parâmetro ano obrigatório"), http.StatusInternalServerError)
	}
	numConcurso := r.URL.Query().Get("numConcurso")
	if numConcurso == "" {
		http.Error(w, fmt.Sprintf("Parâmetro número de Concurso é obrigatório"), http.StatusInternalServerError)
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	numConcursoRepo := repositories.NewFiltroNumConcursoRepository(db)
	numConcursoController, err := numConcursoRepo.GetFiltroNumConcurso(unidadeGestoraCodigo, ano, numConcurso)
	if err != nil {
		http.Error(w, "Erro ao buscar o número de concurso", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(numConcursoController)

}
