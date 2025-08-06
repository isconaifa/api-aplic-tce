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

func (repository *OrgaoRepository) GetAllOrgaos() ([]models.Orgao, error) {
	rows, err := repository.db.Query("select o.ent_codigo, o.exercicio, o.org_codigo, o.org_nome from aplic2008.orgao o")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orgaos []models.Orgao
	for rows.Next() {
		var orgao models.Orgao
		if err := rows.Scan(
			&orgao.Ent_codigo,
			&orgao.Exercicio,
			&orgao.Org_codigo,
			&orgao.Org_nome); err != nil {
			return nil, err
		}
		orgaos = append(orgaos, orgao)
	}
	return orgaos, nil
}
