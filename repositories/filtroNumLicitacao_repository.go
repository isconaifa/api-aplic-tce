package repositories

import "database/sql"

type FiltroNumLicitacaoRepository struct {
	db *sql.DB
}

func NewFiltroNumLicitacaoRepository(db *sql.DB) *FiltroNumLicitacaoRepository {
	return &FiltroNumLicitacaoRepository{db: db}
}
