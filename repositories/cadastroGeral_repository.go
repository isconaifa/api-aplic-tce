package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type CadastroGeralRepository struct {
	db *sql.DB
}

func NewCadastroGeralRepository(db *sql.DB) *CadastroGeralRepository {
	return &CadastroGeralRepository{db: db}
}

func (repository *CadastroGeralRepository) GetAllCadastroGeral(unidadeGestoraCodigo string) ([]models.CadastroGeral, error) {
	query := "select \n   " +
		"CG.CG_IDENTIFICACAO as \"CPF/CNPJ\",\n  " +
		"CG.CG_NOME as \"CREDOR\", \n   " +
		"CG.CG_UF as \"UF\",\n  " +
		"CG.CG_CIDADE as \"MUNICÍPIO\",\n  " +
		"DECODE(CG.CG_TIPOPESSOA, '1', 'FÍSICA', '2', 'JURÍDICA') as \"TIPO DE PESSOA\", \n  " +
		"E.TEMP_DESCRICAO as \"TIPO DE EMPRESA\"   \n" +
		"from aplic2008. CADASTRO_GERAL CG   \n" +
		"left join aplic2008.TIPO_EMPRESA E on (CG.CG_TIPOEMPRESA = E.TEMP_CODIGO) \n" +
		"where CG.ENT_CODIGO = :1\n" +
		"and CG.EXERCICIO >= 2015\n" +
		"order by CG.CG_NOME"

	rows, err := repository.db.Query(query, unidadeGestoraCodigo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cadastroGerais []models.CadastroGeral
	for rows.Next() {
		var cadastroGeral models.CadastroGeral
		if err := rows.Scan(
			&cadastroGeral.CpfCnpj,
			&cadastroGeral.Credor,
			&cadastroGeral.Uf,
			&cadastroGeral.Municipio,
			&cadastroGeral.TipoPessoa,
			&cadastroGeral.TipoEmpresa,
		); err != nil {
			return nil, err
		}
		cadastroGerais = append(cadastroGerais, cadastroGeral)
	}
	return cadastroGerais, nil
}
