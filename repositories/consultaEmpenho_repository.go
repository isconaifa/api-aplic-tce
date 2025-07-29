package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type ConsultaEmpenhoRepository struct {
	db *sql.DB
}

func NewConsultaEmpenhoRepository(db *sql.DB) *ConsultaEmpenhoRepository {
	return &ConsultaEmpenhoRepository{db: db}
}

func (repository *ConsultaEmpenhoRepository) GetAllConsultaEmpenhos(unidadeGestoraCodigo string, ano string) ([]models.Empenho, error) {
	query := "SELECT e.emp_data AS empenho_data,\n" +
		"  e.emp_numero AS empenho_numero,\n " +
		" c.cg_nome AS credor,\n " +
		" e.emp_valor - NVL((SELECT SUM(a.anul_valor) FROM aplic2008.anulacao_empenho a\n    " +
		"                WHERE a.ent_codigo = e.ent_codigo\n         " +
		"             AND a.exercicio = e.exercicio\n            " +
		"          AND a.org_codigo = e.org_codigo\n           " +
		"           AND a.unor_codigo = e.unor_codigo\n            " +
		"          AND a.emp_numero = e.emp_numero), 0) AS valor_empenhado,\n " +
		" (SELECT NVL(SUM(l.liq_valor),0) FROM aplic2008.liquidacao_empenho l\n " +
		"   WHERE l.ent_codigo = e.ent_codigo\n  " +
		"    AND l.exercicio = e.exercicio\n  " +
		"    AND l.org_codigo = e.org_codigo\n   " +
		"  AND l.unor_codigo = e.unor_codigo\n   " +
		"   AND l.emp_numero = e.emp_numero) -\n " +
		" (SELECT NVL(SUM(al.anul_valor),0) FROM aplic2008.anulacao_liquidacao_empenho al\n  " +
		"  WHERE al.ent_codigo = e.ent_codigo\n  " +
		"    AND al.exercicio = e.exercicio\n " +
		"     AND al.org_codigo = e.org_codigo\n   " +
		"   AND al.unor_codigo = e.unor_codigo\n     " +
		" AND al.emp_numero = e.emp_numero) AS valor_liquidado,\n " +
		" (SELECT NVL(SUM(p.pgto_valor),0) FROM aplic2008.pagamento_empenho p\n  " +
		"  JOIN aplic2008.pagamento_empenho_liquidacao pl ON p.ent_codigo = pl.ent_codigo\n   " +
		"   AND p.exercicio = pl.exercicio\n  " +
		"    AND p.pgto_numero = pl.pgto_numero\n  " +
		"  WHERE pl.ent_codigo = e.ent_codigo\n " +
		"     AND pl.exercicio = e.exercicio\n   " +
		"   AND pl.org_codigo = e.org_codigo\n " +
		"     AND pl.unor_codigo = e.unor_codigo\n  " +
		"    AND pl.emp_numero = e.emp_numero) -\n " +
		" (SELECT NVL(SUM(alp.anul_valor),0) FROM aplic2008.anulacao_pagamento_empenho alp\n  " +
		"  JOIN aplic2008.pagamento_empenho_liquidacao pel ON alp.ent_codigo = pel.ent_codigo\n   " +
		"   AND alp.exercicio = pel.exercicio\n  " +
		"    AND alp.pgto_numero = pel.pgto_numero\n  " +
		"  WHERE pel.ent_codigo = e.ent_codigo\n   " +
		"   AND pel.org_codigo = e.org_codigo\n   " +
		"   AND pel.unor_codigo = e.unor_codigo\n   " +
		"   AND pel.emp_numero = e.emp_numero) AS valor_pago,\n " +
		" (SELECT NVL(SUM(p.pgto_valor),0) FROM aplic2008.pagamento_empenho p\n   " +
		" JOIN aplic2008.pagamento_empenho_liquidacao pl ON p.ent_codigo = pl.ent_codigo\n  " +
		"    AND p.exercicio = pl.exercicio\n   " +
		"   AND p.pgto_numero = pl.pgto_numero\n " +
		"   WHERE pl.ent_codigo = e.ent_codigo\n  " +
		"    AND pl.exercicio = e.exercicio\n  " +
		"    AND pl.org_codigo = e.org_codigo\n   " +
		"   AND pl.unor_codigo = e.unor_codigo\n   " +
		"   AND pl.emp_numero = e.emp_numero) +\n  " +
		"(SELECT COALESCE(SUM(DLIQ_VALOR),0) FROM aplic2008.desconto_liquidado d\n   " +
		" WHERE d.ent_codigo = e.ent_codigo\n   " +
		"   AND d.exercicio = e.exercicio\n  " +
		"    AND d.org_codigo = e.org_codigo\n  " +
		"    AND d.unor_codigo = e.unor_codigo\n  " +
		"    AND d.emp_numero = e.emp_numero) -\n " +
		" (SELECT NVL(SUM(aedl.aedliq_valor),0) FROM aplic2008.anulacao_estorno_desc_liquidad aedl\n  " +
		"  WHERE aedl.ent_codigo = e.ent_codigo\n   " +
		"   AND aedl.exercicio = e.exercicio\n  " +
		"    AND aedl.org_codigo = e.org_codigo\n    " +
		"  AND aedl.unor_codigo = e.unor_codigo\n   " +
		"   AND aedl.emp_numero = e.emp_numero) AS valor_pago_mais_retencao,\n " +
		" NVL((SELECT SUM(a.anul_valor) FROM aplic2008.anulacao_empenho a\n    " +
		"    WHERE a.ent_codigo = e.ent_codigo\n    " +
		"      AND a.exercicio = e.exercicio\n     " +
		"     AND a.org_codigo = e.org_codigo\n    " +
		"      AND a.unor_codigo = e.unor_codigo\n     " +
		"     AND a.emp_numero = e.emp_numero), 0) AS valor_anulado,\n " +
		" (SELECT COUNT(1) FROM aplic2008.nota_fiscal n\n  " +
		"  WHERE n.ent_codigo = e.ent_codigo\n    " +
		"  AND n.exercicio = e.exercicio\n   " +
		"   AND n.org_codigo = e.org_codigo\n   " +
		"   AND n.unor_codigo = e.unor_codigo\n  " +
		"    AND n.emp_numero = e.emp_numero) AS qtde_notas_fiscais,\n " +
		" (SELECT COUNT(1) FROM aplic2008.nota_fiscal n\n " +
		"  WHERE n.ent_codigo = e.ent_codigo\n   " +
		"   AND n.exercicio = e.exercicio\n  " +
		"    AND n.org_codigo = e.org_codigo\n  " +
		"    AND n.unor_codigo = e.unor_codigo\n   " +
		"   AND n.emp_numero = e.emp_numero\n    " +
		"  AND n.ntfsc_numeronfe IS NOT NULL) AS qtde_nfe,\n " +
		" (SELECT COUNT(1) FROM aplic2008.contrato_empenho ce\n  " +
		"  WHERE ce.ent_codigo = e.ent_codigo\n  " +
		"    AND ce.exercicio = e.exercicio\n   " +
		"   AND ce.org_codigo = e.org_codigo\n   " +
		"   AND ce.unor_codigo = e.unor_codigo\n   " +
		"   AND ce.emp_numero = e.emp_numero) AS contratos\n" +
		"FROM aplic2008.empenho e\n" +
		"LEFT JOIN aplic2008.cadastro_geral c ON (e.ent_codigo = c.ent_codigo AND c.exercicio >= 2015 AND e.cg_identificacao = c.cg_identificacao)\n" +
		"WHERE e.ent_codigo = :1\n " +
		" AND e.exercicio = :2\n" +
		"ORDER BY e.emp_data, e.emp_numero"

	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var empenhos []models.Empenho
	for rows.Next() {
		var empenho models.Empenho
		if err := rows.Scan(
			&empenho.Emp_Data,
			&empenho.Emp_Numero,
			&empenho.Credor,
			&empenho.Valor_Empenhado,
			&empenho.Valor_Liquidado,
			&empenho.Valor_Pago,
			&empenho.Valor_PagoRetencao,
			&empenho.Anulado_Empenho,
			&empenho.Qtde_Notas_Fiscais,
			&empenho.Qtde_NFe,
			&empenho.Contratos,
		); err != nil {
			return nil, err
		}
		empenhos = append(empenhos, empenho)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return empenhos, nil
}
