package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type ExercicioRepository struct {
	db *sql.DB
}

func NewExercicioRepository(db *sql.DB) *ExercicioRepository {
	return &ExercicioRepository{db: db}
}

func (repository *ExercicioRepository) GetAllExercicios() ([]models.Exercicio, error) {
	rows, err := repository.db.Query("SELECT exercicio FROM aplic2008.libera_envio_pug_carga_mensal")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercicios []models.Exercicio

	for rows.Next() {
		var exercicio models.Exercicio
		if err := rows.Scan(&exercicio.Ano); err != nil {
			return nil, err
		}
		exercicios = append(exercicios, exercicio)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercicios, nil

}
