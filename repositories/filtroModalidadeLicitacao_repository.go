package repositories

import "database/sql"

type FiltroModalidadeLicitacaoRepository struct {
	db *sql.DB
}

func NewFiltroModalidadeLicitacaoRepository(db *sql.DB) *FiltroModalidadeLicitacaoRepository {
	return &FiltroModalidadeLicitacaoRepository{db: db}
}
