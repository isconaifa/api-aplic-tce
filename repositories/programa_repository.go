package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type ProgramaRepository struct {
	db *sql.DB
}

func NewProgramaRepository(db *sql.DB) *ProgramaRepository {
	return &ProgramaRepository{db: db}
}

func (repository *ProgramaRepository) GetAllProgramas(unidadeGestora, exercicio string, codigoFuncao int) ([]models.Programa, error) {
	query := "select distinct p.prg_codigo, p.prg_descricao\n" +
		"from aplic2008.empenho e, aplic2008.Programa p\n" +
		"where e.prg_codigo = p.prg_codigo\n" +
		"and e.ent_codigo = p.ent_codigo \n" +
		"and e.exercicio = p.exercicio\n" +
		"and e.Ent_Codigo = :1\n" +
		"and e.Exercicio = :2\n" +
		"and e.fn_Codigo = :3\n" +
		"order by p.prg_descricao"

	rows, err := repository.db.Query(query, unidadeGestora, exercicio, codigoFuncao)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programas []models.Programa
	for rows.Next() {
		var programa models.Programa
		if err := rows.Scan(
			&programa.PRG_Codigo,
			&programa.PRG_Desc,
		); err != nil {
			return nil, err
		}
		programas = append(programas, programa)
	}
	return programas, nil
}
