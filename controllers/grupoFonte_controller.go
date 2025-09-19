package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

type GrupoFonteController struct {
	repository repositories.GrupoFonteRepository
}

func NewGrupoFonteController(repository *repositories.GrupoFonteRepository) *GrupoFonteController {
	return &GrupoFonteController{repository: *repository}
}

func (controller *GrupoFonteController) GetAllGruposFonte(w http.ResponseWriter, r *http.Request) {
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
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	gruposFonteRepository := repositories.NewGrupoFonteRepository(db)
	gruposFonte, err := gruposFonteRepository.GetAllGruposFonte(ano)
	if err != nil {
		http.Error(w, "Erro ao buscar gruposFonte", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(gruposFonte)
}
