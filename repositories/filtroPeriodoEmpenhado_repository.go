package repositories

import "database/sql"

type FiltroPeriodoEmpenhadoRepository struct {
	db *sql.DB
}

func NewFiltroPeriodoEmpenhadoRepository(db *sql.DB) *FiltroPeriodoEmpenhadoRepository {
	return &FiltroPeriodoEmpenhadoRepository{db: db}
}
