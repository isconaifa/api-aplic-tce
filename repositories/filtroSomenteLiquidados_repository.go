package repositories

import "database/sql"

type FiltroSomenteLiquidadosRepository struct {
	db *sql.DB
}

func NewFiltroSomenteLiquidadosRepository(db *sql.DB) *FiltroSomenteLiquidadosRepository {
	return &FiltroSomenteLiquidadosRepository{db: db}
}
