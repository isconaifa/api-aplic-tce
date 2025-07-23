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

func (repository *UnidadeGestoraRepository) GetUnidadesGestorasPorMunicipio(munCodigo string) ([]models.UnidadeGestora, error) {
	query := "SELECT DISTINCT v.ent_codigo AS codigo,\n" +
		" v.ent_nome   AS nome\n" +
		" FROM vw_entidade_aplic v\n" +
		" WHERE v.mun_codigo = :1\n" +
		" ORDER BY v.ent_nome"

	rows, err := repository.db.Query(query, munCodigo)
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
