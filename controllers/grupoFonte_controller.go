package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
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
	defer db.Close()
	gruposFonteRepository := repositories.NewGrupoFonteRepository(db)
	gruposFonte, err := gruposFonteRepository.GetAllGruposFonte(ano)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonGruposFonte, err := json.Marshal(gruposFonte)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonGruposFonte)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
