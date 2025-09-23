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
	rows, err := repository.db.Query(
		queries.FolhaPagamentoParametrizada,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano),
		sql.Named("codigoOrgao", codigoOrgao),
		sql.Named("codigoUnidadeOrcamentaria", codigoUnidadeOrcamentaria),
		sql.Named("mesReferencia", mesReferencia))
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
