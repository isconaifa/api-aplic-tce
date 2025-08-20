package repositories

import (
	"api-aplic-web/models"
	"database/sql"
	"fmt"
)

type FiltroSomenteAnuladosRepository struct {
	db *sql.DB
}

func NewFiltroSomenteAnuladosRepository(db *sql.DB) *FiltroSomenteAnuladosRepository {
	return &FiltroSomenteAnuladosRepository{db: db}
}

func (repository *FiltroSomenteAnuladosRepository) GetAllFiltroSomenteAnulados(unidadeGestoraCodigo, ano, dataInicioStr, dataFimStr string) ([]models.FiltroSomenteAnulados, error) {
	query := "SELECT DISTINCT\n    e.emp_data AS \"empenho_data\",\n   " +
		" e.emp_numero AS \"empenho_numero\",\n    c.cg_nome AS \"credor\",\n " +
		"   e.emp_valor - (\n        SELECT NVL(SUM(a.anul_valor), 0)\n    " +
		"    FROM aplic2008.anulacao_empenho a\n        WHERE a.ent_codigo = e.ent_codigo\n   " +
		"     AND a.exercicio = e.exercicio\n        AND a.org_codigo = e.org_codigo\n   " +
		"     AND a.unor_codigo = e.unor_codigo\n        AND a.emp_numero = e.emp_numero\n    " +
		"    AND a.anul_data BETWEEN TO_DATE(:startDate, 'DD/MM/YYYY') AND TO_DATE(:endDate, 'DD/MM/YYYY')\n    ) AS \"valor_empenhado\",\n " +
		"   (\n        SELECT NVL(SUM(l.liq_valor), 0)\n        FROM aplic2008.liquidacao_empenho l\n  " +
		"      WHERE l.ENT_CODIGO = E.ENT_CODIGO\n        AND l.EXERCICIO = E.EXERCICIO\n   " +
		"     AND l.ORG_CODIGO = E.ORG_CODIGO\n        AND l.UNOR_CODIGO = E.UNOR_CODIGO\n   " +
		"     AND l.EMP_NUMERO = E.EMP_NUMERO\n    ) - (\n        SELECT NVL(SUM(al.anul_valor), 0)\n   " +
		"     FROM aplic2008.anulacao_liquidacao_empenho al\n        WHERE al.ent_codigo = e.ent_codigo\n  " +
		"      AND al.exercicio = e.exercicio\n        AND al.org_codigo = e.org_codigo\n  " +
		"      AND al.unor_codigo = e.unor_codigo\n        AND al.emp_numero = e.emp_numero\n  " +
		"      AND al.anul_data BETWEEN TO_DATE(:startDate, 'DD/MM/YYYY') AND TO_DATE(:endDate, 'DD/MM/YYYY')\n    ) AS \"valor_liquidado\",\n " +
		"   (\n        SELECT NVL(SUM(p.pgto_valor), 0)\n        FROM aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl\n " +
		"       WHERE p.ENT_CODIGO = PL.ENT_CODIGO\n        AND p.EXERCICIO = PL.EXERCICIO\n        AND p.PGTO_NUMERO = PL.PGTO_NUMERO\n   " +
		"     AND PL.ENT_CODIGO = E.ENT_CODIGO\n        AND PL.EXERCICIO = E.EXERCICIO\n        AND PL.ORG_CODIGO = E.ORG_CODIGO\n   " +
		"     AND PL.UNOR_CODIGO = E.UNOR_CODIGO\n        AND PL.EMP_NUMERO = E.EMP_NUMERO\n    ) - (\n        SELECT NVL(SUM(alp.anul_valor), 0)\n   " +
		"     FROM aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel\n        WHERE alp.ent_codigo = pel.ent_codigo\n  " +
		"      AND alp.exercicio = pel.exercicio\n        AND alp.pgto_numero = pel.pgto_numero\n        AND pel.ent_codigo = e.ent_codigo\n " +
		"       AND pel.org_codigo = e.org_codigo\n        AND pel.unor_codigo = e.unor_codigo\n        AND pel.emp_numero = e.emp_numero\n  " +
		"      AND alp.anul_data BETWEEN TO_DATE(:startDate, 'DD/MM/YYYY') AND TO_DATE(:endDate, 'DD/MM/YYYY')\n    ) AS \"valor_pago\",\n " +
		"   (\n        SELECT NVL(SUM(p.pgto_valor), 0)\n        FROM aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl\n  " +
		"     WHERE p.ENT_CODIGO = PL.ENT_CODIGO\n        AND p.EXERCICIO = PL.EXERCICIO\n        AND p.PGTO_NUMERO = PL.PGTO_NUMERO\n    " +
		"    AND PL.ENT_CODIGO = E.ENT_CODIGO\n        AND PL.EXERCICIO = E.EXERCICIO\n        AND PL.ORG_CODIGO = E.ORG_CODIGO\n  " +
		"      AND PL.UNOR_CODIGO = E.UNOR_CODIGO\n        AND PL.EMP_NUMERO = E.EMP_NUMERO\n    ) - (\n        SELECT NVL(SUM(alp.anul_valor), 0)\n  " +
		"      FROM aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel\n        WHERE alp.ent_codigo = pel.ent_codigo\n  " +
		"      AND alp.exercicio = pel.exercicio\n        AND alp.pgto_numero = pel.pgto_numero\n        AND pel.ent_codigo = e.ent_codigo\n  " +
		"      AND pel.org_codigo = e.org_codigo\n        AND pel.unor_codigo = e.unor_codigo\n        AND pel.emp_numero = e.emp_numero\n    " +
		"    AND alp.anul_data BETWEEN TO_DATE(:startDate, 'DD/MM/YYYY') AND TO_DATE(:endDate, 'DD/MM/YYYY')\n    ) + (\n        SELECT COALESCE(SUM(DLIQ_VALOR), 0)\n  " +
		"      FROM aplic2008.desconto_liquidado d\n        WHERE d.ent_codigo = e.ent_codigo\n        AND d.exercicio = e.exercicio\n  " +
		"      AND d.org_codigo = e.org_codigo\n        AND d.unor_codigo = e.unor_codigo\n   " +
		"     AND d.emp_numero = e.emp_numero\n    ) - (\n        SELECT NVL(SUM(aedl.aedliq_valor), 0)\n   " +
		"     FROM aplic2008.anulacao_estorno_desc_liquidad aedl\n        WHERE aedl.ent_codigo = e.ent_codigo\n   " +
		"     AND aedl.exercicio = e.exercicio\n        AND aedl.org_codigo = e.org_codigo\n    " +
		"    AND aedl.unor_codigo = e.unor_codigo\n        AND aedl.emp_numero = e.emp_numero\n    ) AS \"valor_pago_mais_retencao\",\n " +
		"   (SELECT NVL(SUM(a.anul_valor), 0)\n        FROM aplic2008.anulacao_empenho a\n        WHERE a.ent_codigo = e.ent_codigo\n  " +
		"      AND a.exercicio = e.exercicio\n        AND a.org_codigo = e.org_codigo\n   " +
		"     AND a.unor_codigo = e.unor_codigo\n        AND a.emp_numero = e.emp_numero\n    " +
		"    AND a.anul_data BETWEEN TO_DATE(:startDate, 'DD/MM/YYYY') AND TO_DATE(:endDate, 'DD/MM/YYYY')\n    ) AS \"valor_anulado\",\n  " +
		"  (SELECT COUNT(1) FROM aplic2008.NOTA_FISCAL N WHERE N.ENT_CODIGO = E.ENT_CODIGO AND N.EXERCICIO = E.EXERCICIO AND N.ORG_CODIGO = E.ORG_CODIGO AND N.UNOR_CODIGO = E.UNOR_CODIGO AND N.EMP_NUMERO = E.EMP_NUMERO) AS \"qtde_notas_fiscais\",\n  " +
		"  (SELECT COUNT(1) FROM aplic2008.nota_fiscal n WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo AND n.unor_codigo = e.unor_codigo AND n.emp_numero = e.emp_numero AND n.ntfsc_numeronfe IS NOT NULL) AS \"qtde_nfe\",\n  " +
		"  (SELECT COUNT(1) FROM aplic2008.contrato_empenho n WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo AND n.unor_codigo = e.unor_codigo AND n.emp_numero = e.emp_numero) AS \"contratos\",\n " +
		"   e.ent_codigo AS \"ent_codigo\",\n    (SELECT COUNT(1) FROM aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas WHERE bas.ENT_CODIGO = E.ENT_CODIGO AND bas.EXERCICIO = E.EXERCICIO AND bas.ORG_CODIGO = E.ORG_CODIGO AND bas.UNOR_CODIGO = E.UNOR_CODIGO AND bas.EMP_NUMERO = E.EMP_NUMERO) AS \"qtde_beneficiarios\",\n " +
		"   DECODE(e.emp_instrumentocontrato, '1', 'Contrato', '2', 'Carta-Contrato', '3', 'Nota de Empenho da Despesa', '4', 'Autorização de Compra', '5', 'Ordem de Execução de Serviço', '6', 'Outros Instrumentos Hábeis') AS \"instrumento_contrato\",\n    FRC.FREC_DESCRICAO AS \"fonte_recurso_fonte\",\n  " +
		"  e.conc_tipo AS \"tipo_concurso\",\n    e.conc_numero AS \"num_concurso\",\n " +
		"   DECODE(e.emp_benefassistenciasocial, '1', 'Não existem beneficiários', '2', 'Existem beneficiários sem cadastro informatizado', '3', 'Existem beneficiários com cadastro informatizado') AS \"assistencia_social\",\n    (SELECT COUNT(1) FROM aplic2008.diaria dir WHERE dir.ent_codigo = e.ent_codigo AND dir.exercicio = e.exercicio AND dir.org_codigo = e.org_codigo AND dir.unor_codigo = e.unor_codigo AND dir.emp_numero = e.emp_numero) AS \"diarias\",\n" +
		"    ' ' AS \"relevante\",\n    v.ent_nome AS \"unidade_gestora\",\n  " +
		"  v.mun_nome AS \"municipio\",\n    e.exercicio AS \"exercicio\",\n " +
		"   e.ent_codigo AS \"codigo_ug\",\n    prj.prat_descricao AS \"projeto_atividade\",\n " +
		"   e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS \"dotacao\",\n " +
		"   aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) AS \"num_obra\",\n " +
		"   TSE.TSENG_DESCRICAO AS \"tipo_serv_engenharia\",\n    UPPER(drg.drgrp_descricao) AS \"destinacao_recurso_grupo\",\n " +
		"   UPPER(dri.drids_descricao) AS \"destinacao_recurso_iduso\",\n    e.drids_codigo AS \"destinacao_recurso_codiog_iduso\",\n " +
		"   e.destrec_codigo AS \"codigo_destinacao_recurso\",\n    e.dresp_codigo AS \"codigo_destinacao_recurso_especificacao\",\n   " +
		" e.drgrp_codigo AS \"destinacao_recurso_codigo_grupo\",\n    e.emp_fundamentocompradireta AS \"fund_compra_direta_codigo\",\n " +
		"   FC.FCD_DESCRICAO AS \"fund_compra_direta_descricao\",\n    UPPER(dr.destrec_descricao) AS \"destinacao_recurso\",\n   " +
		" o.org_nome AS \"orgao\",\n    u.unor_nome AS \"unidade_orcamentaria\",\n    e.unor_codigo AS \"unidade_orcamentaria_codigo\",\n " +
		"   e.fn_codigo AS \"funcao_codigo\",\n    f.fn_descricao AS \"funcao_descricao\",\n    e.sfn_codigo AS \"subfuncao_codigo\",\n " +
		"   sf.sfn_descricao AS \"subfuncao_descricao\",\n    e.prg_codigo AS \"programa_codigo\",\n " +
		"   e.prat_numero AS \"num_projeto_atividade\",\n    e.ctec_codigo AS \"categoria_economica\",\n " +
		"   e.ndesp_codigo AS \"natureza_despesa\",\n    e.mdap_codigo AS \"modalidade_aplicacao_codigo\",\n " +
		"   e.elde_codigo AS \"elemento_despesa_codigo\",\n    el.elde_descricao AS \"elemento_despesa_descricao\",\n  " +
		"  e.selde_codigo AS \"subelemento_despesa_codigo\",\n    sub.selde_descricao AS \"subelemento de despesa\",\n  " +
		"  e.emp_descricao AS \"descricao\",\n    e.plic_numero AS \"num_processo_licitatorio\",\n    e.cont_tipo AS \"tipo_contrato\",\n" +
		"    e.cont_numaditivo AS \"num_aditivo_contrato\",\n    e.conv_numero AS \"num_convenio\",\n  " +
		"  e.conv_numaditivo AS \"num_aditivo_convenio\",\n    DECODE(e.emp_compradiretaprocesso, '1', 'Não', '2', 'Sim', 'N', 'Não', 'S', 'Sim') AS \"compra_direta\",\n" +
		"    DECODE(e.emp_tipo, '1', 'Estimativo', '2', 'Global', '3', 'Ordinário') AS \"tipo\",\n " +
		"   e.mesreferencia AS \"mes_referencia\",\n    e.cg_identificacao AS \"identificacao_credor\",\n  " +
		"  c.cg_tipopessoa AS \"tipo_pesssoa_codigo\"\nFROM\n    aplic2008.empenho e\nINNER JOIN\n" +
		"    aplic2008.cadastro_geral c\n    ON (e.cg_identificacao = c.cg_identificacao) AND (c.exercicio >= 2015) AND (e.ent_codigo = c.ent_codigo)\n" +
		"LEFT JOIN\n    aplic2008.anulacao_empenho a\n    ON (e.emp_numero = a.emp_numero) AND (e.unor_codigo = a.unor_codigo) AND (e.org_codigo = a.org_codigo) AND (e.exercicio = a.exercicio) AND (e.ent_codigo = a.ent_codigo)\n" +
		"LEFT JOIN\n    aplic2008.empenho_obra eo\n    ON (e.emp_numero = eo.emp_numero) AND (e.unor_codigo = eo.unor_codigo) AND (e.org_codigo = eo.org_codigo) AND (e.exercicio = eo.exercicio) AND (e.ent_codigo = eo.ent_codigo)\nLEFT JOIN\n " +
		"   aplic2008.orgao o\n    ON (e.org_codigo = o.org_codigo) AND (e.exercicio = o.exercicio) AND (e.ent_codigo = o.ent_codigo)\nLEFT JOIN\n  " +
		"  aplic2008.unidade_orcamentaria u\n    ON (e.unor_codigo = u.unor_codigo) AND (e.org_codigo = u.org_codigo) AND (e.exercicio = u.exercicio) AND (e.ent_codigo = u.ent_codigo)\n" +
		"LEFT JOIN\n    aplic2008.projeto_atividade prj\n    ON (e.ent_codigo = prj.ent_codigo) AND (e.exercicio = prj.exercicio) AND (e.prat_numero = prj.prat_numero) AND (e.prg_codigo = prj.prg_codigo)\nLEFT JOIN\n " +
		"   aplic2008.FONTE_RECURSO FRC\n    ON (E.FREC_CODIGO = FRC.FREC_CODIGO)\nLEFT JOIN\n" +
		"    aplic2008.DESTINACAO_RECURSO_IDUSO DRI\n    ON (E.EXERCICIO = DRI.EXERCICIO AND E.DRIDS_CODIGO = DRI.DRIDS_CODIGO)\n" +
		"LEFT JOIN\n    aplic2008.DESTINACAO_RECURSO_GRUPO DRG\n    ON (E.EXERCICIO = DRG.EXERCICIO AND E.DRGRP_CODIGO = DRG.DRGRP_CODIGO)\n" +
		"LEFT JOIN\n    aplic2008.DESTINACAO_RECURSO_ESPECIFIC DRE\n    ON (E.EXERCICIO = DRE.EXERCICIO AND E.DRESP_CODIGO = DRE.DRESP_CODIGO)\n" +
		"LEFT JOIN\n    aplic2008.DESTINACAO_RECURSO DR\n    ON (E.EXERCICIO = DR.EXERCICIO AND E.DESTREC_CODIGO = DR.DESTREC_CODIGO)\nLEFT JOIN\n " +
		"   aplic2008.elemento_despesa el\n    ON (e.elde_codigo = el.elde_codigo) AND (e.exercicio = el.exercicio)\nLEFT JOIN\n " +
		"   aplic2008.subelemento_despesa sub\n    ON (e.exercicio = sub.selde_exercicio) AND (e.elde_codigo = sub.elde_codigo) AND (e.selde_codigo = sub.selde_codigo)\n" +
		"LEFT JOIN\n    aplic2008.funcao f\n    ON (e.fn_codigo = f.fn_codigo)\nLEFT JOIN\n " +
		"   aplic2008.subfuncao sf\n    ON (e.sfn_codigo = sf.sfn_codigo)\nINNER JOIN\n    vw_entidade_aplic v\n " +
		"   ON (e.ent_codigo = v.ent_codigo)\nLEFT JOIN\n    aplic2008.tipo_despesa_rpps t\n " +
		"   ON (e.emp_tipodespesarpps = t.trpps_codigo)\nLEFT JOIN\n    aplic2008.contrato_empenho ce\n " +
		"   ON (ce.ent_codigo = e.ent_codigo) AND (ce.exercicio = e.exercicio) AND (ce.org_codigo = e.org_codigo) AND (ce.unor_codigo = e.unor_codigo) AND (ce.emp_numero = e.emp_numero)\n" +
		"LEFT JOIN\n    aplic2008.TIPO_SERVICO_ENGENHARIA TSE\n    ON (E.EMP_TIPOSERVICOENGENHARIA = TSE.TSENG_CODIGO)\nLEFT JOIN\n    aplic2008.MODALIDADE_LICITACAO MLIC\n" +
		"    ON (E.MLIC_CODIGO = MLIC.MLIC_CODIGO)\nLEFT JOIN\n    aplic2008.FUNDAMENTO_COMPRA_DIRETA FC\n    ON FC.FCD_CODIGO = E.EMP_FUNDAMENTOCOMPRADIRETA\n" +
		"WHERE\n a.anul_data IS NOT NULL\n " +
		"   AND e.ent_codigo = :codigoUg\n " +
		" AND e.exercicio = :exercicio\n" +
		"    AND a.anul_data BETWEEN TO_DATE(:startDate, 'DD/MM/YYYY') AND TO_DATE(:endDate, 'DD/MM/YYYY')\n" +
		"ORDER BY\n " +
		"   v.mun_nome, v.ent_nome, e.emp_numero, e.emp_data"
	rows, err := repository.db.Query(query,
		sql.Named("codigoUg", unidadeGestoraCodigo),
		sql.Named("exercicio", ano),
		sql.Named("startDate", dataInicioStr),
		sql.Named("endDate", dataFimStr))
	if err != nil {
		return nil, err
		fmt.Printf("Error to execute query: %v\n", err)
	}
	defer rows.Close()
	var filtroSomenteAnulados []models.FiltroSomenteAnulados
	for rows.Next() {
		var f models.FiltroSomenteAnulados
		if err := rows.Scan(
			&f.Emp_Data,
			&f.Emp_Numero,
			&f.Credor,
			&f.Valor_Empenhado,
			&f.Valor_Liquidado,
			&f.Valor_Pago,
			&f.Valor_PagoRetencao,
			&f.Anulado_Empenho,
			&f.Qtde_Notas_Fiscais,
			&f.Qtde_NFe,
			&f.Contratos,
			&f.Ent_codigo,
			&f.Qtde_beneficiarios,
			&f.Instrumento_contrato,
			&f.Fonte_recurso_fonte,
			&f.Tipo_concurso,
			&f.Num_Concurso,
			&f.Assistencia_social,
			&f.Diarias,
			&f.Relevante,
			&f.Unidade_gestora,
			&f.Municipio,
			&f.Exercicio,
			&f.Codigo_ug,
			&f.Projeto_atividade,
			&f.Dotacao,
			&f.Num_obra,
			&f.Tipo_serv_engenharia,
			&f.Destinacao_recurso_grupo,
			&f.Destinacao_recurso_iduso,
			&f.Destinacao_recurso_codiog_iduso,
			&f.Codigo_destinacao_recurso,
			&f.Codigo_destinacao_recurso_especificacao,
			&f.Destinacao_recurso_codigo_grupo,
			&f.Fund_compra_direta_codigo,
			&f.Fund_compra_direta_descricao,
			&f.Destinacao_recurso,
			&f.Orgao,
			&f.UnidadeOrcamentaria,
			&f.Unidade_orcamentaria_codigo,
			&f.Funcao_codigo,
			&f.Funcao_descricao,
			&f.Subfuncao_codigo,
			&f.Subfuncao_descricao,
			&f.Programa_codigo,
			&f.Num_Projeto_Atividade,
			&f.Categoria_Economica,
			&f.Natureza_Despesa,
			&f.Modalidade_aplicacao_codigo,
			&f.Elemento_despesa_codigo,
			&f.Elemento_despesa_descricao,
			&f.Subelemento_despesa_codigo,
			&f.Subelemento_despesa,
			&f.Descricao,
			&f.Num_processo_licitatorio,
			&f.Tipo_Contrato,
			&f.Num_aditivo_contrato,
			&f.Num_convenio,
			&f.Num_aditivo_convenio,
			&f.Compra_direta,
			&f.Tipo,
			&f.Mes_referencia,
			&f.Identificacao_credor,
			&f.Tipo_pesssoaCodigo,
		); err != nil {
			return nil, err
		}
		filtroSomenteAnulados = append(filtroSomenteAnulados, f)
	}
	return filtroSomenteAnulados, nil
}
