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
	rows, err := repository.db.Query("select * from aplic2008.tipo_carga_aplic")
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
