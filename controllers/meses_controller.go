package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

type MesesController struct {
	Repository *repositories.MesRepository
}

func NewMesesController(repository *repositories.MesRepository) *MesesController {
	return &MesesController{Repository: repository}
}

func (controller *MesesController) GetMeses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	mesesRepository := repositories.NewMesRepository(db)
	meses, err := mesesRepository.GetMeses()
	//fmt.Println(meses)
	if err != nil {
		http.Error(w, "Erro ao buscar meses", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(meses)
}
