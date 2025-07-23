package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type UnidadeGestoraController struct {
	repository *repositories.UnidadeGestoraRepository
}

func NewUnidadeGestoraController(repository *repositories.UnidadeGestoraRepository) *UnidadeGestoraController {
	return &UnidadeGestoraController{repository: repository}
}

func (controller *UnidadeGestoraController) GetUnidadesGestoras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Pega o parâmetro mun_codigo da URL
	munCodigo := r.URL.Query().Get("mun_codigo")
	if munCodigo == "" {
		http.Error(w, "Parâmetro 'mun_codigo' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	unidadesGestorasRepository := repositories.NewUnidadeGestoraRepository(db)
	unidadesGestoras, err := unidadesGestorasRepository.GetUnidadesGestorasPorMunicipio(munCodigo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUnidadesGestoras, err := json.Marshal(unidadesGestoras)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonUnidadesGestoras)
}
