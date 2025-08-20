package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type ModalidadeLicitacaoRepository struct {
	db *sql.DB
}

func NewModalidadeLicitacaoRepository(db *sql.DB) *ModalidadeLicitacaoRepository {
	return &ModalidadeLicitacaoRepository{db: db}
}

func (repository *ModalidadeLicitacaoRepository) GetAllModalidadeLicitacao() ([]models.ModalidadeLicitacao, error) {
	rows, err := repository.db.Query(queries.ModalidadeLicitacaoQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var modalidadeLicitacoes []models.ModalidadeLicitacao
	for rows.Next() {
		var m models.ModalidadeLicitacao
		if err := rows.Scan(&m.Mlic_codigo, &m.Mlic_descricao); err != nil {
			return nil, err
		}
		modalidadeLicitacoes = append(modalidadeLicitacoes, m)
	}
	return modalidadeLicitacoes, nil
}
