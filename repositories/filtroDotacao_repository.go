package repositories

import "database/sql"

type FiltroDotacaoRepository struct {
	db *sql.DB
}

func NewFiltroDotacaoRepository(db *sql.DB) *FiltroDotacaoRepository {
	return &FiltroDotacaoRepository{db: db}
}
