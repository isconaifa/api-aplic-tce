package models

type ModalidadeLicitacao struct {
	Mlic_codigo    int    `json:"codigo_modalidade"`
	Mlic_descricao string `json:"modalidade_licitacao"`
}
