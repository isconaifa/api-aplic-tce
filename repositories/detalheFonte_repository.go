package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type DetalhefonteRepository struct {
	db *sql.DB
}

func NewDetalhefonteRepository(db *sql.DB) *DetalhefonteRepository {
	return &DetalhefonteRepository{db: db}
}

func (repository *DetalhefonteRepository) GetAllDetalhefonte(ano string) ([]models.DetalheFonte, error) {
	query := "select d.destrec_codigo, d.destrec_descricao\n" +
		"from aplic2008.destinacao_recurso d\n" +
		"where d.exercicio = :1\n"

	rows, err := repository.db.Query(query, ano)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var detalheFontes []models.DetalheFonte
	for rows.Next() {
		var detalheFonte models.DetalheFonte
		if err := rows.Scan(&detalheFonte.Destrect_codigo, &detalheFonte.Destrect_descricao); err != nil {
			return nil, err
		}
		detalheFontes = append(detalheFontes, detalheFonte)
	}
	return detalheFontes, nil
}
