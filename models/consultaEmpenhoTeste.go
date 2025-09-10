package models

type EmpenhoTeste struct {
	Emp_Data           string  `json:"empenho_data"`
	Emp_Numero         string  `json:"empenho_numero"`
	Credor             string  `json:"credor"`
	Valor_Empenhado    float64 `json:"valor_empenhado"`
	Valor_Liquidado    float64 `json:"valor_liquidado"`
	Valor_Pago         float64 `json:"valor_pago"`
	Valor_PagoRetencao float64 `json:"valor_pago_mais_retencao"`
	Anulado_Empenho    float64 `json:"valor_anulado"`
	Qtde_Notas_Fiscais string  `json:"qtde_notas_fiscais"`
	Qtde_NFe           string  `json:"qtde_nfe"`
	Contratos          string  `json:"contratos"`
}
