package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FiltroUnidadeOrcamentariaRepository struct {
	db *sql.DB
}

func NewFiltroUnidadeOrcamentariaRepository(db *sql.DB) *FiltroUnidadeOrcamentariaRepository {
	return &FiltroUnidadeOrcamentariaRepository{db: db}
}

func (repository *FiltroUnidadeOrcamentariaRepository) GetAllFiltroUnidadeOrcamentaria(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria string) ([]models.FiltroUnidadeOrcamentaria, error) {
	query := "SELECT DISTINCT\n    -- 1..3\n    e.emp_data AS empenho_data,\n    e.emp_numero AS empenho_numero,\n    c.cg_nome AS credor,\n\n    -- 4..8 (valores)\n    ( e.emp_valor - (\n        SELECT NVL(SUM(a.anul_valor),0)\n          FROM aplic2008.anulacao_empenho a\n         WHERE a.ent_codigo = e.ent_codigo\n           AND a.exercicio  = e.exercicio\n           AND a.org_codigo = e.org_codigo\n           AND a.unor_codigo= e.unor_codigo\n           AND a.emp_numero = e.emp_numero ) ) AS valor_empenhado,\n\n    ( (SELECT NVL(SUM(l.liq_valor),0)\n         FROM aplic2008.liquidacao_empenho l\n        WHERE l.ent_codigo = e.ent_codigo\n          AND l.exercicio  = e.exercicio\n          AND l.org_codigo = e.org_codigo\n          AND l.unor_codigo= e.unor_codigo\n          AND l.emp_numero = e.emp_numero)\n      - (SELECT NVL(SUM(al.anul_valor),0)\n           FROM aplic2008.anulacao_liquidacao_empenho al\n          WHERE al.ent_codigo = e.ent_codigo\n            AND al.exercicio  = e.exercicio\n            AND al.org_codigo = e.org_codigo\n            AND al.unor_codigo= e.unor_codigo\n            AND al.emp_numero = e.emp_numero) ) AS valor_liquidado,\n\n    ( (SELECT NVL(SUM(p.pgto_valor),0)\n         FROM aplic2008.pagamento_empenho p,\n              aplic2008.pagamento_empenho_liquidacao pl\n        WHERE p.ent_codigo = pl.ent_codigo\n          AND p.exercicio  = pl.exercicio\n          AND p.pgto_numero= pl.pgto_numero\n          AND pl.ent_codigo= e.ent_codigo\n          AND pl.exercicio = e.exercicio\n          AND pl.org_codigo= e.org_codigo\n          AND pl.unor_codigo= e.unor_codigo\n          AND pl.emp_numero= e.emp_numero)\n      - (SELECT NVL(SUM(alp.anul_valor),0)\n           FROM aplic2008.anulacao_pagamento_empenho alp,\n                aplic2008.pagamento_empenho_liquidacao pel\n          WHERE alp.ent_codigo = pel.ent_codigo\n            AND alp.exercicio  = pel.exercicio\n            AND alp.pgto_numero= pel.pgto_numero\n            AND pel.ent_codigo = e.ent_codigo\n            AND pel.org_codigo = e.org_codigo\n            AND pel.unor_codigo= e.unor_codigo\n            AND pel.emp_numero = e.emp_numero) ) AS valor_pago,\n\n    ( (SELECT NVL(SUM(p.pgto_valor),0)\n         FROM aplic2008.pagamento_empenho p,\n              aplic2008.pagamento_empenho_liquidacao pl\n        WHERE p.ent_codigo = pl.ent_codigo\n          AND p.exercicio  = pl.exercicio\n          AND p.pgto_numero= pl.pgto_numero\n          AND pl.ent_codigo= e.ent_codigo\n          AND pl.exercicio = e.exercicio\n          AND pl.org_codigo= e.org_codigo\n          AND pl.unor_codigo= e.unor_codigo\n          AND pl.emp_numero= e.emp_numero)\n      - (SELECT NVL(SUM(alp.anul_valor),0)\n           FROM aplic2008.anulacao_pagamento_empenho alp,\n                aplic2008.pagamento_empenho_liquidacao pel\n          WHERE alp.ent_codigo = pel.ent_codigo\n            AND alp.exercicio  = pel.exercicio\n            AND alp.pgto_numero= pel.pgto_numero\n            AND pel.ent_codigo = e.ent_codigo\n            AND pel.org_codigo = e.org_codigo\n            AND pel.unor_codigo= e.unor_codigo\n            AND pel.emp_numero = e.emp_numero)\n      + (SELECT COALESCE(SUM(d.dliq_valor),0)\n           FROM aplic2008.desconto_liquidado d\n          WHERE d.ent_codigo = e.ent_codigo\n            AND d.exercicio  = e.exercicio\n            AND d.org_codigo = e.org_codigo\n            AND d.unor_codigo= e.unor_codigo\n            AND d.emp_numero = e.emp_numero)\n      - (SELECT NVL(SUM(aedl.aedliq_valor),0)\n           FROM aplic2008.anulacao_estorno_desc_liquidad aedl\n          WHERE aedl.ent_codigo = e.ent_codigo\n            AND aedl.exercicio  = e.exercicio\n            AND aedl.org_codigo = e.org_codigo\n            AND aedl.unor_codigo= e.unor_codigo\n            AND aedl.emp_numero = e.emp_numero) ) AS valor_pago_mais_retencao,\n\n    ( SELECT NVL(SUM(a.anul_valor),0)\n        FROM aplic2008.anulacao_empenho a\n       WHERE a.ent_codigo = e.ent_codigo\n         AND a.exercicio  = e.exercicio\n         AND a.org_codigo = e.org_codigo\n         AND a.unor_codigo= e.unor_codigo\n         AND a.emp_numero = e.emp_numero ) AS anulado_empenho,\n\n    -- 9..11 (contagens como string)\n    TO_CHAR( (SELECT COUNT(1)\n                FROM aplic2008.nota_fiscal n\n               WHERE n.ent_codigo = e.ent_codigo\n                 AND n.exercicio  = e.exercicio\n                 AND n.org_codigo = e.org_codigo\n                 AND n.unor_codigo= e.unor_codigo\n                 AND n.emp_numero = e.emp_numero) ) AS qtde_notas_fiscais,\n\n    TO_CHAR( (SELECT COUNT(1)\n                FROM aplic2008.nota_fiscal n\n               WHERE n.ent_codigo = e.ent_codigo\n                 AND n.exercicio  = e.exercicio\n                 AND n.org_codigo = e.org_codigo\n                 AND n.unor_codigo= e.unor_codigo\n                 AND n.emp_numero = e.emp_numero\n                 AND n.ntfsc_numeronfe IS NOT NULL) ) AS qtde_nfe,\n\n    TO_CHAR( (SELECT COUNT(1)\n                FROM aplic2008.contrato_empenho n\n               WHERE n.ent_codigo = e.ent_codigo\n                 AND n.exercicio  = e.exercicio\n                 AND n.org_codigo = e.org_codigo\n                 AND n.unor_codigo= e.unor_codigo\n                 AND n.emp_numero = e.emp_numero) ) AS contratos,\n\n    -- 12..18\n    e.ent_codigo AS ent_codigo,\n\n    (SELECT COUNT(1)\n       FROM aplic2008.benef_assistencia_social_emp bas\n      WHERE bas.ent_codigo = e.ent_codigo\n        AND bas.exercicio  = e.exercicio\n        AND bas.org_codigo = e.org_codigo\n        AND bas.unor_codigo= e.unor_codigo\n        AND bas.emp_numero = e.emp_numero) AS qtde_beneficiarios,\n\n    DECODE(e.emp_instrumentocontrato,'1','Contrato','2','Carta-Contrato','3','Nota de Empenho da Despesa','4','Autorização de Compra','5','Ordem de Execução de Serviço','6','Outros Instrumentos Hábeis') AS instrumento_contrato,\n\n    frc.frec_descricao AS fonte_recurso_fonte,\n    e.conc_tipo  AS tipo_concurso,\n    e.conc_numero AS num_concurso,\n    DECODE(e.emp_benefassistenciasocial,'1','Não existem beneficiários','2','Existem beneficiários sem cadastro informatizado','3','Existem beneficiários com cadastro informatizado') AS assistencia_social,\n\n    -- 19..24\n    (SELECT COUNT(1)\n       FROM aplic2008.diaria dir\n      WHERE dir.ent_codigo = e.ent_codigo\n        AND dir.exercicio  = e.exercicio\n        AND dir.org_codigo = e.org_codigo\n        AND dir.unor_codigo= e.unor_codigo\n        AND dir.emp_numero = e.emp_numero) AS diarias,\n\n    ' ' AS relevante,\n    v.ent_nome AS unidade_gestora,\n    v.mun_nome AS municipio,\n    e.exercicio AS exercicio,\n    e.ent_codigo AS codigo_ug,\n\n    -- 25..33\n    prj.prat_descricao AS projeto_atividade,\n    e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS dotacao,\n    aplic2008.FN_EMPENHO_OBRA_PROJETO(e.ent_codigo, e.exercicio, e.org_codigo, e.unor_codigo, e.emp_numero) AS num_obra,\n    tse.tseng_descricao AS tipo_serv_engenharia,\n    UPPER(drg.drgrp_descricao) AS destinacao_recurso_grupo,\n    UPPER(dri.drids_descricao) AS destinacao_recurso_iduso,\n    e.drids_codigo AS destinacao_recurso_codiog_iduso,\n    e.destrec_codigo AS codigo_destinacao_recurso,\n    e.dresp_codigo AS codigo_destinacao_recurso_especificacao,\n    e.drgrp_codigo AS destinacao_recurso_codigo_grupo,\n    e.emp_fundamentocompradireta AS fund_compra_direta_codigo,\n    fc.fcd_descricao AS fund_compra_direta_descricao,\n    UPPER(dr.destrec_descricao) AS destinacao_recurso,\n\n    -- 38..46\n    o.org_nome AS orgao,\n    u.unor_nome AS unidade_orcamentaria,\n    e.unor_codigo AS unidade_orcamentaria_codigo,\n    e.fn_codigo AS funcao_codigo,\n    f.fn_descricao AS funcao_descricao,\n    e.sfn_codigo AS subfuncao_codigo,\n    sf.sfn_descricao AS subfuncao_descricao,\n    e.prg_codigo AS programa_codigo,\n    e.prat_numero AS num_projeto_atividade,\n\n    -- 47..55\n    e.ctec_codigo AS categoria_economica,\n    e.ndesp_codigo AS natureza_despesa,\n    e.mdap_codigo AS modalidade_aplicacao_codigo,\n    e.elde_codigo AS elemento_despesa_codigo,\n    el.elde_descricao AS elemento_despesa_descricao,\n    e.selde_codigo AS subelemento_despesa_codigo,\n    sub.selde_descricao AS subelemento_despesa,\n    e.emp_descricao AS descricao,\n    e.plic_numero AS num_processo_licitatorio,\n\n    -- 56..64\n    e.cont_tipo AS tipo_contrato,\n    e.cont_numaditivo AS num_aditivo_contrato,\n    e.conv_numero AS num_convenio,\n    e.conv_numaditivo AS num_aditivo_convenio,\n    DECODE(e.emp_compradiretaprocesso,'1','Não','2','Sim','N','Não','S','Sim') AS compra_direta,\n    DECODE(e.emp_tipo,'1','Estimativo','2','Global','3','Ordinário') AS tipo,\n    e.mesreferencia AS mes_referencia,\n    e.cg_identificacao AS identificacao_credor,\n    c.cg_tipopessoa AS tipo_pesssoa_codigo\n\nFROM aplic2008.empenho e\nLEFT JOIN aplic2008.cadastro_geral c\n       ON (e.ent_codigo = c.ent_codigo AND c.exercicio >= 2015 AND e.cg_identificacao = c.cg_identificacao)\nLEFT JOIN aplic2008.empenho_obra eo\n       ON (e.ent_codigo = eo.ent_codigo AND e.exercicio = eo.exercicio AND e.org_codigo = eo.org_codigo AND e.unor_codigo = eo.unor_codigo AND e.emp_numero = eo.emp_numero)\nLEFT JOIN aplic2008.orgao o\n       ON (e.ent_codigo = o.ent_codigo AND e.exercicio = o.exercicio AND e.org_codigo = o.org_codigo)\nLEFT JOIN aplic2008.unidade_orcamentaria u\n       ON (e.ent_codigo = u.ent_codigo AND e.exercicio = u.exercicio AND e.org_codigo = u.org_codigo AND e.unor_codigo = u.unor_codigo)\nLEFT JOIN aplic2008.projeto_atividade prj\n       ON (e.ent_codigo = prj.ent_codigo AND e.exercicio = prj.exercicio AND e.prat_numero = prj.prat_numero AND e.prg_codigo = prj.prg_codigo)\nLEFT JOIN aplic2008.fonte_recurso frc\n       ON (e.frec_codigo = frc.frec_codigo)\nLEFT JOIN aplic2008.destinacao_recurso_iduso dri\n       ON (e.exercicio = dri.exercicio AND e.drids_codigo = dri.drids_codigo)\nLEFT JOIN aplic2008.destinacao_recurso_grupo drg\n       ON (e.exercicio = drg.exercicio AND e.drgrp_codigo = drg.drgrp_codigo)\nLEFT JOIN aplic2008.destinacao_recurso_especific dre\n       ON (e.exercicio = dre.exercicio AND e.dresp_codigo = dre.dresp_codigo)\nLEFT JOIN aplic2008.destinacao_recurso dr\n       ON (e.exercicio = dr.exercicio AND e.destrec_codigo = dr.destrec_codigo)\nLEFT JOIN aplic2008.elemento_despesa el\n       ON (e.elde_codigo = el.elde_codigo AND e.exercicio = el.exercicio)\nLEFT JOIN aplic2008.subelemento_despesa sub\n       ON (e.exercicio = sub.selde_exercicio AND e.elde_codigo = sub.elde_codigo AND e.selde_codigo = sub.selde_codigo)\nLEFT JOIN aplic2008.funcao f\n       ON (e.fn_codigo = f.fn_codigo)\nLEFT JOIN aplic2008.subfuncao sf\n       ON (e.sfn_codigo = sf.sfn_codigo)\nINNER JOIN vw_entidade_aplic v\n       ON (e.ent_codigo = v.ent_codigo)\nLEFT JOIN aplic2008.tipo_despesa_rpps t\n       ON (e.emp_tipodespesarpps = t.trpps_codigo)\nLEFT JOIN aplic2008.contrato_empenho ce\n       ON (ce.ent_codigo = e.ent_codigo AND ce.exercicio = e.exercicio AND ce.org_codigo = e.org_codigo AND ce.unor_codigo = e.unor_codigo AND ce.emp_numero = e.emp_numero)\nLEFT JOIN aplic2008.tipo_servico_engenharia tse\n  " +
		"     ON (e.emp_tiposervicoengenharia = tse.tseng_codigo)\nLEFT JOIN aplic2008.modalidade_licitacao mlic\n       ON (e.mlic_codigo = mlic.mlic_codigo)\n" +
		"LEFT JOIN aplic2008.fundamento_compra_direta fc\n   " +
		"    ON (fc.fcd_codigo = e.emp_fundamentocompradireta)\n" +
		"WHERE 1 = 1\n " +
		" AND e.ent_codigo = :1\n " +
		" AND e.exercicio  = :2\n " +
		" AND e.org_codigo = :3\n " +
		" AND e.unor_codigo= :4\n" +
		"ORDER BY v.mun_nome, v.ent_nome, e.emp_numero, e.emp_data"

	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var filtroUnidadeOrcamentarias []models.FiltroUnidadeOrcamentaria
	for rows.Next() {
		var filtroUnidadeOrcamentaria models.FiltroUnidadeOrcamentaria
		if err := rows.Scan(
			&filtroUnidadeOrcamentaria.Emp_Data,
			&filtroUnidadeOrcamentaria.Emp_Numero,
			&filtroUnidadeOrcamentaria.Credor,
			&filtroUnidadeOrcamentaria.Valor_Empenhado,
			&filtroUnidadeOrcamentaria.Valor_Liquidado,
			&filtroUnidadeOrcamentaria.Valor_Pago,
			&filtroUnidadeOrcamentaria.Valor_PagoRetencao,
			&filtroUnidadeOrcamentaria.Anulado_Empenho,
			&filtroUnidadeOrcamentaria.Qtde_Notas_Fiscais,
			&filtroUnidadeOrcamentaria.Qtde_NFe,
			&filtroUnidadeOrcamentaria.Contratos,
			&filtroUnidadeOrcamentaria.Ent_codigo,
			&filtroUnidadeOrcamentaria.Qtde_beneficiarios,
			&filtroUnidadeOrcamentaria.Instrumento_contrato,
			&filtroUnidadeOrcamentaria.Fonte_recurso_fonte,
			&filtroUnidadeOrcamentaria.Tipo_concurso,
			&filtroUnidadeOrcamentaria.Num_Concurso,
			&filtroUnidadeOrcamentaria.Assistencia_social,
			&filtroUnidadeOrcamentaria.Diarias,
			&filtroUnidadeOrcamentaria.Relevante,
			&filtroUnidadeOrcamentaria.Unidade_gestora,
			&filtroUnidadeOrcamentaria.Municipio,
			&filtroUnidadeOrcamentaria.Exercicio,
			&filtroUnidadeOrcamentaria.Codigo_ug,
			&filtroUnidadeOrcamentaria.Projeto_atividade,
			&filtroUnidadeOrcamentaria.Dotacao,
			&filtroUnidadeOrcamentaria.Num_obra,
			&filtroUnidadeOrcamentaria.Tipo_serv_engenharia,
			&filtroUnidadeOrcamentaria.Destinacao_recurso_grupo,
			&filtroUnidadeOrcamentaria.Destinacao_recurso_iduso,
			&filtroUnidadeOrcamentaria.Destinacao_recurso_codiog_iduso,
			&filtroUnidadeOrcamentaria.Codigo_destinacao_recurso,
			&filtroUnidadeOrcamentaria.Codigo_destinacao_recurso_especificacao,
			&filtroUnidadeOrcamentaria.Destinacao_recurso_codigo_grupo,
			&filtroUnidadeOrcamentaria.Fund_compra_direta_codigo,
			&filtroUnidadeOrcamentaria.Fund_compra_direta_descricao,
			&filtroUnidadeOrcamentaria.Destinacao_recurso,
			&filtroUnidadeOrcamentaria.Orgao,
			&filtroUnidadeOrcamentaria.UnidadeOrcamentaria,
			&filtroUnidadeOrcamentaria.Unidade_orcamentaria_codigo,
			&filtroUnidadeOrcamentaria.Funcao_codigo,
			&filtroUnidadeOrcamentaria.Funcao_descricao,
			&filtroUnidadeOrcamentaria.Subfuncao_codigo,
			&filtroUnidadeOrcamentaria.Subfuncao_descricao,
			&filtroUnidadeOrcamentaria.Programa_codigo,
			&filtroUnidadeOrcamentaria.Num_Projeto_Atividade,
			&filtroUnidadeOrcamentaria.Categoria_Economica,
			&filtroUnidadeOrcamentaria.Natureza_Despesa,
			&filtroUnidadeOrcamentaria.Modalidade_aplicacao_codigo,
			&filtroUnidadeOrcamentaria.Elemento_despesa_codigo,
			&filtroUnidadeOrcamentaria.Elemento_despesa_descricao,
			&filtroUnidadeOrcamentaria.Subelemento_despesa_codigo,
			&filtroUnidadeOrcamentaria.Subelemento_despesa,
			&filtroUnidadeOrcamentaria.Descricao,
			&filtroUnidadeOrcamentaria.Num_processo_licitatorio,
			&filtroUnidadeOrcamentaria.Tipo_Contrato,
			&filtroUnidadeOrcamentaria.Num_aditivo_contrato,
			&filtroUnidadeOrcamentaria.Num_convenio,
			&filtroUnidadeOrcamentaria.Num_aditivo_convenio,
			&filtroUnidadeOrcamentaria.Compra_direta,
			&filtroUnidadeOrcamentaria.Tipo,
			&filtroUnidadeOrcamentaria.Mes_referencia,
			&filtroUnidadeOrcamentaria.Identificacao_credor,
			&filtroUnidadeOrcamentaria.Tipo_pesssoaCodigo,
		); err != nil {
			return nil, err
		}
		filtroUnidadeOrcamentarias = append(filtroUnidadeOrcamentarias, filtroUnidadeOrcamentaria)
	}
	return filtroUnidadeOrcamentarias, nil
}
