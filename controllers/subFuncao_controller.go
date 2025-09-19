package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type SubFuncaoController struct {
	repository *repositories.SubFuncaoRepository
}

func NewSubFuncaoController(repository *repositories.SubFuncaoRepository) *SubFuncaoController {
	return &SubFuncaoController{repository: repository}
}

func (controller *SubFuncaoController) GetAllSubFuncoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	unidadeGestoraCodigo := r.URL.Query().Get("unidadeGestoraCodigo")
	if unidadeGestoraCodigo == "" {
		http.Error(w, "Parâmetro 'unidadeGestora' é obrigatório", http.StatusBadRequest)
		return
	}
	ano := r.URL.Query().Get("ano")
	if ano == "" {
		http.Error(w, "Parâmetro 'exercicio' é obrigatório", http.StatusBadRequest)
		return
	}
	codigoFuncao, err := strconv.Atoi(r.URL.Query().Get("codigoFuncao"))
	if err != nil {
		http.Error(w, "Parâmetro 'codigoFuncao' é obrigatório", http.StatusBadRequest)
		return
	}

	if codigoFuncao == 0 {
		http.Error(w, "Parâmetro 'codigoFuncao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	subFuncoesRepository := repositories.NewSubFuncaoRepository(db)
	subFuncoes, err := subFuncoesRepository.GetAllSubFuncoes(unidadeGestoraCodigo, ano, codigoFuncao)
	if err != nil {
		http.Error(w, "Erro ao buscar Sub Função", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(subFuncoes)
}
