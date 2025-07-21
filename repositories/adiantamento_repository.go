package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type AdiantamentoRepository struct {
	db *sql.DB
}

func NewAdiantamentoRepository(db *sql.DB) *AdiantamentoRepository {
	return &AdiantamentoRepository{db: db}
}

func (repository *AdiantamentoRepository) GetAllAdiantamentos() ([]models.Adiantamento, error) {
	linhas, err := repository.db.Query("SELECT solic_numprocesso, pess_matricula, adto_tipo, adto_data, adto_datalimite, adto_valor, adto_objetivo, lei_numero, decr_numero, org_codigo, unor_codigo, emp_numero FROM aplic2008.adiantamento")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var adiantamentos []models.Adiantamento
	for linhas.Next() {
		var adiantamento models.Adiantamento
		if err := linhas.Scan(
			&adiantamento.SOLIC_NumProcesso,
			&adiantamento.PESS_Matricula,
			&adiantamento.ADTO_Tipo,
			&adiantamento.ADTO_Data,
			&adiantamento.ADTO_DataLimite,
			&adiantamento.ADTO_Valor,
			&adiantamento.ADTO_Objetivo,
			&adiantamento.LEI_Numero,
			&adiantamento.DECR_Numero,
			&adiantamento.ORG_Codigo,
			&adiantamento.UNOR_Codigo,
			&adiantamento.EMP_Numero,
		); err != nil {
			return nil, err
		}
		adiantamentos = append(adiantamentos, adiantamento)
	}
	return adiantamentos, nil
}
