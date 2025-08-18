package repositories

import "database/sql"

type FiltroValorEmpenhadoRepository struct {
	db *sql.DB
}

func NewFiltroValorEmpenhadoRepository(db *sql.DB) *FiltroValorEmpenhadoRepository {
	return &FiltroValorEmpenhadoRepository{db: db}
}
