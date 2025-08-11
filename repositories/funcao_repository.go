package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FuncaoRepository struct {
	db *sql.DB
}

func NewFuncaoRepository(db *sql.DB) *FuncaoRepository {
	return &FuncaoRepository{db: db}
}

func (repository *FuncaoRepository) GetAllFuncoes(unidadeGestora, exercicio string) ([]models.Funcao, error) {
	query := "select distinct f.fn_codigo, f.fn_descricao\n" +
		"from aplic2008.empenho e, aplic2008.funcao f\n" +
		"where e.fn_codigo = f.fn_codigo\n" +
		"and e.Ent_Codigo = :1\n " +
		"and e.Exercicio = :2\n" +
		"order by f.fn_descricao"
	rows, err := repository.db.Query(query, unidadeGestora, exercicio)
	if err != nil {
		return nil, err
	}
	var funcoes []models.Funcao
	for rows.Next() {
		var funcao models.Funcao
		if err := rows.Scan(&funcao.FN_Codigo, &funcao.FN_Desc); err != nil {
			return nil, err
		}
		funcoes = append(funcoes, funcao)
	}
	return funcoes, nil
}
