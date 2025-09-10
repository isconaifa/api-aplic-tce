package models

import "time"

type Liquidacao struct {
	NumLiquidacao             string     `json:"num_liquidacao"`
	Data                      time.Time  `json:"data"`
	Valor                     float64    `json:"valor"`
	Credor                    string     `json:"credor"`
	OrgaoCodigo               string     `json:"orgao_codigo"`
	UnidadeOrcamentariaCodigo string     `json:"unidade_orcamentaria_codigo"`
	NumEmpenho                string     `json:"num_empenho"`
	Itens                     int        `json:"itens"`
	DescontoLiqPago           float64    `json:"desconto_liq_pago"`
	DescontoLiquidado         float64    `json:"desconto_liquidado"`
	NotaFiscal                int        `json:"nota_fiscal"`
	Recibos                   int        `json:"recibos"`
	ValorPago                 float64    `json:"valor_pago"`
	AnulacaoLiquidacao        float64    `json:"anulacao_liquidacao"`
	Convenio                  *string    `json:"convenio,omitempty"`
	ConvenioAditivo           *string    `json:"convenio_aditivo,omitempty"`
	TipoDocumentoHabil        *string    `json:"tipo_documento_habil,omitempty"`
	DataAtesto                *time.Time `json:"data_atesto,omitempty"`
	MatriculaRespLiquidacao   *string    `json:"matricula_resp_liquidacao,omitempty"`
	MatriculaRespAtesto       *string    `json:"matricula_resp_atesto,omitempty"`
	ResponsavelLiquidacao     *string    `json:"responsavel_liquidacao,omitempty"`
	ResponsavelAtesto         *string    `json:"responsavel_atesto,omitempty"`
}
