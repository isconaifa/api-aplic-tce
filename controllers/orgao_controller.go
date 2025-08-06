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
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	orgaosRepository := repositories.NewOrgaoRepository(db)
	orgaos, err := orgaosRepository.GetAllOrgaos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonOrgaos, err := json.Marshal(orgaos)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonOrgaos)
}
