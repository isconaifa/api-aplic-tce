package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type MesRepository struct {
	db *sql.DB
}

func NewMesRepository(db *sql.DB) *MesRepository {
	return &MesRepository{db: db}
}
func (repository *MesRepository) GetMeses() ([]models.Meses, error) {
	query, err := queries.Load("Mes.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var meses []models.Meses
	for rows.Next() {
		var m models.Meses
		if err := rows.Scan(
			&m.MesReferencia,
			&m.Descricao); err != nil {
			return nil, err
		}
		meses = append(meses, m)
	}
	return meses, nil
}
