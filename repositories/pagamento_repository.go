package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type PagamentoRepository struct {
	db *sql.DB
}

func NewPagamentoRepository(db *sql.DB) *PagamentoRepository {
	return &PagamentoRepository{db: db}
}

func (repository *PagamentoRepository) GetPagamentos(unidadeGestoraCodigo, ano, numEmpenho string) ([]models.Pagamento, error) {
	rows, err := repository.db.Query(queries.Pagamento,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano),
		sql.Named("numEmpenho", numEmpenho))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pagamentos []models.Pagamento
	for rows.Next() {
		var pag models.Pagamento
		if err := rows.Scan(
			&pag.NumPagamento,              // 1 - "Nº Pagamento"
			&pag.Data,                      // 2 - "Data"
			&pag.Valor,                     // 3 - "Valor"
			&pag.OrgaoCodigo,               // 4 - "Órgão(código)"
			&pag.UnidadeOrcamentariaCodigo, // 5 - "Unidade Orçamentária(código)"
			&pag.NumEmpenho,                // 6 - "Nº Empenho"
			&pag.UndOrcamentaria,           // 7 - "Und.Orçcamentaria"
			&pag.Orgao,                     // 8 - "Órgão"
			&pag.QtdeDocumentos,            // 9 - "Qtde Documentos"
			&pag.AnulacaoPagamento,         // 10 - "Anulação Pagamento"
			&pag.NumConvenio,               // 11 - "Nº Convênio"
			&pag.NumConvenioAdt,            // 12 - "Nº Convênio Adt"
			&pag.Justificativa,             // 13 - "Jusitificativa"
			&pag.NumLiquidacao,
		); err != nil {
			return nil, err
		}
		pagamentos = append(pagamentos, pag)
	}
	return pagamentos, nil
}
