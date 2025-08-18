package repositories

import "database/sql"

type FiltroSomenteAnuladosRepository struct {
	db *sql.DB
}

func NewFiltroSomenteAnuladosRepository(db *sql.DB) *FiltroSomenteAnuladosRepository {
	return &FiltroSomenteAnuladosRepository{db: db}
}
