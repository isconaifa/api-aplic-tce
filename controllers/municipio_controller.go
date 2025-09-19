package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type MunicipioController struct {
	repository *repositories.MunicipioRepository
}

func NewMunicipioController(repository *repositories.MunicipioRepository) *MunicipioController {
	return &MunicipioController{repository: repository}
}

func (controller *MunicipioController) GetMunicipios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	anoExercicio := r.URL.Query().Get("anoExercicio")
	if anoExercicio == "" {
		http.Error(w, "Parâmetro 'anoExercicio' é obrigatório", http.StatusBadRequest)
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
	municipiosRepository := repositories.NewMunicipioRepository(db)
	municipios, err := municipiosRepository.GetAllMunicipios(anoExercicio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Aconteceu um erro ao buscar municipios: %v", err), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(municipios)
}
