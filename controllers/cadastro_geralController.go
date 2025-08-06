package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type CadastroGeralController struct {
	repository *repositories.CadastroGeralRepository
}

func NewCadastroGeralController(repository *repositories.CadastroGeralRepository) *CadastroGeralController {
	return &CadastroGeralController{repository: repository}
}

func (controller *CadastroGeralController) GetAllCadastroGeral(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, "Parâmetro 'unidadeGestoraCodigo' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	cadGeralRepository := repositories.NewCadastroGeralRepository(db)
	cadastroGeral, err := cadGeralRepository.GetAllCadastroGeral(unidadeGestoraCodigo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonCadastroGeral, err := json.Marshal(cadastroGeral)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCadastroGeral)
}
