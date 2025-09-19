package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

type FonteDestinacaoRecursoController struct {
	repository *repositories.FonteDestinacaoRecursoRepository
}

func NewFonteDestinacaoRecursoController(repository *repositories.FonteDestinacaoRecursoRepository) *FonteDestinacaoRecursoController {
	return &FonteDestinacaoRecursoController{repository: repository}

}

func (controller *FonteDestinacaoRecursoController) GetFontesDestinacaoRecurso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, "Parâmetro 'unidadeGestoraCodigo' é obrigatório", http.StatusBadRequest)
		return
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, "Parâmetro 'exercicio' é obrigatório", http.StatusBadRequest)
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
	fontesDestinacaoRecursoRepository := repositories.NewFonteDestinacaoRecursoRepository(db)
	fontesDestinacaoRecurso, err := fontesDestinacaoRecursoRepository.GetAllFontesDestinacaoRecurso(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, "Erro ao buscar fontesDestinacaoRecurso", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(fontesDestinacaoRecurso)
}
