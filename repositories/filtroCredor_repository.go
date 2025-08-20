package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FiltroCredorRepository struct {
	db *sql.DB
}

func NewFiltroCredorRepository(db *sql.DB) *FiltroCredorRepository {
	return &FiltroCredorRepository{db: db}
}

func (repository *FiltroCredorRepository) GetFiltroCredor(unidadeGestoraCodigo, ano, codigoCredor string) ([]models.FiltroCredor, error) {
	return nil, nil
}
