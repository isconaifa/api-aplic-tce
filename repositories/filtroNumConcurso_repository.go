package repositories

import "database/sql"

type FilroNumConcursoRepository struct {
	db *sql.DB
}

func NewFiltroNumConcursoRepository(db *sql.DB) *FilroNumConcursoRepository {
	return &FilroNumConcursoRepository{db: db}
}
