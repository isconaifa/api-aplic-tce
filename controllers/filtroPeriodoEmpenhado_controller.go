package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type FiltroPeriodoEmpenhadoController struct {
	repository *repositories.FiltroPeriodoEmpenhadoRepository
}

func NewFiltroPeriodoEmpenhadoController(repository *repositories.FiltroPeriodoEmpenhadoRepository) *FiltroPeriodoEmpenhadoController {
	return &FiltroPeriodoEmpenhadoController{repository: repository}
}

func (controller *FiltroPeriodoEmpenhadoController) GetAllFiltroPeriodoEmpenhado(w http.ResponseWriter, r *http.Request) {
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
	dataInicioStr := r.URL.Query().Get("dataInicioStr")
	if dataInicioStr == "" {
		http.Error(w, "Parâmetro 'dataInicioStr' é obrigatório", http.StatusBadRequest)
		return
	}
	dataFimStr := r.URL.Query().Get("dataFimStr")
	if dataFimStr == "" {
		http.Error(w, "Parâmetro 'dataFimStr' é obrigatório", http.StatusBadRequest)
		return
	}
	if dataFimStr < dataInicioStr {
		http.Error(w, " Data final tem que ser maior que data inicial", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filtroPeriodoEmpenhadoRepository := repositories.NewFiltroPeriodoEmpenhadoRepository(db)
	filtroPeriodoEmpenhados, err := filtroPeriodoEmpenhadoRepository.GetAllFiltroPeriodoEmpenhado(unidadeGestoraCodigo, ano, dataInicioStr, dataFimStr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar filtroPeriodoEmpenhados", http.StatusInternalServerError)
		return
	}
	jsonfilterPeriodoEmpenhados, err := json.Marshal(filtroPeriodoEmpenhados)
	if err != nil {
		http.Error(w, "Erro ao converter filtroPeriodoEmpenhados para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonfilterPeriodoEmpenhados)
}
