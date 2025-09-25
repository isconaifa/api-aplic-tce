package models

type Meses struct {
	MesReferencia string `json:"mes_referencia" db:"mes_referencia"`
	Descricao     string `json:"descricao" db:"descricao"`
}
