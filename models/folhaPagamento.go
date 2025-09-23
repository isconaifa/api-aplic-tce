package models

type FolhaPagamento struct {
	AnoReferencia    int     `json:"ano_referencia"`
	CodMesReferencia int     `json:"cod_mes_referencia"`
	MesReferencia    string  `json:"mes_referencia"`
	ValorBruto       float64 `json:"valor_bruto"`
	Gratificacoes    float64 `json:"gratificacoes"`
	Beneficios       float64 `json:"beneficios"`
	Descontos        float64 `json:"descontos"`
	ValorLiquido     float64 `json:"valor_liquido"`
	QtdeFuncionarios int     `json:"qtde_funcionarios"`
}
