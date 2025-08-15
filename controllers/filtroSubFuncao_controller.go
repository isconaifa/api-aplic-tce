package controllers

import "api-aplic-web/repositories"

type FiltroSubFuncaoController struct {
	repository *repositories.FiltroSubFuncaoRepository
}

func NewFiltroSubFuncaoController(repository *repositories.FiltroSubFuncaoRepository) *FiltroSubFuncaoController {
	return &FiltroSubFuncaoController{repository: repository}
}
