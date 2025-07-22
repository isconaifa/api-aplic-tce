package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type ConselheiroRepository struct {
	db *sql.DB
}

func NewConselheiroRepository(db *sql.DB) *ConselheiroRepository {
	return &ConselheiroRepository{db: db}
}

func (repository *ConselheiroRepository) GetAllConselheiros(exercicio int) ([]models.Conselheiro, error) {
	rows, err := repository.db.Query("select distinct d.cod_conselheiro as codigo,\n "+
		"c.apelido as nome \n "+
		"from publico.distribuicao_relator d,  \n "+
		"controlp.conselheiro  c  \n"+
		" where d.ano_relatoria = :exercicio\n "+
		"and d.cod_conselheiro = c.cod_conselheiro \n "+
		" order by c.apelido  ", exercicio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conselheiros []models.Conselheiro
	for rows.Next() {
		var conselheiro models.Conselheiro
		if err := rows.Scan(&conselheiro.Codigo, &conselheiro.Nome); err != nil {
			return nil, err
		}
		conselheiros = append(conselheiros, conselheiro)
	}
	return conselheiros, nil
}
