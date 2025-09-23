package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type FolhaPagamentoRepository struct {
	db *sql.DB
}

func NewFolhaPagamentoRepository(db *sql.DB) *FolhaPagamentoRepository {
	return &FolhaPagamentoRepository{db: db}
}

func (repository *FolhaPagamentoRepository) GetFolhaPagamento(unidadeGestoraCodigo string, ano string) ([]models.FolhaPagamento, error) {
	rows, err := repository.db.Query(
		queries.FolhaPagamento,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var folhasPagamentos []models.FolhaPagamento
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
		folhasPagamentos = append(folhasPagamentos, folpag)
	}
	return folhasPagamentos, nil
}
