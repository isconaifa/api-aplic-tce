package models

type Adiantamento struct {
	NumProcesso               string  `json:"num_processo"`
	Matricula                 string  `json:"matricula"`
	CodigoTipo                string  `json:"codigo_tipo"`
	Tipo                      string  `json:"tipo"`
	DataDoc                   string  `json:"data_doc"`
	ValorAdiantamento         float64 `json:"valor_adiantamento"`
	Objetivo                  string  `json:"objetivo"`
	NumeroLei                 string  `json:"numero_lei"`
	Orgao                     string  `json:"orgao"`
	CodigoUnidadeOrcamentaria string  `json:"codigo_unidade_orcamentaria"`
	Empenho                   string  `json:"empenho"`
	DataLimite                string  `json:"data_limite"`
	NumeroDecreto             *string `json:"numero_decreto,omitempty"`
	Funcionario               *string `json:"funcionario,omitempty"`
	DescricaoOrgao            *string `json:"descricao_orgao"`
	UnidadeOrcamentaria       *string `json:"unidade_orcamentaria"`
	DataPrestacaoContas       *string `json:"data_prestacao_contas,omitempty"`
	SituacaoPrestacaoContas   *string `json:"situacao_prestacao_contas,omitempty"`
	PCADAprovada              *string `json:"pcad_aprovada,omitempty"`
	QtdAdSemPrestacao         *int    `json:"qtd_ad_sem_prestacao"`
	PrestacaoContas           *int    `json:"prestacao_contas"`
	AdDevolucao               *int    `json:"ad_devolucao"`
}
