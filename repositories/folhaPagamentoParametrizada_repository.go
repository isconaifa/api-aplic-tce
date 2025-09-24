package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type FolhaPagamentoParametrizadaRepository struct {
	db *sql.DB
}

func NewFolhaPagamentoParametrizadaRepository(db *sql.DB) *FolhaPagamentoParametrizadaRepository {
	return &FolhaPagamentoParametrizadaRepository{db: db}
}

func (repository *FolhaPagamentoParametrizadaRepository) GetFolhaPagamentoParametrizada(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, mesReferencia string) ([]models.FolhaPagamento, error) {
	var codigoOrgaoParam interface{}
	if codigoOrgao != "" {
		codigoOrgaoParam = codigoOrgao
	} else {
		codigoOrgaoParam = nil
	}
	var codigoUnidadeOrcamentariaParam interface{}
	if codigoUnidadeOrcamentaria != "" {
		codigoUnidadeOrcamentariaParam = codigoUnidadeOrcamentaria
	} else {
		codigoUnidadeOrcamentariaParam = nil
	}
	var mesReferenciaParam interface{}
	if mesReferencia != "" {
		mesReferenciaParam = mesReferencia
	} else {
		mesReferenciaParam = nil
	}
	rows, err := repository.db.Query(
		queries.FolhaPagamentParametrizada,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano),
		sql.Named("codigoOrgao", codigoOrgaoParam),
		sql.Named("codigoUnidadeOrcamentaria", codigoUnidadeOrcamentariaParam),
		sql.Named("mesReferencia", mesReferenciaParam),
	)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var folhasPagamentoParametrizadas []models.FolhaPagamento
	for rows.Next() {
		var folpag models.FolhaPagamento
		if err := rows.Scan(
			&folpag.AnoReferencia,
			&folpag.CodMesReferencia,
			&folpag.MesReferencia,
			&folpag.ValorBruto,
			&folpag.Gratificacoes,
			&folpag.Beneficios,
			&folpag.Descontos,
			&folpag.ValorLiquido,
			&folpag.QtdeFuncionarios); err != nil {
			return nil, err
		}
		folhasPagamentoParametrizadas = append(folhasPagamentoParametrizadas, folpag)
	}
	return folhasPagamentoParametrizadas, nil
}
