package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type OrgaoController struct {
	repository *repositories.OrgaoRepository
}

func NewOrgaoController(repository *repositories.OrgaoRepository) *OrgaoController {
	return &OrgaoController{repository: repository}
}

func (controller *OrgaoController) GetAllOrgaos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestora := r.URL.Query().Get("unidadeGestora")
	if unidadeGestora == "" {
		http.Error(w, "Parâmetro 'unidadeGestora' é obrigatório", http.StatusBadRequest)
		return
	}
	exercicio := r.URL.Query().Get("exercicio")
	if exercicio == "" {
		http.Error(w, "Parâmetro 'exercicio' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	orgaoRepository := repositories.NewOrgaoRepository(db)
	orgaos, err := orgaoRepository.GetAllOrgaos(unidadeGestora, exercicio)
	if err != nil {
		http.Error(w, "Erro ao buscar orgaos", http.StatusInternalServerError)
		return
	}
	jsonOrgaos, err := json.Marshal(orgaos)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonOrgaos)
}
