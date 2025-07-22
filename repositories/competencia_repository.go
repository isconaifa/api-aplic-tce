package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type CompetenciaRepository struct {
	db *sql.DB
}

func NewCompetenciaRepository(db *sql.DB) *CompetenciaRepository {
	return &CompetenciaRepository{db: db}
}

func (repository *CompetenciaRepository) GetAllCompetencias() ([]models.Competencia, error) {
	rows, err := repository.db.Query("select descricao\n " +
		"from aplic2008.competencia_aplic\n" +
		" where mes_referencia between '01' and '12'\n" +
		" order by mes_referencia")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competencias []models.Competencia

	for rows.Next() {
		var competencia models.Competencia
		if err := rows.Scan(&competencia.Descricao); err != nil {
			return nil, err
		}
		competencias = append(competencias, competencia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return competencias, nil
}
