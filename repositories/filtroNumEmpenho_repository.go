package repositories

import "database/sql"

type FiltroNumEmpenhoRepository struct {
	db *sql.DB
}

func NewFiltroNumEmpenhoRepository(db *sql.DB) *FiltroNumEmpenhoRepository {
	return &FiltroNumEmpenhoRepository{db: db}
}
