package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type ModalidadeLicitacaoController struct {
	repository *repositories.ModalidadeLicitacaoRepository
}

func NewModalidadeLicitacaoController(repository *repositories.ModalidadeLicitacaoRepository) *ModalidadeLicitacaoController {
	return &ModalidadeLicitacaoController{repository: repository}
}

func (controller *ModalidadeLicitacaoController) GetAllModalidadeLicitacao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	modalidadeLicitacaoRepository := repositories.NewModalidadeLicitacaoRepository(db)
	modalidadeLicitacao, err := modalidadeLicitacaoRepository.GetAllModalidadeLicitacao()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar modalidadeLicitacao: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(modalidadeLicitacao)
}
