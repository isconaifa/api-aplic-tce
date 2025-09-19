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

type FiltroDestinacaoRecursoController struct {
	repository *repositories.FiltroDestinacaoRecursoRepository
}

func NewFiltroDestinacaoRecursoController(repository *repositories.FiltroDestinacaoRecursoRepository) *FiltroDestinacaoRecursoController {
	return &FiltroDestinacaoRecursoController{repository: repository}
}

func (controller *FiltroDestinacaoRecursoController) GetAllFiltroDestinacaoRecurso(w http.ResponseWriter, r *http.Request) {
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
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	filtroDestinacaoRecursoRepository := repositories.NewFiltroDestinacaoRecursoRepository(db)
	filtroDestinacaoRecursos, err := filtroDestinacaoRecursoRepository.GetAllFiltroDestinacaoRecurso(unidadeGestoraCodigo, ano, codigoDestinacaoRecurso)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroDestinacaoRecursos", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(filtroDestinacaoRecursos)

}
