package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
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
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	cadGeralRepository := repositories.NewCadastroGeralRepository(db)
	cadastroGeral, err := cadGeralRepository.GetAllCadastroGeral(unidadeGestoraCodigo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar cadastro geral: %v", err), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(cadastroGeral)
}
