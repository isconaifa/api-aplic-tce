package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fontesDestinacaoRecursoRepository := repositories.NewFonteDestinacaoRecursoRepository(db)
	fontesDestinacaoRecurso, err := fontesDestinacaoRecursoRepository.GetAllFontesDestinacaoRecurso(unidadeGestoraCodigo, ano)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonFontesDestinacaoRecurso, err := json.Marshal(fontesDestinacaoRecurso)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFontesDestinacaoRecurso)
}
