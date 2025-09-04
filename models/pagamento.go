package models

import "time"

type Pagamento struct {
	NumPagamento              string    `json:"num_pagamento"`
	Data                      time.Time `json:"data"`
	Valor                     float64   `json:"valor"`
	OrgaoCodigo               string    `json:"orgao_codigo"`
	UnidadeOrcamentariaCodigo string    `json:"unidade_orcamentaria_codigo"`
	NumEmpenho                string    `json:"num_empenho"`
	UndOrcamentaria           string    `json:"und_orcamentaria"`
	Orgao                     string    `json:"orgao"`
	QtdeDocumentos            int       `json:"qtde_documentos"`
	AnulacaoPagamento         float64   `json:"anulacao_pagamento"`
	NumConvenio               *string   `json:"num_convenio"`
	NumConvenioAdt            *string   `json:"num_convenio_adt"`
	Justificativa             *string   `json:"justificativa"`
	NumLiquidacao             string    `json:"num_liquidacao"`
}
