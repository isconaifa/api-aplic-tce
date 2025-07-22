package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type UnidadeGestoraRepository struct {
	db *sql.DB
}

func NewUnidadeGestoraRepository(db *sql.DB) *UnidadeGestoraRepository {
	return &UnidadeGestoraRepository{db: db}
}
func (repository *UnidadeGestoraRepository) GetAllUnidadesGestoras() ([]models.UnidadeGestora, error) {
	rows, err := repository.db.Query(" select distinct v.ent_codigo as codigo,\n" +
		"v.ent_nome as nome\n " +
		"from vw_entidade_aplic v\n " +
		"where v.mun_codigo = '510665'\n " +
		"order by v.ent_nome")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var unidadesGestoras []models.UnidadeGestora

	for rows.Next() {
		var unidadeGestora models.UnidadeGestora
		if err := rows.Scan(&unidadeGestora.Codigo, &unidadeGestora.Nome); err != nil {
			return nil, err
		}
		unidadesGestoras = append(unidadesGestoras, unidadeGestora)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return unidadesGestoras, nil
}
