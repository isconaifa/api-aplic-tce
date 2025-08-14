package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if codigoFuncao == 0 {
		http.Error(w, "Parâmetro 'codigoFuncao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	subFuncoesRepository := repositories.NewSubFuncaoRepository(db)
	subFuncoes, err := subFuncoesRepository.GetAllSubFuncoes(unidadeGestoraCodigo, ano, codigoFuncao)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonSubFuncoes, err := json.Marshal(subFuncoes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonSubFuncoes)
}
