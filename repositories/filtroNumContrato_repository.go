package repositories

import "database/sql"

type FiltroNumContratoRepository struct {
	db *sql.DB
}

func NewFiltroNumContratoRepository(db *sql.DB) *FiltroNumContratoRepository {
	return &FiltroNumContratoRepository{db: db}
}
