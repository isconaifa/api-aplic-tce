package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
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
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	tiposDeCargasRepository := repositories.NewTipoDeCargaRepository(db)
	tiposDeCargas, err := tiposDeCargasRepository.GetAllTiposDeCarga()
	if err != nil {
		http.Error(w, "Erro ao buscar Tipos de Carga", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(tiposDeCargas)
}
