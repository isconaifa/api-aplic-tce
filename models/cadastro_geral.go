package models

type CadastroGeral struct {
	CpfCnpj     *string `json:"cpf_cnpj"`
	Credor      *string `json:"credor"`
	Uf          *string `json:"uf"`
	Municipio   *string `json:"municipio"`
	TipoPessoa  *string `json:"tipo_pessoa"`
	TipoEmpresa *string `json:"tipo_empresa"`
}
