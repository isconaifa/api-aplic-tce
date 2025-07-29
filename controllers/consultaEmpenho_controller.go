package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type ConsultaEmpenhoController struct {
	repository *repositories.ConsultaEmpenhoRepository
}

func NewConsultaEmpenhoController(repository *repositories.ConsultaEmpenhoRepository) *ConsultaEmpenhoController {
	return &ConsultaEmpenhoController{repository: repository}
}
func (controller *ConsultaEmpenhoController) GetAllConsultaEmpenhos(w http.ResponseWriter, r *http.Request) {
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
	db, err := database.Connectdb()
	consultaEmpenhoRepository := repositories.NewConsultaEmpenhoRepository(db)
	consultaEmpenhos, err := consultaEmpenhoRepository.GetAllConsultaEmpenhos(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonConsultaEmpenhos, err := json.Marshal(consultaEmpenhos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonConsultaEmpenhos)
}
