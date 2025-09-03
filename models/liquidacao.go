package models

import "time"

type Liquidacao struct {
	NumLiquidacao             string     `json:"Nº da Liquidação"`
	Data                      time.Time  `json:"Data"`
	Valor                     float64    `json:"Valor"`
	Credor                    string     `json:"Credor"`
	OrgaoCodigo               string     `json:"Órgão(código)"`
	UnidadeOrcamentariaCodigo string     `json:"Unidade Orçamentária(código)"`
	NumEmpenho                string     `json:"Nº Empenho"`
	Itens                     int        `json:"Item(ns)"`
	DescontoLiqPago           float64    `json:"Desconto Liq.Pago"`
	DescontoLiquidado         float64    `json:"Desconto Liquidado"`
	NotaFiscal                int        `json:"Nota Fiscal"`
	Recibos                   int        `json:"Recibo(s)"`
	ValorPago                 float64    `json:"Valor Pago"`
	AnulacaoLiquidacao        float64    `json:"Anulação Liquidação"`
	Convenio                  *string    `json:"Convênio"`
	ConvenioAditivo           *string    `json:"Convênio Aditivo"`
	TipoDocumentoHabil        *string    `json:"Tipo Documento Hábil"`
	DataAtesto                *time.Time `json:"Data do atesto"`
	MatriculaRespLiquidacao   *string    `json:"Matr. do resp. pela liquidação"`
	MatriculaRespAtesto       *string    `json:"Matr. do resp. pelo atesto"`
	ResponsavelLiquidacao     *string    `json:"Responsável pela liquidação"`
	ResponsavelAtesto         *string    `json:"Responsável pelo atesto"`
}
