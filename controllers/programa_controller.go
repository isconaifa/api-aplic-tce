package controllers

import (
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProgramaController struct {
	repository *repositories.ProgramaRepository
}

func NewProgramaController(repository *repositories.ProgramaRepository) *ProgramaController {
	return &ProgramaController{repository: repository}
}

func (controller *ProgramaController) GetAllProgramas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	unidadeGestora := r.URL.Query().Get("unidadeGestora")
	if unidadeGestora == "" {
		http.Error(w, "Parâmetro 'unidadeGestora' é obrigatório", http.StatusBadRequest)
		return
	}
	exercicio := r.URL.Query().Get("exercicio")
	if exercicio == "" {
		http.Error(w, "Parâmetro 'exercicio' é obrigatório", http.StatusBadRequest)
		return
	}
	codigoFuncao, err := strconv.Atoi(r.URL.Query().Get("codigoFuncao"))
	if codigoFuncao == 0 {
		http.Error(w, "Parâmetro 'codigoFuncao' é obrigatório", http.StatusBadRequest)
		return
	}
	db, err := database.Connectdb()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	programaRepository := repositories.NewProgramaRepository(db)
	programas, err := programaRepository.GetAllProgramas(unidadeGestora, exercicio, codigoFuncao)
	if err != nil {
		http.Error(w, "Erro ao buscar programas", http.StatusInternalServerError)
		return
	}
	jsonProgramas, err := json.Marshal(programas)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonProgramas)
}
