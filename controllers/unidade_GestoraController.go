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
	defer db.Close()

	// Chama o repositório passando munCodigo e ano
	unidadesGestorasRepository := repositories.NewUnidadeGestoraRepository(db)
	unidadesGestoras, err := unidadesGestorasRepository.GetUnidadesGestorasPorMunicipio(munCodigo, ano)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serializa para JSON e retorna
	jsonUnidadesGestoras, err := json.Marshal(unidadesGestoras)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonUnidadesGestoras)
}
