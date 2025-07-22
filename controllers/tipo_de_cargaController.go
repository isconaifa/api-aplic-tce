package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
)

type TipoDeCargaController struct {
	repository *repositories.TipoDeCargaRepository
}

func NewTipoDeCargaController(repository *repositories.TipoDeCargaRepository) *TipoDeCargaController {
	return &TipoDeCargaController{repository: repository}
}

func (controller *TipoDeCargaController) GetTiposDeCargas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	tiposDeCargasRepository := repositories.NewTipoDeCargaRepository(db)
	tiposDeCargas, err := tiposDeCargasRepository.GetAllTiposDeCarga()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonTiposDeCargas, err := json.Marshal(tiposDeCargas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTiposDeCargas)
}
