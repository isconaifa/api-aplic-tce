package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type TipoDeCargaRepository struct {
	db *sql.DB
}

func NewTipoDeCargaRepository(db *sql.DB) *TipoDeCargaRepository {
	return &TipoDeCargaRepository{db: db}
}

func (repository *TipoDeCargaRepository) GetAllTiposDeCarga() ([]models.TipoCarga, error) {
	rows, err := repository.db.Query("SELECT ca_descricao\n " +
		"FROM aplic2008.cargas_aplic\n" +
		"WHERE ca_codigo IN ('CT', 'FP', 'PA', 'CC')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tipoCargas []models.TipoCarga
	for rows.Next() {
		var tipoCarga models.TipoCarga
		if err := rows.Scan(&tipoCarga.Descricao); err != nil {
			return nil, err
		}
		tipoCargas = append(tipoCargas, tipoCarga)
	}
	return tipoCargas, nil
}
