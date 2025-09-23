package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type AdiantamentoRepository struct {
	db *sql.DB
}

func NewAdiantamentoRepository(db *sql.DB) *AdiantamentoRepository {
	return &AdiantamentoRepository{db: db}
}

func (repository *AdiantamentoRepository) GetAdiantamento(unidadeGestoraCodigo, ano string) ([]models.Adiantamento, error) {
	rows, err := repository.db.Query(
		queries.Adiantamento,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var adiantamentos []models.Adiantamento
	for rows.Next() {
		var ad models.Adiantamento
		if err := rows.Scan(
			&ad.NumProcesso,
			&ad.Matricula,
			&ad.CodigoTipo,
			&ad.Tipo,
			&ad.DataDoc,
			&ad.ValorAdiantamento,
			&ad.Objetivo,
			&ad.NumeroLei,
			&ad.Orgao,
			&ad.CodigoUnidadeOrcamentaria,
			&ad.Empenho,
			&ad.DataLimite,
			&ad.NumeroDecreto,
			&ad.Funcionario,
			&ad.DescricaoOrgao,
			&ad.UnidadeOrcamentaria,
			&ad.DataPrestacaoContas,
			&ad.SituacaoPrestacaoContas,
			&ad.PCADAprovada,
			&ad.QtdAdSemPrestacao,
			&ad.PrestacaoContas,
			&ad.AdDevolucao); err != nil {
			return nil, err
		}
		adiantamentos = append(adiantamentos, ad)
	}
	return adiantamentos, nil
}
