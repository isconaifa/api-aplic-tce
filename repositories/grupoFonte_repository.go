package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type GrupoFonteRepository struct {
	db *sql.DB
}

func NewGrupoFonteRepository(db *sql.DB) *GrupoFonteRepository {
	return &GrupoFonteRepository{db: db}
}

func (r *GrupoFonteRepository) GetAllGruposFonte(ano string) ([]models.GrupoFonte, error) {
	query := "select d.drgrp_codigo, d.drgrp_descricao\n" +
		"from aplic2008.destinacao_recurso_grupo d\n" +
		"where d.exercicio = :1\n"

	rows, err := r.db.Query(query, ano)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var gruposFonte []models.GrupoFonte
	for rows.Next() {
		var grupoFonte models.GrupoFonte
		if err := rows.Scan(
			&grupoFonte.Drgrp_codigo,
			&grupoFonte.Drgrp_descricao,
		); err != nil {
			return nil, err
		}
		gruposFonte = append(gruposFonte, grupoFonte)
	}
	return gruposFonte, nil
}
