package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
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

	// Pega o parâmetro mun_codigo da URL
	munCodigo := r.URL.Query().Get("mun_codigo")
	if munCodigo == "" {
		http.Error(w, "Parâmetro 'mun_codigo' é obrigatório", http.StatusBadRequest)
		return
	}

	// Pega o parâmetro ano da URL
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, "Parâmetro 'ano' é obrigatório", http.StatusBadRequest)
		return
	}

	// Conecta ao banco
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// Chama o repositório passando munCodigo e ano
	unidadesGestorasRepository := repositories.NewUnidadeGestoraRepository(db)
	unidadesGestoras, err := unidadesGestorasRepository.GetUnidadesGestorasPorMunicipio(munCodigo, ano)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(unidadesGestoras)
}
