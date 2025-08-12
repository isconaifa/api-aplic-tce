package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type SubFuncaoRepository struct {
	db *sql.DB
}

func NewSubFuncaoRepository(db *sql.DB) *SubFuncaoRepository {
	return &SubFuncaoRepository{db: db}
}

func (r *SubFuncaoRepository) GetAllSubFuncoes(unidadeGestora, exercicio string, codigoFuncao int) ([]models.Subfuncao, error) {
	query := "select distinct sf.sfn_codigo, sf.sfn_descricao\n" +
		"from aplic2008.empenho e, aplic2008.subfuncao sf\n" +
		"where e.sfn_codigo = sf.sfn_codigo\n" +
		"and e.Ent_Codigo = :1\n" +
		"and e.Exercicio = :2\n" +
		"and e.fn_Codigo = :3\n" +
		"order by sf.sfn_descricao"
	rows, err := r.db.Query(query, unidadeGestora, exercicio, codigoFuncao)
	if err != nil {
		return nil, err
	}
	var subFuncoes []models.Subfuncao
	for rows.Next() {
		var subFuncao models.Subfuncao
		if err := rows.Scan(
			&subFuncao.SF_Codigo,
			&subFuncao.SF_Desc,
		); err != nil {
			return nil, err
		}
		subFuncoes = append(subFuncoes, subFuncao)
	}
	return subFuncoes, nil
}
