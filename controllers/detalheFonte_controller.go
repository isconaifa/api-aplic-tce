package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type DetalheFonteController struct {
	repository *repositories.DetalhefonteRepository
}

func NewDetalheFonteController(repository *repositories.DetalhefonteRepository) *DetalheFonteController {
	return &DetalheFonteController{repository: repository}
}
func (controller *DetalheFonteController) GetAllDetalheFonte(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	defer db.Close()
	detalheFonteRepository := repositories.NewDetalhefonteRepository(db)
	detalheFonte, err := detalheFonteRepository.GetAllDetalhefonte(ano)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonDetalheFonte, err := json.Marshal(detalheFonte)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonDetalheFonte)
}
