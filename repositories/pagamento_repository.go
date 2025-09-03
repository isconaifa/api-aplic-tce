package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type PagamentoRepository struct {
	db *sql.DB
}

func NewPagamentoRepository(db *sql.DB) *PagamentoRepository {
	return &PagamentoRepository{db: db}
}

func (repository *PagamentoRepository) GetPagamentos() ([]models.Pagamento, error) {
	return nil, nil
}
