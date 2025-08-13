package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type OrgaoRepository struct {
	db *sql.DB
}

func NewOrgaoRepository(db *sql.DB) *OrgaoRepository {
	return &OrgaoRepository{db: db}
}

func (repository *OrgaoRepository) GetAllOrgaos(unidadeGestora, exercicio string) ([]models.Orgao, error) {
	query := "select distinct o.org_codigo, o.org_nome\n" +
		"from aplic2008.empenho e, aplic2008.orgao o\n" +
		"where e.org_codigo = o.org_codigo\n" +
		"and e.ent_codigo = o.ent_codigo\n" +
		"and e.exercicio = o.exercicio\n" +
		"and e.Ent_Codigo = :1\n" +
		"and e.Exercicio = :2\n" +
		"order by o.org_nome"
	rows, err := repository.db.Query(query, unidadeGestora, exercicio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orgaos []models.Orgao
	for rows.Next() {
		var orgao models.Orgao
		if err := rows.Scan(
			&orgao.Org_codigo,
			&orgao.Org_nome,
		); err != nil {
			return nil, err
		}
		orgaos = append(orgaos, orgao)
	}
	return orgaos, nil
}
