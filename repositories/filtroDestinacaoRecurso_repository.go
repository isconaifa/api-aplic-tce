package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FiltroDestinacaoRecursoRepository struct {
	db *sql.DB
}

func NewFiltroDestinacaoRecursoRepository(db *sql.DB) *FiltroDestinacaoRecursoRepository {
	return &FiltroDestinacaoRecursoRepository{db: db}
}

func (repository *FiltroDestinacaoRecursoRepository) GetAllFiltroDestinacaoRecurso(unidadeGestoraCodigo, ano string, codigoDestinacaoRecurso int) ([]models.FiltroDestinacaoRecurso, error) {
	query := "SELECT DISTINCT\n  " +
		"  e.emp_data AS Emp_Data,\n  " +
		"  e.emp_numero AS Emp_Numero,\n " +
		"   c.cg_nome AS Credor,\n  " +
		"  (e.emp_valor - COALESCE( (SELECT SUM(a.anul_valor) FROM aplic2008.anulacao_empenho a WHERE a.ent_codigo = e.ent_codigo AND a.exercicio = e.exercicio AND a.org_codigo = e.org_codigo AND a.unor_codigo = e.unor_codigo AND a.emp_numero = e.emp_numero), 0 )) AS Valor_Empenhado,\n " +
		"   ((SELECT COALESCE(SUM(l.liq_valor),0) FROM aplic2008.liquidacao_empenho l WHERE l.ENT_CODIGO = E.ENT_CODIGO AND l.EXERCICIO = E.EXERCICIO AND l.ORG_CODIGO = E.ORG_CODIGO AND l.UNOR_CODIGO = E.UNOR_CODIGO AND l.EMP_NUMERO = E.EMP_NUMERO) - (SELECT COALESCE(SUM(al.anul_valor),0) FROM aplic2008.anulacao_liquidacao_empenho al WHERE al.ent_codigo = e.ent_codigo AND al.exercicio = e.exercicio AND al.org_codigo = e.org_codigo AND al.unor_codigo = e.unor_codigo AND al.emp_numero = e.emp_numero)) AS Valor_Liquidado,\n " +
		"   ((SELECT COALESCE(SUM(p.pgto_valor),0) FROM aplic2008.pagamento_empenho p JOIN aplic2008.pagamento_empenho_liquidacao pl ON p.ent_codigo = pl.ent_codigo AND p.exercicio = pl.exercicio AND p.pgto_numero = pl.pgto_numero WHERE pl.ent_codigo = e.ent_codigo AND pl.exercicio = e.exercicio AND pl.org_codigo = e.org_codigo AND pl.unor_codigo = e.unor_codigo AND pl.emp_numero = e.emp_numero) - (SELECT COALESCE(SUM(alp.anul_valor),0) FROM aplic2008.anulacao_pagamento_empenho alp JOIN aplic2008.pagamento_empenho_liquidacao pel ON\n" +
		" alp.ent_codigo = pel.ent_codigo AND alp.exercicio = pel.exercicio AND alp.pgto_numero = pel.pgto_numero WHERE pel.ent_codigo = e.ent_codigo AND pel.org_codigo = e.org_codigo AND pel.unor_codigo = e.unor_codigo AND pel.emp_numero = e.emp_numero)) AS Valor_Pago,\n    ((SELECT COALESCE(SUM(p.pgto_valor),0) FROM aplic2008.pagamento_empenho p JOIN aplic2008.pagamento_empenho_liquidacao pl ON p.ent_codigo = pl.ent_codigo AND p.exercicio = pl.exercicio AND p.pgto_numero = pl.pgto_numero WHERE pl.ent_codigo = e.ent_codigo AND pl.exercicio = e.exercicio AND pl.org_codigo = e.org_codigo AND pl.unor_codigo = e.unor_codigo AND pl.emp_numero = e.emp_numero) - (SELECT COALESCE(SUM(alp.anul_valor),0) FROM aplic2008.anulacao_pagamento_empenho alp JOIN aplic2008.pagamento_empenho_liquidacao pel ON alp.ent_codigo = pel.ent_codigo AND alp.exercicio = pel.exercicio AND alp.pgto_numero = pel.pgto_numero WHERE pel.ent_codigo = e.ent_codigo AND pel.org_codigo = e.org_codigo AND pel.unor_codigo = e.unor_codigo AND pel.emp_numero = e.emp_numero) + (SELECT COALESCE(SUM(d.DLIQ_VALOR),0) FROM aplic2008.desconto_liquidado d WHERE d.ent_codigo = e.ent_codigo AND d.exercicio = e.exercicio AND d.org_codigo = e.org_codigo AND d.unor_codigo = e.unor_codigo AND d.emp_numero = e.emp_numero) - (SELECT COALESCE(SUM(aedl.aedliq_valor),0) FROM aplic2008.anulacao_estorno_desc_liquidad aedl WHERE aedl.ent_codigo = e.ent_codigo AND aedl.exercicio = e.exercicio AND aedl.org_codigo = e.org_codigo AND aedl.unor_codigo = e.unor_codigo AND aedl.emp_numero = e.emp_numero)) AS Valor_PagoRetencao,\n    (SELECT COALESCE(SUM(a.anul_valor),0) FROM aplic2008.anulacao_empenho a WHERE a.ent_codigo = e.ent_codigo AND a.exercicio = e.exercicio AND a.org_codigo = e.org_codigo AND a.unor_codigo = e.unor_codigo AND a.emp_numero = e.emp_numero) AS Anulado_Empenho,\n    CAST((SELECT COUNT(1) FROM aplic2008.NOTA_FISCAL N WHERE N.ENT_CODIGO = E.ENT_CODIGO AND N.EXERCICIO = E.EXERCICIO AND N.ORG_CODIGO = E.ORG_CODIGO AND N.UNOR_CODIGO = E.UNOR_CODIGO AND N.EMP_NUMERO = E.EMP_NUMERO) AS VARCHAR(10)) AS Qtde_Notas_Fiscais,\n    CAST((SELECT COUNT(1) FROM aplic2008.nota_fiscal n WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo AND n.unor_codigo = e.unor_codigo AND n.emp_numero = e.emp_numero AND n.ntfsc_numeronfe IS NOT NULL) AS VARCHAR(10)) AS Qtde_NFe,\n    CAST((SELECT COUNT(1) FROM aplic2008.contrato_empenho n WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo AND n.unor_codigo = e.unor_codigo AND n.emp_numero = e.emp_numero) AS VARCHAR(10)) AS Contratos,\n    e.ent_codigo AS Ent_codigo,\n    (SELECT COUNT(1) FROM aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas WHERE bas.ENT_CODIGO = E.ENT_CODIGO AND bas.EXERCICIO = E.EXERCICIO AND bas.ORG_CODIGO = E.ORG_CODIGO AND bas.UNOR_CODIGO = E.UNOR_CODIGO AND bas.EMP_NUMERO = E.EMP_NUMERO) AS Qtde_beneficiarios,\n    DECODE(e.emp_instrumentocontrato, '1', 'Contrato', '2', 'Carta-Contrato', '3', 'Nota de Empenho da Despesa', '4', 'Autorização de Compra', '5', 'Ordem de Execução de Serviço', '6', 'Outros Instrumentos Hábeis') AS Instrumento_contrato,\n    frc.frec_descricao AS Fonte_recurso_fonte,\n    e.conc_tipo AS Tipo_concurso,\n    e.conc_numero AS Num_Concurso,\n    DECODE(e.emp_benefassistenciasocial, '1', 'Não existem beneficiários', '2', 'Existem beneficiários sem cadastro informatizado', '3', 'Existem beneficiários com cadastro informatizado') AS Assistencia_social,\n  " +
		"  (SELECT COUNT(1) FROM aplic2008.diaria dir WHERE dir.ent_codigo = e.ent_codigo AND dir.exercicio = e.exercicio AND dir.org_codigo = e.org_codigo AND dir.unor_codigo = e.unor_codigo AND dir.emp_numero = e.emp_numero) AS Diarias,\n    ' ' AS Relevante,\n " +
		"   v.ent_nome AS Unidade_gestora,\n  " +
		"  v.mun_nome AS Municipio,\n    e.exercicio AS Exercicio,\n    e.ent_codigo AS Codigo_ug,\n " +
		"   prj.prat_descricao AS Projeto_atividade,\n    e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS Dotacao,\n    aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) AS Num_obra,\n    TSE.TSENG_DESCRICAO AS Tipo_serv_engenharia,\n    DRG.DRGRP_DESCRICAO AS Destinacao_recurso_grupo,\n    DRI.DRIDS_DESCRICAO AS Destinacao_recurso_iduso,\n    e.drids_codigo AS Destinacao_recurso_codiog_iduso,\n    dr.destrec_codigo AS Codigo_destinacao_recurso,\n    dre.dresp_codigo AS Codigo_destinacao_recurso_especificacao,\n    e.drgrp_codigo AS Destinacao_recurso_codigo_grupo,\n    FC.FCD_CODIGO AS Fund_compra_direta_codigo,\n    FC.FCD_DESCRICAO AS Fund_compra_direta_descricao,\n    dr.destrec_descricao AS Destinacao_recurso,\n    o.org_nome AS Orgao,\n    u.unor_nome AS UnidadeOrcamentaria,\n    e.unor_codigo AS Unidade_orcamentaria_codigo,\n    e.fn_codigo AS Funcao_codigo,\n    f.fn_descricao AS Funcao_descricao,\n    e.sfn_codigo AS Subfuncao_codigo,\n    sf.sfn_descricao AS Subfuncao_descricao,\n    e.prg_codigo AS Programa_codigo,\n    e.prat_numero AS Num_Projeto_Atividade,\n    e.ctec_codigo AS Categoria_Economica,\n    e.ndesp_codigo AS Natureza_Despesa,\n    e.mdap_codigo AS Modalidade_aplicacao_codigo,\n    e.elde_codigo AS Elemento_despesa_codigo,\n    el.elde_descricao AS Elemento_despesa_descricao,\n    e.selde_codigo AS Subelemento_despesa_codigo,\n    sub.selde_descricao AS Subelemento_despesa,\n    e.emp_descricao AS Descricao,\n    e.plic_numero AS Num_processo_licitatorio,\n    e.cont_tipo AS Tipo_Contrato,\n    e.cont_numaditivo AS Num_aditivo_contrato,\n    e.conv_numero AS Num_convenio,\n    e.conv_numaditivo AS Num_aditivo_convenio,\n    DECODE(e.emp_compradiretaprocesso, '1', 'Não', '2', 'Sim', 'N', 'Não', 'S', 'Sim') AS Compra_direta,\n    DECODE(e.emp_tipo, '1', 'Estimativo', '2', 'Global', '3', 'Ordinário') AS Tipo,\n    e.mesreferencia AS Mes_referencia,\n    e.cg_identificacao AS Identificacao_credor,\n    c.cg_tipopessoa AS Tipo_pesssoaCodigo\nFROM\n    aplic2008.empenho e\nLEFT JOIN\n    aplic2008.cadastro_geral c ON e.ent_codigo = c.ent_codigo AND c.exercicio >= 2015 AND e.cg_identificacao = c.cg_identificacao\nLEFT JOIN\n    aplic2008.orgao o ON e.ent_codigo = o.ent_codigo AND e.exercicio = o.exercicio AND e.org_codigo = o.org_codigo\nLEFT JOIN\n    aplic2008.unidade_orcamentaria u ON e.ent_codigo = u.ent_codigo AND e.exercicio = u.exercicio AND e.org_codigo = u.org_codigo AND e.unor_codigo = u.unor_codigo\nLEFT JOIN\n    aplic2008.projeto_atividade prj ON e.ent_codigo = prj.ent_codigo AND e.exercicio = prj.exercicio AND e.prat_numero = prj.prat_numero AND e.prg_codigo = prj.prg_codigo\nLEFT JOIN\n    aplic2008.FONTE_RECURSO FRC ON e.frec_codigo = FRC.frec_codigo\nLEFT JOIN\n    aplic2008.DESTINACAO_RECURSO_IDUSO DRI ON e.exercicio = DRI.exercicio AND e.drids_codigo = DRI.drids_codigo\nLEFT JOIN\n    aplic2008.DESTINACAO_RECURSO_GRUPO DRG ON e.exercicio = DRG.exercicio AND e.drgrp_codigo = DRG.drgrp_codigo\nLEFT JOIN\n    aplic2008.DESTINACAO_RECURSO_ESPECIFIC DRE ON e.exercicio = DRE.exercicio AND e.dresp_codigo = DRE.dresp_codigo\nLEFT JOIN\n    aplic2008.DESTINACAO_RECURSO DR ON e.exercicio = DR.exercicio AND e.destrec_codigo = DR.destrec_codigo\nLEFT JOIN\n    aplic2008.elemento_despesa el ON e.elde_codigo = el.elde_codigo AND e.exercicio = el.exercicio\nLEFT JOIN\n    aplic2008.subelemento_despesa sub ON e.exercicio = sub.selde_exercicio AND e.elde_codigo = sub.elde_codigo AND e.selde_codigo = sub.selde_codigo\nLEFT JOIN\n    aplic2008.funcao f ON e.fn_codigo = f.fn_codigo\nLEFT JOIN\n    aplic2008.subfuncao sf ON e.sfn_codigo = sf.sfn_codigo\nINNER JOIN\n    vw_entidade_aplic v ON e.ent_codigo = v.ent_codigo\nLEFT JOIN\n    aplic2008.TIPO_SERVICO_ENGENHARIA TSE ON e.emp_tiposervicoengenharia = TSE.tseng_codigo\nLEFT JOIN\n    aplic2008.FUNDAMENTO_COMPRA_DIRETA FC ON FC.fcd_codigo = e.emp_fundamentocompradireta\nWHERE\n    1 = 1\n " +
		"   AND e.ent_codigo = :1\n  " +
		"  AND e.exercicio = :2\n  " +
		"  AND dre.dresp_codigo = :3\n" +
		"ORDER BY\n " +
		"v.mun_nome,\n    v.ent_nome,\n    e.emp_numero,\n    e.emp_data"
	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano, codigoDestinacaoRecurso)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var filtroDestinacaoRecursos []models.FiltroDestinacaoRecurso
	for rows.Next() {
		var filtroDestinacaoRecurso models.FiltroDestinacaoRecurso
		if err := rows.Scan(
			&filtroDestinacaoRecurso.Emp_Data,
			&filtroDestinacaoRecurso.Emp_Numero,
			&filtroDestinacaoRecurso.Credor,
			&filtroDestinacaoRecurso.Valor_Empenhado,
			&filtroDestinacaoRecurso.Valor_Liquidado,
			&filtroDestinacaoRecurso.Valor_Pago,
			&filtroDestinacaoRecurso.Valor_PagoRetencao,
			&filtroDestinacaoRecurso.Anulado_Empenho,
			&filtroDestinacaoRecurso.Qtde_Notas_Fiscais,
			&filtroDestinacaoRecurso.Qtde_NFe,
			&filtroDestinacaoRecurso.Contratos,
			&filtroDestinacaoRecurso.Ent_codigo,
			&filtroDestinacaoRecurso.Qtde_beneficiarios,
			&filtroDestinacaoRecurso.Instrumento_contrato,
			&filtroDestinacaoRecurso.Fonte_recurso_fonte,
			&filtroDestinacaoRecurso.Tipo_concurso,
			&filtroDestinacaoRecurso.Num_Concurso,
			&filtroDestinacaoRecurso.Assistencia_social,
			&filtroDestinacaoRecurso.Diarias,
			&filtroDestinacaoRecurso.Relevante,
			&filtroDestinacaoRecurso.Unidade_gestora,
			&filtroDestinacaoRecurso.Municipio,
			&filtroDestinacaoRecurso.Exercicio,
			&filtroDestinacaoRecurso.Codigo_ug,
			&filtroDestinacaoRecurso.Projeto_atividade,
			&filtroDestinacaoRecurso.Dotacao,
			&filtroDestinacaoRecurso.Num_obra,
			&filtroDestinacaoRecurso.Tipo_serv_engenharia,
			&filtroDestinacaoRecurso.Destinacao_recurso_grupo,
			&filtroDestinacaoRecurso.Destinacao_recurso_iduso,
			&filtroDestinacaoRecurso.Destinacao_recurso_codiog_iduso,
			&filtroDestinacaoRecurso.Codigo_destinacao_recurso,
			&filtroDestinacaoRecurso.Codigo_destinacao_recurso_especificacao,
			&filtroDestinacaoRecurso.Destinacao_recurso_codigo_grupo,
			&filtroDestinacaoRecurso.Fund_compra_direta_codigo,
			&filtroDestinacaoRecurso.Fund_compra_direta_descricao,
			&filtroDestinacaoRecurso.Destinacao_recurso,
			&filtroDestinacaoRecurso.Orgao,
			&filtroDestinacaoRecurso.UnidadeOrcamentaria,
			&filtroDestinacaoRecurso.Unidade_orcamentaria_codigo,
			&filtroDestinacaoRecurso.Funcao_codigo,
			&filtroDestinacaoRecurso.Funcao_descricao,
			&filtroDestinacaoRecurso.Subfuncao_codigo,
			&filtroDestinacaoRecurso.Subfuncao_descricao,
			&filtroDestinacaoRecurso.Programa_codigo,
			&filtroDestinacaoRecurso.Num_Projeto_Atividade,
			&filtroDestinacaoRecurso.Categoria_Economica,
			&filtroDestinacaoRecurso.Natureza_Despesa,
			&filtroDestinacaoRecurso.Modalidade_aplicacao_codigo,
			&filtroDestinacaoRecurso.Elemento_despesa_codigo,
			&filtroDestinacaoRecurso.Elemento_despesa_descricao,
			&filtroDestinacaoRecurso.Subelemento_despesa_codigo,
			&filtroDestinacaoRecurso.Subelemento_despesa,
			&filtroDestinacaoRecurso.Descricao,
			&filtroDestinacaoRecurso.Num_processo_licitatorio,
			&filtroDestinacaoRecurso.Tipo_Contrato,
			&filtroDestinacaoRecurso.Num_aditivo_contrato,
			&filtroDestinacaoRecurso.Num_convenio,
			&filtroDestinacaoRecurso.Num_aditivo_convenio,
			&filtroDestinacaoRecurso.Compra_direta,
			&filtroDestinacaoRecurso.Tipo,
			&filtroDestinacaoRecurso.Mes_referencia,
			&filtroDestinacaoRecurso.Identificacao_credor,
			&filtroDestinacaoRecurso.Tipo_pesssoaCodigo,
		); err != nil {
			return nil, err
		}
		filtroDestinacaoRecursos = append(filtroDestinacaoRecursos, filtroDestinacaoRecurso)
	}
	return filtroDestinacaoRecursos, nil
}
