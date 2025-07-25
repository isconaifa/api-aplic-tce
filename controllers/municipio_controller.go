package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	municipiosRepository := repositories.NewMunicipioRepository(db)
	municipios, err := municipiosRepository.GetAllMunicipios(anoExercicio)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonMunicipios, err := json.Marshal(municipios)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMunicipios)
}
