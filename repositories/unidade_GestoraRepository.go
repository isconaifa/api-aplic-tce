package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type UnidadeGestoraRepository struct {
	db *sql.DB
}

func NewUnidadeGestoraRepository(db *sql.DB) *UnidadeGestoraRepository {
	return &UnidadeGestoraRepository{db: db}
}

func (repository *UnidadeGestoraRepository) GetUnidadesGestorasPorMunicipio(munCodigo string, ano string) ([]models.UnidadeGestora, error) {
	query := "SELECT DISTINCT\n" +
		"v.ent_codigo AS codigo,\n" +
		"v.ent_nome AS nome,\n" +
		"c.nome_conselheiro AS conselheiro\n" +
		"FROM vw_entidade_aplic v JOIN publico.distribuicao_relator d\n" +
		"ON d.cnpj_cpf_cod_tce_entidade = v.ent_codigo JOIN controlp.conselheiro c\n" +
		"ON d.cod_conselheiro = c.cod_conselheiro WHERE v.mun_codigo = :1\n" +
		"AND d.ano_relatoria = :2\n" +
		"ORDER BY v.ent_nome"

	rows, err := repository.db.Query(query, munCodigo, ano)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var unidadesGestoras []models.UnidadeGestora

	for rows.Next() {
		var unidadeGestora models.UnidadeGestora
		if err := rows.Scan(&unidadeGestora.Codigo, &unidadeGestora.Nome, &unidadeGestora.Conselheiro); err != nil {
			return nil, err
		}
		unidadesGestoras = append(unidadesGestoras, unidadeGestora)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return unidadesGestoras, nil
}
