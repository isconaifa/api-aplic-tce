package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type ServidorRepository struct {
	db *sql.DB
}

func NewServidorRepository(db *sql.DB) *ServidorRepository {
	return &ServidorRepository{db: db}
}

func (repository *ServidorRepository) GetServidores(unidadeGestoraCodigo, ano string) ([]models.Servidor, error) {
	rows, err := repository.db.Query(queries.ServidorQuery,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var servidores []models.Servidor
	for rows.Next() {
		var serv models.Servidor
		if err := rows.Scan(
			&serv.Matricula,
			&serv.Nome); err != nil {
			return nil, err
		}
		servidores = append(servidores, serv)
	}
	return servidores, nil
}
