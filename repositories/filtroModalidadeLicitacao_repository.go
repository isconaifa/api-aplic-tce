package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FiltroModalidadeLicitacaoRepository struct {
	db *sql.DB
}

func NewFiltroModalidadeLicitacaoRepository(db *sql.DB) *FiltroModalidadeLicitacaoRepository {
	return &FiltroModalidadeLicitacaoRepository{db: db}
}

func (repository *FiltroModalidadeLicitacaoRepository) GetAllFilterModalidadeLicitacao(unidadeGestoraCodigo, ano string, codigoModalidadeLicitacao int) ([]models.FiltroModalidadeLicitacao, error) {
	query := "SELECT DISTINCT\n    e.emp_data AS empenho_data,\n    e.emp_numero AS empenho_numero,\n    c.cg_nome AS credor,\n    e.emp_valor - (\n        SELECT NVL(SUM(a.anul_valor), 0)\n        FROM aplic2008.anulacao_empenho a\n        WHERE a.ent_codigo = e.ent_codigo\n        AND a.exercicio = e.exercicio\n        AND a.org_codigo = e.org_codigo\n        AND a.unor_codigo = e.unor_codigo\n        AND a.emp_numero = e.emp_numero\n    ) AS valor_empenhado,\n    (\n        SELECT NVL(SUM(l.liq_valor), 0)\n        FROM aplic2008.liquidacao_empenho l\n        WHERE l.ENT_CODIGO = E.ENT_CODIGO\n        AND l.EXERCICIO = E.EXERCICIO\n        AND l.ORG_CODIGO = E.ORG_CODIGO\n        AND l.UNOR_CODIGO = E.UNOR_CODIGO\n        AND l.EMP_NUMERO = E.EMP_NUMERO\n    ) - (\n        SELECT NVL(SUM(al.anul_valor), 0)\n        FROM aplic2008.anulacao_liquidacao_empenho al\n        WHERE al.ent_codigo = e.ent_codigo\n        AND al.exercicio = e.exercicio\n        AND al.org_codigo = e.org_codigo\n        AND al.unor_codigo = e.unor_codigo\n        AND al.emp_numero = e.emp_numero\n    ) AS valor_liquidado,\n    (\n        SELECT NVL(SUM(p.pgto_valor), 0)\n        FROM aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl\n        WHERE p.ENT_CODIGO = PL.ENT_CODIGO\n        AND p.EXERCICIO = PL.EXERCICIO\n        AND p.PGTO_NUMERO = PL.PGTO_NUMERO\n        AND PL.ENT_CODIGO = E.ENT_CODIGO\n        AND PL.EXERCICIO = E.EXERCICIO\n        AND PL.ORG_CODIGO = E.ORG_CODIGO\n        AND PL.UNOR_CODIGO = E.UNOR_CODIGO\n        AND PL.EMP_NUMERO = E.EMP_NUMERO\n    ) - (\n        SELECT NVL(SUM(alp.anul_valor), 0)\n        FROM aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel\n        WHERE alp.ent_codigo = pel.ent_codigo\n        AND alp.exercicio = pel.exercicio\n        AND alp.pgto_numero = pel.pgto_numero\n        AND pel.ent_codigo = e.ent_codigo\n        AND pel.org_codigo = e.org_codigo\n        AND pel.unor_codigo = e.unor_codigo\n        AND pel.emp_numero = e.emp_numero\n    ) AS valor_pago,\n    (\n        SELECT NVL(SUM(p.pgto_valor), 0)\n        FROM aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl\n        WHERE p.ENT_CODIGO = PL.ENT_CODIGO\n        AND p.EXERCICIO = PL.EXERCICIO\n        AND p.PGTO_NUMERO = PL.PGTO_NUMERO\n        AND PL.ENT_CODIGO = E.ENT_CODIGO\n        AND PL.EXERCICIO = E.EXERCICIO\n        AND PL.ORG_CODIGO = E.ORG_CODIGO\n        AND PL.UNOR_CODIGO = E.UNOR_CODIGO\n        AND PL.EMP_NUMERO = E.EMP_NUMERO\n    ) - (\n        SELECT NVL(SUM(alp.anul_valor), 0)\n        FROM aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel\n        WHERE alp.ent_codigo = pel.ent_codigo\n        AND alp.exercicio = pel.exercicio\n        AND alp.pgto_numero = pel.pgto_numero\n        AND pel.ent_codigo = e.ent_codigo\n        AND pel.org_codigo = e.org_codigo\n        AND pel.unor_codigo = e.unor_codigo\n        AND pel.emp_numero = e.emp_numero\n    ) + (\n        SELECT COALESCE(SUM(DLIQ_VALOR), 0)\n        FROM aplic2008.desconto_liquidado d\n        WHERE d.ent_codigo = e.ent_codigo\n        AND d.exercicio = e.exercicio\n        AND d.org_codigo = e.org_codigo\n        AND d.unor_codigo = e.unor_codigo\n        AND d.emp_numero = e.emp_numero\n    ) - (\n        SELECT NVL(SUM(aedl.aedliq_valor), 0)\n        FROM aplic2008.anulacao_estorno_desc_liquidad aedl\n        WHERE aedl.ent_codigo = e.ent_codigo\n        AND aedl.exercicio = e.exercicio\n        AND aedl.org_codigo = e.org_codigo\n        AND aedl.unor_codigo = e.unor_codigo\n        AND aedl.emp_numero = e.emp_numero\n    ) AS valor_pago_mais_retencao,\n    (\n        SELECT NVL(SUM(a.anul_valor), 0)\n        FROM aplic2008.anulacao_empenho a\n        WHERE a.ent_codigo = e.ent_codigo\n        AND a.exercicio = e.exercicio\n        AND a.org_codigo = e.org_codigo\n        AND a.unor_codigo = e.unor_codigo\n        AND a.emp_numero = e.emp_numero\n    ) AS valor_anulado,\n    (\n        SELECT COUNT(1)\n        FROM aplic2008.NOTA_FISCAL N\n        WHERE N.ENT_CODIGO = E.ENT_CODIGO\n        AND N.EXERCICIO = E.EXERCICIO\n        AND N.ORG_CODIGO = E.ORG_CODIGO\n        AND N.UNOR_CODIGO = E.UNOR_CODIGO\n        AND N.EMP_NUMERO = E.EMP_NUMERO\n    ) AS qtde_notas_fiscais,\n    (\n        SELECT COUNT(1)\n        FROM aplic2008.nota_fiscal n\n        WHERE n.ent_codigo = e.ent_codigo\n        AND n.exercicio = e.exercicio\n        AND n.org_codigo = e.org_codigo\n        AND n.unor_codigo = e.unor_codigo\n        AND n.emp_numero = e.emp_numero\n        AND n.ntfsc_numeronfe IS NOT NULL\n    ) AS qtde_nfe,\n    (\n        SELECT COUNT(1)\n        FROM aplic2008.contrato_empenho n\n        WHERE n.ent_codigo = e.ent_codigo\n        AND n.exercicio = e.exercicio\n        AND n.org_codigo = e.org_codigo\n        AND n.unor_codigo = e.unor_codigo\n        AND n.emp_numero = e.emp_numero\n    ) AS contratos,\n    e.ent_codigo,\n    (\n        SELECT COUNT(1)\n        FROM aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas\n        WHERE bas.ENT_CODIGO = E.ENT_CODIGO\n        AND bas.EXERCICIO = E.EXERCICIO\n        AND bas.ORG_CODIGO = E.ORG_CODIGO\n        AND bas.UNOR_CODIGO = E.UNOR_CODIGO\n        AND bas.EMP_NUMERO = E.EMP_NUMERO\n    ) AS qtde_beneficiarios,\n    DECODE(e.emp_instrumentocontrato, '1', 'Contrato', '2', 'Carta-Contrato', '3', 'Nota de Empenho da Despesa', '4', 'Autorização de Compra', '5', 'Ordem de Execução de Serviço', '6', 'Outros Instrumentos Hábeis') AS instrumento_contrato,\n    FRC.FREC_DESCRICAO AS fonte_recurso_fonte,\n    e.conc_tipo AS tipo_concurso,\n    e.conc_numero AS num_concurso,\n    DECODE(e.emp_benefassistenciasocial, '1', 'Não existem beneficiários', '2', 'Existem beneficiários sem cadastro informatizado', '3', 'Existem beneficiários com cadastro informatizado') AS assistencia_social,\n    (\n        SELECT COUNT(1)\n        FROM aplic2008.diaria dir\n        WHERE dir.ent_codigo = e.ent_codigo\n        AND dir.exercicio = e.exercicio\n        AND dir.org_codigo = e.org_codigo\n        AND dir.unor_codigo = e.unor_codigo\n        AND dir.emp_numero = e.emp_numero\n    ) AS diarias,\n    ' ' AS relevante,\n    v.ent_nome AS unidade_gestora,\n    v.mun_nome AS municipio,\n    e.exercicio AS exercicio,\n    e.ent_codigo AS codigo_ug,\n    prj.prat_descricao AS projeto_atividade,\n    e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS dotacao,\n    aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) AS num_obra,\n    TSE.TSENG_DESCRICAO AS tipo_serv_engenharia,\n    upper(drg.drgrp_descricao) AS destinacao_recurso_grupo,\n    upper(dri.drids_descricao) AS destinacao_recurso_iduso,\n    e.drids_codigo AS destinacao_recurso_codiog_iduso,\n    e.destrec_codigo AS codigo_destinacao_recurso,\n    e.dresp_codigo AS codigo_destinacao_recurso_especificacao,\n    e.drgrp_codigo AS destinacao_recurso_codigo_grupo,\n    e.emp_fundamentocompradireta AS fund_compra_direta_codigo,\n    FC.FCD_DESCRICAO AS fund_compra_direta_descricao,\n    upper(dr.destrec_descricao) AS destinacao_recurso,\n    o.org_nome AS orgao,\n    u.unor_nome AS unidade_orcamentaria,\n    e.unor_codigo AS unidade_orcamentaria_codigo,\n    e.fn_codigo AS funcao_codigo,\n    f.fn_descricao AS funcao_descricao,\n    e.sfn_codigo AS subfuncao_codigo,\n    sf.sfn_descricao AS subfuncao_descricao,\n    e.prg_codigo AS programa_codigo,\n    e.prat_numero AS num_projeto_atividade,\n    e.ctec_codigo AS categoria_economica,\n    e.ndesp_codigo AS natureza_despesa,\n    e.mdap_codigo AS modalidade_aplicacao_codigo,\n    e.elde_codigo AS elemento_despesa_codigo,\n    el.elde_descricao AS elemento_despesa_descricao,\n    e.selde_codigo AS subelemento_despesa_codigo,\n    sub.selde_descricao AS subelemento_despesa,\n    e.emp_descricao AS descricao,\n    e.plic_numero AS num_processo_licitatorio,\n    e.cont_tipo AS tipo_contrato,\n    e.cont_numaditivo AS num_aditivo_contrato,\n    e.conv_numero AS num_convenio,\n    e.conv_numaditivo AS num_aditivo_convenio,\n    DECODE(e.emp_compradiretaprocesso, '1', 'Não', '2', 'Sim', 'N', 'Não', 'S', 'Sim') AS compra_direta,\n    DECODE(e.emp_tipo, '1', 'Estimativo', '2', 'Global', '3', 'Ordinário') AS tipo,\n    e.mesreferencia AS mes_referencia,\n    e.cg_identificacao AS identificacao_credor,\n    c.cg_tipopessoa AS tipo_pesssoa_codigo\nFROM aplic2008.empenho e\nLEFT JOIN aplic2008.cadastro_geral c ON e.ent_codigo = c.ent_codigo AND c.exercicio >= 2015 AND e.cg_identificacao = c.cg_identificacao\nLEFT JOIN aplic2008.empenho_obra eo ON e.ent_codigo = eo.ent_codigo AND e.exercicio = eo.exercicio AND e.org_codigo = eo.org_codigo AND e.unor_codigo = eo.unor_codigo AND e.emp_numero = eo.emp_numero\nLEFT JOIN aplic2008.orgao o ON e.ent_codigo = o.ent_codigo AND e.exercicio = o.exercicio AND e.org_codigo = o.org_codigo\nLEFT JOIN aplic2008.unidade_orcamentaria u ON e.ent_codigo = u.ent_codigo AND e.exercicio = u.exercicio AND e.org_codigo = u.org_codigo AND e.unor_codigo = u.unor_codigo\nLEFT JOIN aplic2008.projeto_atividade prj ON e.ent_codigo = prj.ent_codigo AND e.exercicio = prj.exercicio AND e.prat_numero = prj.prat_numero AND e.prg_codigo = prj.prg_codigo\nLEFT JOIN aplic2008.fonte_recurso FRC ON E.FREC_CODIGO = FRC.FREC_CODIGO\nLEFT JOIN aplic2008.destinacao_recurso_iduso DRI ON E.EXERCICIO = DRI.EXERCICIO AND E.DRIDS_CODIGO = DRI.DRIDS_CODIGO\nLEFT JOIN aplic2008.destinacao_recurso_grupo DRG ON E.EXERCICIO = DRG.EXERCICIO AND E.DRGRP_CODIGO = DRG.DRGRP_CODIGO\nLEFT JOIN aplic2008.destinacao_recurso_especific DRE ON E.EXERCICIO = DRE.EXERCICIO AND E.DRESP_CODIGO = DRE.DRESP_CODIGO\nLEFT JOIN aplic2008.destinacao_recurso DR ON E.EXERCICIO = DR.EXERCICIO AND E.DESTREC_CODIGO = DR.DESTREC_CODIGO\nLEFT JOIN aplic2008.elemento_despesa el ON e.elde_codigo = el.elde_codigo AND e.exercicio = el.exercicio\nLEFT JOIN aplic2008.subelemento_despesa sub ON e.exercicio = sub.selde_exercicio AND e.elde_codigo = sub.elde_codigo AND e.selde_codigo = sub.selde_codigo\nLEFT JOIN aplic2008.funcao f ON e.fn_codigo = f.fn_codigo\nLEFT JOIN aplic2008.subfuncao sf ON e.sfn_codigo = sf.sfn_codigo\nINNER JOIN vw_entidade_aplic v ON e.ent_codigo = v.ent_codigo\nLEFT JOIN aplic2008.tipo_despesa_rpps t ON e.emp_tipodespesarpps = t.trpps_codigo\nLEFT JOIN aplic2008.contrato_empenho ce ON ce.ent_codigo = e.ent_codigo AND ce.exercicio = e.exercicio AND ce.org_codigo = e.org_codigo AND ce.unor_codigo = e.unor_codigo AND ce.emp_numero = e.emp_numero\nLEFT JOIN aplic2008.tipo_servico_engenharia TSE ON E.EMP_TIPOSERVICOENGENHARIA = TSE.TSENG_CODIGO\nLEFT JOIN aplic2008.modalidade_licitacao MLIC ON E.MLIC_CODIGO = MLIC.MLIC_CODIGO\n" +
		"LEFT JOIN aplic2008.fundamento_compra_direta FC ON FC.FCD_CODIGO = E.EMP_FUNDAMENTOCOMPRADIRETA\n" +
		"WHERE 1 = 1\n" +
		"AND e.ent_codigo = :1\n" +
		"AND e.exercicio = :2\n" +
		"AND E.MLIC_CODIGO = :3\n" +
		"ORDER BY v.mun_nome, v.ent_nome, e.emp_numero, e.emp_data"
	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano, codigoModalidadeLicitacao)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var filtroModalidadeLicitacoes []models.FiltroModalidadeLicitacao
	for rows.Next() {
		var filtroModalidadeLicitacao models.FiltroModalidadeLicitacao
		if err := rows.Scan(
			&filtroModalidadeLicitacao.Emp_Data,
			&filtroModalidadeLicitacao.Emp_Numero,
			&filtroModalidadeLicitacao.Credor,
			&filtroModalidadeLicitacao.Valor_Empenhado,
			&filtroModalidadeLicitacao.Valor_Liquidado,
			&filtroModalidadeLicitacao.Valor_Pago,
			&filtroModalidadeLicitacao.Valor_PagoRetencao,
			&filtroModalidadeLicitacao.Anulado_Empenho,
			&filtroModalidadeLicitacao.Qtde_Notas_Fiscais,
			&filtroModalidadeLicitacao.Qtde_NFe,
			&filtroModalidadeLicitacao.Contratos,
			&filtroModalidadeLicitacao.Ent_codigo,
			&filtroModalidadeLicitacao.Qtde_beneficiarios,
			&filtroModalidadeLicitacao.Instrumento_contrato,
			&filtroModalidadeLicitacao.Fonte_recurso_fonte,
			&filtroModalidadeLicitacao.Tipo_concurso,
			&filtroModalidadeLicitacao.Num_Concurso,
			&filtroModalidadeLicitacao.Assistencia_social,
			&filtroModalidadeLicitacao.Diarias,
			&filtroModalidadeLicitacao.Relevante,
			&filtroModalidadeLicitacao.Unidade_gestora,
			&filtroModalidadeLicitacao.Municipio,
			&filtroModalidadeLicitacao.Exercicio,
			&filtroModalidadeLicitacao.Codigo_ug,
			&filtroModalidadeLicitacao.Projeto_atividade,
			&filtroModalidadeLicitacao.Dotacao,
			&filtroModalidadeLicitacao.Num_obra,
			&filtroModalidadeLicitacao.Tipo_serv_engenharia,
			&filtroModalidadeLicitacao.Destinacao_recurso_grupo,
			&filtroModalidadeLicitacao.Destinacao_recurso_iduso,
			&filtroModalidadeLicitacao.Destinacao_recurso_codiog_iduso,
			&filtroModalidadeLicitacao.Codigo_destinacao_recurso,
			&filtroModalidadeLicitacao.Codigo_destinacao_recurso_especificacao,
			&filtroModalidadeLicitacao.Destinacao_recurso_codigo_grupo,
			&filtroModalidadeLicitacao.Fund_compra_direta_codigo,
			&filtroModalidadeLicitacao.Fund_compra_direta_descricao,
			&filtroModalidadeLicitacao.Destinacao_recurso,
			&filtroModalidadeLicitacao.Orgao,
			&filtroModalidadeLicitacao.UnidadeOrcamentaria,
			&filtroModalidadeLicitacao.Unidade_orcamentaria_codigo,
			&filtroModalidadeLicitacao.Funcao_codigo,
			&filtroModalidadeLicitacao.Funcao_descricao,
			&filtroModalidadeLicitacao.Subfuncao_codigo,
			&filtroModalidadeLicitacao.Subfuncao_descricao,
			&filtroModalidadeLicitacao.Programa_codigo,
			&filtroModalidadeLicitacao.Num_Projeto_Atividade,
			&filtroModalidadeLicitacao.Categoria_Economica,
			&filtroModalidadeLicitacao.Natureza_Despesa,
			&filtroModalidadeLicitacao.Modalidade_aplicacao_codigo,
			&filtroModalidadeLicitacao.Elemento_despesa_codigo,
			&filtroModalidadeLicitacao.Elemento_despesa_descricao,
			&filtroModalidadeLicitacao.Subelemento_despesa_codigo,
			&filtroModalidadeLicitacao.Subelemento_despesa,
			&filtroModalidadeLicitacao.Descricao,
			&filtroModalidadeLicitacao.Num_processo_licitatorio,
			&filtroModalidadeLicitacao.Tipo_Contrato,
			&filtroModalidadeLicitacao.Num_aditivo_contrato,
			&filtroModalidadeLicitacao.Num_convenio,
			&filtroModalidadeLicitacao.Num_aditivo_convenio,
			&filtroModalidadeLicitacao.Compra_direta,
			&filtroModalidadeLicitacao.Tipo,
			&filtroModalidadeLicitacao.Mes_referencia,
			&filtroModalidadeLicitacao.Identificacao_credor,
			&filtroModalidadeLicitacao.Tipo_pesssoaCodigo,
		); err != nil {
			return nil, err
		}
		filtroModalidadeLicitacoes = append(filtroModalidadeLicitacoes, filtroModalidadeLicitacao)
	}
	return filtroModalidadeLicitacoes, nil
}
