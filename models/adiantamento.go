package models

type Adiantamento struct {
	SOLIC_NumProcesso *string `json:"solic_numprocesso"`
	PESS_Matricula    *string `json:"pess_matricula"`
	ADTO_Tipo         *string `json:"adto_tipo"`
	ADTO_Data         *string `json:"adto_data"`
	ADTO_DataLimite   *string `json:"adto_datalimite"`
	ADTO_Valor        *string `json:"adto_valor"`
	ADTO_Objetivo     *string `json:"adto_objetivo"`
	LEI_Numero        *string `json:"lei_numero"`
	DECR_Numero       *string `json:"decr_numero"`
	ORG_Codigo        *string `json:"org_codigo"`
	UNOR_Codigo       *string `json:"unor_codigo"`
	EMP_Numero        *string `json:"emp_numero"`
}
