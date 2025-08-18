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

func (repository *FiltroUnidadeOrcamentariaRepository) GetAllFiltroUnidadeOrcamentaria(unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria string, codigoDestinacaoRecurso int) ([]models.FiltroUnidadeOrcamentaria, error) {
	query := "SELECT DISTINCT\n    e.emp_data AS \"Emp_Data\",\n    e.emp_numero AS \"Emp_Numero\",\n    c.cg_nome AS \"Credor\",\n    e.emp_valor - (SELECT NVL(SUM(a.anul_valor), 0) FROM aplic2008.anulacao_empenho a WHERE a.ent_codigo = e.ent_codigo AND a.exercicio = e.exercicio AND a.org_codigo = e.org_codigo AND a.unor_codigo = e.unor_codigo AND a.emp_numero = e.emp_numero) AS \"Valor_Empenhado\",\n    (SELECT NVL(SUM(l.liq_valor), 0) FROM aplic2008.liquidacao_empenho l WHERE l.ENT_CODIGO = E.ENT_CODIGO AND l.EXERCICIO = E.EXERCICIO AND l.ORG_CODIGO = E.ORG_CODIGO AND l.UNOR_CODIGO = E.UNOR_CODIGO AND l.EMP_NUMERO = E.EMP_NUMERO) - (SELECT NVL(SUM(al.anul_valor), 0) FROM aplic2008.anulacao_liquidacao_empenho al WHERE al.ent_codigo = e.ent_codigo AND al.exercicio = e.exercicio AND al.org_codigo = e.org_codigo AND al.unor_codigo = e.unor_codigo AND al.emp_numero = e.emp_numero) AS \"Valor_Liquidado\",\n    (SELECT NVL(SUM(p.pgto_valor), 0) FROM aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl WHERE P.ENT_CODIGO = PL.ENT_CODIGO AND P.EXERCICIO = PL.EXERCICIO AND P.PGTO_NUMERO = PL.PGTO_NUMERO AND PL.ENT_CODIGO = E.ENT_CODIGO AND PL.EXERCICIO = E.EXERCICIO AND PL.ORG_CODIGO = E.ORG_CODIGO AND PL.UNOR_CODIGO = E.UNOR_CODIGO AND PL.EMP_NUMERO = E.EMP_NUMERO) - (SELECT NVL(SUM(alp.anul_valor), 0) FROM aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel WHERE alp.ent_codigo = pel.ent_codigo AND alp.exercicio = pel.exercicio AND alp.pgto_numero = pel.pgto_numero AND pel.ent_codigo = e.ent_codigo AND pel.org_codigo = e.org_codigo AND pel.unor_codigo = e.unor_codigo AND pel.emp_numero = e.emp_numero) AS \"Valor_Pago\",\n    (SELECT NVL(SUM(p.pgto_valor), 0) FROM aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl WHERE P.ENT_CODIGO = PL.ENT_CODIGO AND P.EXERCICIO = PL.EXERCICIO AND P.PGTO_NUMERO = PL.PGTO_NUMERO AND PL.ENT_CODIGO = E.ENT_CODIGO AND PL.EXERCICIO = E.EXERCICIO AND PL.ORG_CODIGO = E.ORG_CODIGO AND PL.UNOR_CODIGO = E.UNOR_CODIGO AND PL.EMP_NUMERO = E.EMP_NUMERO) - (SELECT NVL(SUM(alp.anul_valor), 0) FROM aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel WHERE alp.ent_codigo = pel.ent_codigo AND alp.exercicio = pel.exercicio AND alp.pgto_numero = pel.pgto_numero AND pel.ent_codigo = e.ent_codigo AND pel.org_codigo = e.org_codigo AND pel.unor_codigo = e.unor_codigo AND pel.emp_numero = e.emp_numero) + (SELECT COALESCE(SUM(DLIQ_VALOR), 0) FROM aplic2008.desconto_liquidado d WHERE d.ent_codigo = e.ent_codigo AND d.exercicio = e.exercicio AND d.org_codigo = e.org_codigo AND d.unor_codigo = e.unor_codigo AND d.emp_numero = e.emp_numero) - (SELECT NVL(SUM(aedl.aedliq_valor), 0) FROM aplic2008.anulacao_estorno_desc_liquidad aedl WHERE aedl.ent_codigo = e.ent_codigo AND aedl.exercicio = e.exercicio AND aedl.org_codigo = e.org_codigo AND aedl.unor_codigo = e.unor_codigo AND aedl.emp_numero = e.emp_numero) AS \"Valor_PagoRetencao\",\n    (SELECT NVL(SUM(a.anul_valor), 0) FROM aplic2008.anulacao_empenho a WHERE a.ent_codigo = e.ent_codigo AND a.exercicio = e.exercicio AND a.org_codigo = e.org_codigo AND a.unor_codigo = e.unor_codigo AND a.emp_numero = e.emp_numero) AS \"Anulado_Empenho\",\n    (SELECT COUNT(1) FROM aplic2008.NOTA_FISCAL N WHERE N.ENT_CODIGO = E.ENT_CODIGO AND N.EXERCICIO = E.EXERCICIO AND N.ORG_CODIGO = E.ORG_CODIGO AND N.UNOR_CODIGO = E.UNOR_CODIGO AND N.EMP_NUMERO = E.EMP_NUMERO) AS \"Qtde_Notas_Fiscais\",\n    (SELECT COUNT(1) FROM aplic2008.nota_fiscal n WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo AND n.unor_codigo = e.unor_codigo AND n.emp_numero = e.emp_numero AND n.ntfsc_numeronfe IS NOT NULL) AS \"Qtde_NFe\",\n    (SELECT COUNT(1) FROM aplic2008.contrato_empenho n WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo AND n.unor_codigo = e.unor_codigo AND n.emp_numero = e.emp_numero) AS \"Contratos\",\n    e.ent_codigo AS \"Ent_codigo\",\n    (SELECT COUNT(1) FROM aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas WHERE bas.ENT_CODIGO = E.ENT_CODIGO AND bas.EXERCICIO = E.EXERCICIO AND bas.ORG_CODIGO = E.ORG_CODIGO AND bas.UNOR_CODIGO = E.UNOR_CODIGO AND bas.EMP_NUMERO = E.EMP_NUMERO) AS \"Qtde_beneficiarios\",\n    DECODE(e.emp_instrumentocontrato, '1', 'Contrato', '2', 'Carta-Contrato', '3', 'Nota de Empenho da Despesa', '4', 'Autorização de Compra', '5', 'Ordem de Execução de Serviço', '6', 'Outros Instrumentos Hábeis') AS \"Instrumento_contrato\",\n    FRC.FREC_DESCRICAO AS \"Fonte_recurso_fonte\",\n    e.conc_tipo AS \"Tipo_concurso\",\n    e.conc_numero AS \"Num_Concurso\",\n    DECODE(e.emp_benefassistenciasocial, '1', 'Não existem beneficiários', '2', 'Existem beneficiários sem cadastro informatizado', '3', 'Existem beneficiários com cadastro informatizado') AS \"Assistencia_social\",\n    (SELECT COUNT(1) FROM aplic2008.diaria dir WHERE dir.ent_codigo = e.ent_codigo AND dir.exercicio = e.exercicio AND dir.org_codigo = e.org_codigo AND dir.unor_codigo = e.unor_codigo AND dir.emp_numero = e.emp_numero) AS \"Diarias\",\n    ' ' AS \"Relevante\",\n    v.ent_nome AS \"Unidade_gestora\",\n    v.mun_nome AS \"Municipio\",\n    e.exercicio AS \"Exercicio\",\n    e.ent_codigo AS \"Codigo_ug\",\n    prj.prat_descricao AS \"Projeto_atividade\",\n    e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS \"Dotacao\",\n    aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) AS \"Num_obra\",\n    TSE.TSENG_DESCRICAO AS \"Tipo_serv_engenharia\",\n    UPPER(drg.drgrp_descricao) AS \"Destinacao_recurso_grupo\",\n    UPPER(dri.drids_descricao) AS \"Destinacao_recurso_iduso\",\n    e.drids_codigo AS \"Destinacao_recurso_codiog_iduso\",\n    e.destrec_codigo AS \"Codigo_destinacao_recurso\",\n    e.dresp_codigo AS \"Codigo_destinacao_recurso_especificacao\",\n    e.drgrp_codigo AS \"Destinacao_recurso_codigo_grupo\",\n    E.EMP_FUNDAMENTOCOMPRADIRETA AS \"Fund_compra_direta_codigo\",\n    FC.FCD_DESCRICAO AS \"Fund_compra_direta_descricao\",\n    UPPER(dr.destrec_descricao) AS \"Destinacao_recurso\",\n    o.org_nome AS \"Orgao\",\n    u.unor_nome AS \"UnidadeOrcamentaria\",\n    e.unor_codigo AS \"Unidade_orcamentaria_codigo\",\n    e.fn_codigo AS \"Funcao_codigo\",\n    f.fn_descricao AS \"Funcao_descricao\",\n    e.sfn_codigo AS \"Subfuncao_codigo\",\n    sf.sfn_descricao AS \"Subfuncao_descricao\",\n    e.prg_codigo AS \"Programa_codigo\",\n    e.prat_numero AS \"Num_Projeto_Atividade\",\n    e.ctec_codigo AS \"Categoria_Economica\",\n    e.ndesp_codigo AS \"Natureza_Despesa\",\n    e.mdap_codigo AS \"Modalidade_aplicacao_codigo\",\n    e.elde_codigo AS \"Elemento_despesa_codigo\",\n    el.elde_descricao AS \"Elemento_despesa_descricao\",\n    e.selde_codigo AS \"Subelemento_despesa_codigo\",\n    sub.selde_descricao AS \"Subelemento_despesa\",\n    e.emp_descricao AS \"Descricao\",\n    e.plic_numero AS \"Num_processo_licitatorio\",\n    e.cont_tipo AS \"Tipo_Contrato\",\n    e.cont_numaditivo AS \"Num_aditivo_contrato\",\n    e.conv_numero AS \"Num_convenio\",\n    e.conv_numaditivo AS \"Num_aditivo_convenio\",\n    DECODE(e.emp_compradiretaprocesso, '1', 'Não', '2', 'Sim', 'N', 'Não', 'S', 'Sim') AS \"Compra_direta\",\n    DECODE(e.emp_tipo, '1', 'Estimativo', '2', 'Global', '3', 'Ordinário') AS \"Tipo\",\n    e.mesreferencia AS \"Mes_referencia\",\n    e.cg_identificacao AS \"Identificacao_credor\",\n    c.cg_tipopessoa AS \"Tipo_pesssoaCodigo\"\nFROM\n    aplic2008.empenho e\nLEFT JOIN aplic2008.cadastro_geral c ON (e.ent_codigo = c.ent_codigo AND c.exercicio >= 2015 AND e.cg_identificacao = c.cg_identificacao)\nLEFT JOIN aplic2008.empenho_obra eo ON (e.ent_codigo = eo.ent_codigo AND e.exercicio = eo.exercicio AND e.org_codigo = eo.org_codigo AND e.unor_codigo = eo.unor_codigo AND e.emp_numero = eo.emp_numero)\nLEFT JOIN aplic2008.orgao o ON (e.ent_codigo = o.ent_codigo AND e.exercicio = o.exercicio AND e.org_codigo = o.org_codigo)\nLEFT JOIN aplic2008.unidade_orcamentaria u ON (e.ent_codigo = u.ent_codigo AND e.exercicio = u.exercicio AND e.org_codigo = u.org_codigo AND e.unor_codigo = u.unor_codigo)\nLEFT JOIN aplic2008.projeto_atividade prj ON (e.ent_codigo = prj.ent_codigo AND e.exercicio = prj.exercicio AND e.prat_numero = prj.prat_numero AND e.prg_codigo = prj.prg_codigo)\nLEFT JOIN aplic2008.FONTE_RECURSO FRC ON (E.FREC_CODIGO = FRC.FREC_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO_IDUSO DRI ON (E.EXERCICIO = DRI.EXERCICIO AND E.DRIDS_CODIGO = DRI.DRIDS_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO_GRUPO DRG ON (E.EXERCICIO = DRG.EXERCICIO AND E.DRGRP_CODIGO = DRG.DRGRP_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO_ESPECIFIC DRE ON (E.EXERCICIO = DRE.EXERCICIO AND E.DRESP_CODIGO = DRE.DRESP_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO DR ON (E.EXERCICIO = DR.EXERCICIO AND E.DESTREC_CODIGO = DR.DESTREC_CODIGO)\nLEFT JOIN aplic2008.elemento_despesa el ON (e.elde_codigo = el.elde_codigo AND e.exercicio = el.exercicio)\nLEFT JOIN aplic2008.subelemento_despesa sub ON (e.exercicio = sub.selde_exercicio AND e.elde_codigo = sub.elde_codigo AND e.selde_codigo = sub.selde_codigo)\nLEFT JOIN aplic2008.funcao f ON (e.fn_codigo = f.fn_codigo)\nLEFT JOIN aplic2008.subfuncao sf ON (e.sfn_codigo = sf.sfn_codigo)\nINNER JOIN vw_entidade_aplic v ON (e.ent_codigo = v.ent_codigo)\nLEFT JOIN aplic2008.tipo_despesa_rpps t ON (e.emp_tipodespesarpps = t.trpps_codigo)\nLEFT JOIN aplic2008.contrato_empenho ce ON (ce.ent_codigo = e.ent_codigo AND ce.exercicio = e.exercicio AND ce.org_codigo = e.org_codigo AND ce.unor_codigo = e.unor_codigo AND ce.emp_numero = e.emp_numero)\nLEFT JOIN aplic2008.TIPO_SERVICO_ENGENHARIA TSE ON (E.EMP_TIPOSERVICOENGENHARIA = TSE.TSENG_CODIGO)\nLEFT JOIN aplic2008.MODALIDADE_LICITACAO MLIC ON (E.MLIC_CODIGO = MLIC.MLIC_CODIGO)\n" +
		"LEFT JOIN aplic2008.FUNDAMENTO_COMPRA_DIRETA FC ON FC.FCD_CODIGO = E.EMP_FUNDAMENTOCOMPRADIRETA\nWHERE\n  " +
		"  e.ent_codigo = :1\n " +
		"   AND e.exercicio = :2\n  " +
		"  AND e.org_codigo = :3\n  " +
		"  AND e.unor_codigo = :4\n  " +
		"  AND e.destrec_codigo = :5\n" +
		"ORDER BY\n    v.mun_nome,\n    v.ent_nome,\n    e.emp_numero,\n    e.emp_data"

	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano, codigoOrgao, codigoUnidadeOrcamentaria, codigoDestinacaoRecurso)
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
