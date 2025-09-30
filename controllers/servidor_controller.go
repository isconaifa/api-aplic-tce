package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ServidorController struct {
	repository *repositories.ServidorRepository
}

func NewServidorController(repository *repositories.ServidorRepository) *ServidorController {
	return &ServidorController{repository: repository}
}

func (controller *ServidorController) GetServidores(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	servidorRepository := repositories.NewServidorRepository(db)
	servidores, err := servidorRepository.GetServidores(unidadeGestoraCodigo, ano)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar servidores: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(servidores)
}
