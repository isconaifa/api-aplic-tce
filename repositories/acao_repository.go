package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type AcaoRepository struct {
	db *sql.DB
}

func NewAcaoRepository(db *sql.DB) *AcaoRepository {
	return &AcaoRepository{db: db}
}

func (acaoRepository *AcaoRepository) GetAllAcoes(unidadeGestoraCodigo, ano string) ([]models.Acao, error) {
	query := "select distinct p.prat_numero,\n" +
		"upper(p.prat_descricao) as prat_descricao\n" +
		" from aplic2008.empenho e,\n" +
		"aplic2008.projeto_atividade p\n" +
		"where e.prat_numero = p.prat_numero\n" +
		"and e.prg_codigo = p.prg_codigo\n" +
		" and e.exercicio = p.exercicio\n" +
		" and e.ent_codigo = p.ent_codigo\n" +
		"  and e.Ent_Codigo = :1\n" +
		" and e.Exercicio = :2\n" +
		" order by prat_descricao"

	rows, err := acaoRepository.db.Query(query, unidadeGestoraCodigo, ano)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acoes []models.Acao
	for rows.Next() {
		var acao models.Acao
		if err := rows.Scan(&acao.Prat_Numero, &acao.Prat_Descricao); err != nil {
			return nil, err
		}
		acoes = append(acoes, acao)
	}
	return acoes, nil
}
