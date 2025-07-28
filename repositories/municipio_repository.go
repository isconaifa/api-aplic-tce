package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type MunicipioRepository struct {
	db *sql.DB
}

func NewMunicipioRepository(db *sql.DB) *MunicipioRepository {
	return &MunicipioRepository{db: db}
}

func (repository *MunicipioRepository) GetAllMunicipios(exercicio string) ([]models.Municipio, error) {
	query := "select distinct v.mun_codigo as codigo,\n" +
		"v.mun_nome as nome\n" +
		" from publico.distribuicao_relator d,\n" +
		"controlp.conselheiro  c,\n" +
		"vw_entidade_aplic v\n " +
		"where d.ano_relatoria = :1\n" +
		"and d.cod_conselheiro = c.cod_conselheiro\n" +
		"and d.cnpj_cpf_cod_tce_entidade = v.ent_codigo\n " +
		"order by v.mun_nome"
	rows, err := repository.db.Query(query, exercicio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var municipios []models.Municipio

	for rows.Next() {
		var municipio models.Municipio
		if err := rows.Scan(&municipio.Codigo, &municipio.Nome); err != nil {
			return nil, err
		}
		municipios = append(municipios, municipio)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return municipios, nil
}
