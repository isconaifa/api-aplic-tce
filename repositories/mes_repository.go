package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type MesRepository struct {
	db *sql.DB
}

func NewMesRepository(db *sql.DB) *MesRepository {
	return &MesRepository{db: db}
}
func (repository *MesRepository) GetMeses() ([]models.Meses, error) {
	rows, err := repository.db.Query("select comp.mes_referencia, comp.descricao\n " +
		"from aplic2008.competencia_aplic comp\n" +
		"where mes_referencia between '01' and '12'\n" +
		"order by mes_referencia")
	if err != nil {
		return nil, err
	}
	func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var meses []models.Meses
	for rows.Next() {
		var mes models.Meses
		if err := rows.Scan(
			&mes.MesReferencia,
			&mes.Descricao); err != nil {
			return nil, err
		}
		meses = append(meses, mes)
	}
	return meses, nil
}
