package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FiltroProgramaController struct {
	repository *repositories.FiltroProgramaRepository
}

func NewFiltroProgramaController(repository *repositories.FiltroProgramaRepository) *FiltroProgramaController {
	return &FiltroProgramaController{repository: repository}
}
func (controller *FiltroProgramaController) GetAllProgramas(w http.ResponseWriter, r *http.Request) {
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
	codigoPrograma, err := strconv.Atoi(r.URL.Query().Get("codigoPrograma"))
	if err != nil {
		http.Error(w, "ocorreu um erro", http.StatusBadRequest)
		return
	}
	if codigoPrograma == 0 {
		http.Error(w, "Parâmetro 'codigoPrograma' é obrigatório", http.StatusBadRequest)
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
	filtroProgramaRepository := repositories.NewFiltroProgramaRepository(db)
	filtroProgramas, err := filtroProgramaRepository.GetAllProgramas(unidadeGestoraCodigo, ano, codigoPrograma, codigoGrupoFonte)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao buscar programas", http.StatusInternalServerError)
		return
	}
	jsonFiltroProgramas, err := json.Marshal(filtroProgramas)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFiltroProgramas)

}
