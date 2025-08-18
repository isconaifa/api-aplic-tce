package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FiltroGrupoFonteController struct {
	repository *repositories.FiltroGrupoFonteRepository
}

func NewFiltroGrupoFonteController(repository *repositories.FiltroGrupoFonteRepository) *FiltroGrupoFonteController {
	return &FiltroGrupoFonteController{repository: repository}
}
func (controller *FiltroGrupoFonteController) GetAllFiltroGrupoFonte(w http.ResponseWriter, r *http.Request) {
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
	codigoGrupoFonte, err := strconv.Atoi(r.URL.Query().Get("codigoGrupoFonte"))
	if err != nil {
		http.Error(w, "ocorreu um erro", http.StatusBadRequest)
		return
	}
	if codigoGrupoFonte == 0 {
		http.Error(w, "Parâmetro 'codigoGrupoFonte' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroGrupoFonteRepository := repositories.NewFiltroGrupoFonteRepository(db)
	filtroGrupoFontes, err := filtroGrupoFonteRepository.GetAllFiltroGrupoFonte(unidadeGestoraCodigo, ano, codigoGrupoFonte)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroGrupoFontes", http.StatusInternalServerError)
		return
	}
	jsonFiltroGrupoFontes, err := json.Marshal(filtroGrupoFontes)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroGrupoFontes)
}
