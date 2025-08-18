package repositories

import "database/sql"

type FiltroOpcaoRepository struct {
	db *sql.DB
}

func NewFiltroOpcaoRepository(db *sql.DB) *FiltroOpcaoRepository {
	return &FiltroOpcaoRepository{db: db}
}
