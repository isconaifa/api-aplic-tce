package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type UnidadeOrcamentariaRepository struct {
	db *sql.DB
}

func NewUnidadeOrcamentariaRepository(db *sql.DB) *UnidadeOrcamentariaRepository {
	return &UnidadeOrcamentariaRepository{db: db}
}

func (repository *UnidadeOrcamentariaRepository) GetAllUnidadeOrcamentaria(unidadeGestora, exercicio, codigoOrgao string) ([]models.UnidadeOrcamentaria, error) {
	query := "select distinct u.unor_codigo, u.unor_nome\n" +
		"from aplic2008.empenho e, aplic2008.unidade_orcamentaria u\n" +
		"where e.unor_codigo = u.unor_codigo\n" +
		"and e.org_codigo = u.org_codigo\n" +
		"and e.ent_codigo = u.ent_codigo\n" +
		"and e.exercicio = u.exercicio\n" +
		"and e.Ent_Codigo = :1\n" +
		"and e.Exercicio = :2\n" +
		"and e.org_codigo = :3\n" +
		"order by u.unor_nome"

	rows, err := repository.db.Query(query, unidadeGestora, exercicio, codigoOrgao)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var unidadesOrcamentarias []models.UnidadeOrcamentaria
	for rows.Next() {
		var unidadeOrcamentaria models.UnidadeOrcamentaria
		if err := rows.Scan(
			&unidadeOrcamentaria.Unor_Codigo,
			&unidadeOrcamentaria.Unor_Nome,
		); err != nil {
			return nil, err
		}
		unidadesOrcamentarias = append(unidadesOrcamentarias, unidadeOrcamentaria)
	}
	return unidadesOrcamentarias, nil
}
