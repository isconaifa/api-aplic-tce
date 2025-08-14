package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FonteDestinacaoRecursoRepository struct {
	db *sql.DB
}

func NewFonteDestinacaoRecursoRepository(db *sql.DB) *FonteDestinacaoRecursoRepository {
	return &FonteDestinacaoRecursoRepository{db: db}
}

func (repository *FonteDestinacaoRecursoRepository) GetAllFontesDestinacaoRecurso(unidadeGestoraCodigo, ano string) ([]models.FonteDestinacaoRecurso, error) {
	query := "select distinct d.dresp_codigo as frec_codigo,\n" +
		"d.dresp_descricao as frec_descricao\n" +
		"from aplic2008.empenho e inner join\n" +
		"aplic2008.destinacao_recurso_especific d\n" +
		"on e.dresp_codigo = d.dresp_codigo\n" +
		"and e.exercicio = d.exercicio\nwhere 1 = 1\n" +
		"and e.Ent_Codigo = :1\n" +
		"and e.Exercicio = :2\n" +
		"order by d.dresp_descricao"

	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fontesDestinacaoRecurso []models.FonteDestinacaoRecurso
	for rows.Next() {
		var fonteDestinacaoRecurso models.FonteDestinacaoRecurso
		if err := rows.Scan(
			&fonteDestinacaoRecurso.Frec_codigo,
			&fonteDestinacaoRecurso.Frec_descricao,
		); err != nil {
			return nil, err
		}
		fontesDestinacaoRecurso = append(fontesDestinacaoRecurso, fonteDestinacaoRecurso)
	}
	return fontesDestinacaoRecurso, nil
}
