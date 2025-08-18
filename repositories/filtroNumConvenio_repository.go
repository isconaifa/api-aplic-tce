package repositories

import "database/sql"

type FiltroNumConvenioRepository struct {
	db *sql.DB
}

func NewFiltroNumConvenioRepository(db *sql.DB) *FiltroNumConvenioRepository {
	return &FiltroNumConvenioRepository{db: db}
}
