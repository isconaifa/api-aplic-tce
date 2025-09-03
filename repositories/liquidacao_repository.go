package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type LiquidacaoRepository struct {
	db *sql.DB
}

func NewLiquidacaoRepository(db *sql.DB) *LiquidacaoRepository {
	return &LiquidacaoRepository{db: db}
}

func (repository *LiquidacaoRepository) GetAllLiquidacao(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, numEmpenho string) ([]models.Liquidacao, error) {
	rows, err := repository.db.Query(queries.Liquidacao, unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, numEmpenho)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var liquidacoes []models.Liquidacao
	for rows.Next() {
		var liq models.Liquidacao
		if err := rows.Scan(
			&liq.NumLiquidacao,
			&liq.Data,
			&liq.Valor,
			&liq.Credor,
			&liq.OrgaoCodigo,
			&liq.UnidadeOrcamentariaCodigo,
			&liq.NumEmpenho,
			&liq.Itens,
			&liq.DescontoLiqPago,
			&liq.DescontoLiquidado,
			&liq.NotaFiscal,
			&liq.Recibos,
			&liq.ValorPago,
			&liq.AnulacaoLiquidacao,
			&liq.Convenio,
			&liq.ConvenioAditivo,
			&liq.TipoDocumentoHabil,
			&liq.DataAtesto,
			&liq.MatriculaRespLiquidacao,
			&liq.MatriculaRespAtesto,
			&liq.ResponsavelLiquidacao,
			&liq.ResponsavelAtesto,
		); err != nil {
			return nil, err
		}
		liquidacoes = append(liquidacoes, liq)
	}
	return liquidacoes, nil
}
