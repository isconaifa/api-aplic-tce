package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FiltroSubFuncaoRepository struct {
	db *sql.DB
}

func NewFiltroSubFuncaoRepository(db *sql.DB) *FiltroSubFuncaoRepository {
	return &FiltroSubFuncaoRepository{db: db}
}
func (repository *FiltroSubFuncaoRepository) GetAllFiltroSubFuncao(unidadeGestoraCodigo, ano string, codigoSubFuncao int) ([]models.FiltroSubfuncao, error) {
	query := "SELECT DISTINCT\n    e.emp_data AS \"empenho_data\",\n    e.emp_numero AS \"empenho_numero\",\n  " +
		"  c.cg_nome AS \"credor\",\n " +
		"   e.emp_valor - (select nvl(sum(a.anul_valor),0) from aplic2008.anulacao_empenho a where a.ent_codigo = e.ent_codigo and a.exercicio = e.exercicio and a.org_codigo = e.org_codigo and a.unor_codigo = e.unor_codigo and a.emp_numero = e.emp_numero) AS \"valor_empenhado\",\n    (select nvl(sum(l.liq_valor),0) from aplic2008.liquidacao_empenho l where l.ENT_CODIGO = E.ENT_CODIGO and l.EXERCICIO = E.EXERCICIO and l.ORG_CODIGO = E.ORG_CODIGO and l.UNOR_CODIGO = E.UNOR_CODIGO and l.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(al.anul_valor),0) from aplic2008.anulacao_liquidacao_empenho al where al.ent_codigo = e.ent_codigo and al.exercicio = e.exercicio and al.org_codigo = e.org_codigo and al.unor_codigo = e.unor_codigo and al.emp_numero = e.emp_numero) AS \"valor_liquidado\",\n    (select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo and pel.emp_numero = e.emp_numero) AS \"valor_pago\",\n    (select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo and pel.emp_numero = e.emp_numero) + (select coalesce(sum(DLIQ_VALOR),0) from aplic2008.desconto_liquidado d where d.ent_codigo = e.ent_codigo and d.exercicio = e.exercicio and d.org_codigo = e.org_codigo and d.unor_codigo = e.unor_codigo and d.emp_numero = e.emp_numero) - (select nvl(sum(aedl.aedliq_valor),0) from aplic2008.anulacao_estorno_desc_liquidad aedl where aedl.ent_codigo = e.ent_codigo and aedl.exercicio = e.exercicio and aedl.org_codigo = e.org_codigo and aedl.unor_codigo = e.unor_codigo and aedl.emp_numero = e.emp_numero) AS \"valor_pago_mais_retencao\",\n    (select nvl(sum(a.anul_valor),0) from aplic2008.anulacao_empenho a where a.ent_codigo = e.ent_codigo and a.exercicio = e.exercicio and a.org_codigo = e.org_codigo and a.unor_codigo = e.unor_codigo and a.emp_numero = e.emp_numero) AS \"valor_anulado\",\n    (select count(1) from aplic2008.NOTA_FISCAL N where N.ENT_CODIGO = E.ENT_CODIGO and N.EXERCICIO = E.EXERCICIO and N.ORG_CODIGO = E.ORG_CODIGO and N.UNOR_CODIGO = E.UNOR_CODIGO and N.EMP_NUMERO = E.EMP_NUMERO) AS \"qtde_notas_fiscais\",\n    (select count(1) FROM aplic2008.nota_fiscal n WHERE n.ent_codigo = e.ent_codigo and n.exercicio = e.exercicio and n.org_codigo = e.org_codigo and n.unor_codigo = e.unor_codigo and n.emp_numero = e.emp_numero and n.ntfsc_numeronfe is not null) AS \"qtde_nfe\",\n    (select count(1) FROM aplic2008.contrato_empenho n WHERE n.ent_codigo = e.ent_codigo and n.exercicio = e.exercicio and n.org_codigo = e.org_codigo and n.unor_codigo = e.unor_codigo and n.emp_numero = e.emp_numero) AS \"contratos\",\n    e.ent_codigo AS \"ent_codigo\",\n    (select Count(1) from aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas where bas.ENT_CODIGO = E.ENT_CODIGO and bas.EXERCICIO = E.EXERCICIO and bas.ORG_CODIGO = E.ORG_CODIGO and bas.UNOR_CODIGO = E.UNOR_CODIGO and bas.EMP_NUMERO = E.EMP_NUMERO) AS \"qtde_beneficiarios\",\n    CASE emp_instrumentocontrato WHEN 1 THEN 'Contrato' WHEN 2 THEN 'Carta-Contrato' WHEN 3 THEN 'Nota de Empenho da Despesa' WHEN 4 THEN 'Autorização de Compra' WHEN 5 THEN 'Ordem de Execução de Serviço' WHEN 6 THEN 'Outros Instrumentos Hábeis' END AS \"instrumento_contrato\",\n    FRC.FREC_DESCRICAO AS \"fonte_recurso_fonte\",\n    e.conc_tipo AS \"tipo_concurso\",\n    e.conc_numero AS \"num_concurso\",\n    CASE e.emp_benefassistenciasocial WHEN '1' THEN 'Não existem beneficiários' WHEN '2' THEN 'Existem beneficiários sem cadastro informatizado' WHEN '3' THEN 'Existem beneficiários com cadastro informatizado' END AS \"assistencia_social\",\n    (select Count(1) from aplic2008.diaria dir where dir.ent_codigo = e.ent_codigo and dir.exercicio = e.exercicio and dir.org_codigo = e.org_codigo and dir.unor_codigo = e.unor_codigo and dir.emp_numero = e.emp_numero) AS \"diarias\",\n    '' AS \"relevante\",\n    v.ent_nome AS \"unidade_gestora\",\n    v.mun_nome AS \"municipio\",\n    e.exercicio AS \"exercicio\",\n    e.ent_codigo AS \"codigo_ug\",\n    prj.prat_descricao AS \"projeto_atividade\",\n    e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS \"dotacao\",\n " +
		"   aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) AS \"num_obra\",\n    TSE.TSENG_DESCRICAO AS \"tipo_serv_engenharia\",\n    upper(drg.drgrp_descricao) AS \"destinacao_recurso_grupo\",\n  " +
		"  upper(dri.drids_descricao) AS \"destinacao_recurso_iduso\",\n " +
		"   e.drids_codigo AS \"destinacao_recurso_codiog_iduso\",\n  " +
		"  e.destrec_codigo AS \"codigo_destinacao_recurso\",\n  " +
		"  e.dresp_codigo AS \"codigo_destinacao_recurso_especificacao\",\n " +
		"   e.drgrp_codigo AS \"destinacao_recurso_codigo_grupo\",\n  " +
		"  E.EMP_FUNDAMENTOCOMPRADIRETA AS \"fund_compra_direta_codigo\",\n " +
		"   FC.FCD_DESCRICAO AS \"fund_compra_direta_descricao\",\n " +
		"   upper(dr.destrec_descricao) AS \"destinacao_recurso\",\n " +
		"   o.org_nome AS \"orgao\",\n    u.unor_nome AS \"unidade_orcamentaria\",\n  " +
		"  e.unor_codigo AS \"unidade_orcamentaria_codigo\",\n    e.fn_codigo AS \"funcao_codigo\",\n " +
		"   f.fn_descricao AS \"funcao_descricao\",\n    e.sfn_codigo AS \"subfuncao_codigo\",\n  " +
		"  sf.sfn_descricao AS \"subfuncao_descricao\",\n    e.prg_codigo AS \"programa_codigo\",\n  " +
		"  e.prat_numero AS \"num_projeto_atividade\",\n    e.ctec_codigo AS \"categoria_economica\",\n " +
		"   e.ndesp_codigo AS \"natureza_despesa\",\n    e.mdap_codigo AS \"modalidade_aplicacao_codigo\",\n  " +
		"  e.elde_codigo AS \"elemento_despesa_codigo\",\n    el.elde_descricao AS \"elemento_despesa_descricao\",\n  " +
		"  e.selde_codigo AS \"subelemento_despesa_codigo\",\n    sub.selde_descricao AS \"subelemento_despesa\",\n " +
		"   e.emp_descricao AS \"descricao\",\n    e.plic_numero AS \"num_processo_licitatorio\",\n " +
		"   e.cont_tipo AS \"tipo_contrato\",\n    e.cont_numaditivo AS \"num_aditivo_contrato\",\n  " +
		"  e.conv_numero AS \"num_convenio\",\n    e.conv_numaditivo AS \"num_aditivo_convenio\",\n  " +
		"  CASE e.emp_compradiretaprocesso WHEN '1' THEN 'Não' WHEN '2' THEN 'Sim' WHEN 'N' THEN 'Não' WHEN 'S' THEN 'Sim' END AS \"compra_direta\",\n  " +
		"  CASE e.emp_tipo WHEN '1' THEN 'Estimativo' WHEN '2' THEN 'Global' WHEN '3' THEN 'Ordinário' END AS \"tipo\",\n " +
		"   e.mesreferencia AS \"mes_referencia\",\n    e.cg_identificacao AS \"identificacao_credor\",\n " +
		"   c.cg_tipopessoa AS \"tipo_pesssoa_codigo\"\nFROM aplic2008.empenho e\nLEFT JOIN aplic2008.cadastro_geral c ON (e.ent_codigo = c.ent_codigo and c.exercicio >= 2015 and e.cg_identificacao = c.cg_identificacao)\nLEFT JOIN aplic2008.empenho_obra eo ON ((e.ent_codigo = eo.ent_codigo) and (e.exercicio = eo.exercicio) and (e.org_codigo = eo.org_codigo) and (e.unor_codigo = eo.unor_codigo) and (e.emp_numero = eo.emp_numero))\nLEFT JOIN aplic2008.orgao o ON (e.ent_codigo = o.ent_codigo and e.exercicio = o.exercicio and e.org_codigo = o.org_codigo)\nLEFT JOIN aplic2008.unidade_orcamentaria u ON (e.ent_codigo = u.ent_codigo and e.exercicio = u.exercicio and e.org_codigo = u.org_codigo and e.unor_codigo = u.unor_codigo)\n" +
		"LEFT JOIN aplic2008.projeto_atividade prj on (e.ent_codigo = prj.ent_codigo and e.exercicio = prj.exercicio and e.prat_numero = prj.prat_numero and e.prg_codigo = prj.prg_codigo)\nLEFT JOIN aplic2008.FONTE_RECURSO FRC ON (E.FREC_CODIGO = FRC.FREC_CODIGO)\n" +
		"LEFT JOIN aplic2008.DESTINACAO_RECURSO_IDUSO DRI on (E.EXERCICIO = DRI.EXERCICIO and E.DRIDS_CODIGO = DRI.DRIDS_CODIGO)\n" +
		"LEFT JOIN aplic2008.DESTINACAO_RECURSO_GRUPO DRG on (E.EXERCICIO = DRG.EXERCICIO and E.DRGRP_CODIGO = DRG.DRGRP_CODIGO)\n" +
		"LEFT JOIN aplic2008.DESTINACAO_RECURSO_ESPECIFIC DRE on (E.EXERCICIO = DRE.EXERCICIO and E.DRESP_CODIGO = DRE.DRESP_CODIGO)\n" +
		"LEFT JOIN aplic2008.DESTINACAO_RECURSO DR on (E.EXERCICIO = DR.EXERCICIO and E.DESTREC_CODIGO = DR.DESTREC_CODIGO)\n" +
		"LEFT JOIN aplic2008.elemento_despesa el on (e.elde_codigo = el.elde_codigo and e.exercicio = el.exercicio)\n" +
		"LEFT JOIN aplic2008.subelemento_despesa sub on (e.exercicio = sub.selde_exercicio and e.elde_codigo = sub.elde_codigo and e.selde_codigo = sub.selde_codigo)\n" +
		"LEFT JOIN aplic2008.funcao f on (e.fn_codigo = f.fn_codigo)\nLEFT JOIN aplic2008.subfuncao sf on (e.sfn_codigo = sf.sfn_codigo)\nINNER JOIN vw_entidade_aplic v on (e.ent_codigo = v.ent_codigo)\nLEFT JOIN aplic2008.tipo_despesa_rpps t on (e.emp_tipodespesarpps = t.trpps_codigo)\n" +
		"LEFT JOIN aplic2008.contrato_empenho ce on (ce.ent_codigo = e.ent_codigo and ce.exercicio = e.exercicio and ce.org_codigo = e.org_codigo and ce.unor_codigo = e.unor_codigo and ce.emp_numero = e.emp_numero)\nLEFT JOIN aplic2008.TIPO_SERVICO_ENGENHARIA TSE ON (E.EMP_TIPOSERVICOENGENHARIA = TSE.TSENG_CODIGO)\nLEFT JOIN aplic2008.MODALIDADE_LICITACAO MLIC ON (E.MLIC_CODIGO = MLIC.MLIC_CODIGO)\nLEFT JOIN aplic2008.FUNDAMENTO_COMPRA_DIRETA FC on FC.FCD_CODIGO = E.EMP_FUNDAMENTOCOMPRADIRETA\nWHERE 1 = 1\n " +
		" AND e.ent_codigo = :1\n " +
		" AND e.exercicio = :2\n " +
		" AND e.sfn_codigo = :3\n" +
		"ORDER BY v.mun_nome, v.ent_nome, e.emp_numero, e.emp_data"
	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano, codigoSubFuncao)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var filtroSubFuncoes []models.FiltroSubfuncao
	for rows.Next() {
		var filtroSubFuncao models.FiltroSubfuncao
		if err := rows.Scan(
			&filtroSubFuncao.Emp_Data,
			&filtroSubFuncao.Emp_Numero,
			&filtroSubFuncao.Credor,
			&filtroSubFuncao.Valor_Empenhado,
			&filtroSubFuncao.Valor_Liquidado,
			&filtroSubFuncao.Valor_Pago,
			&filtroSubFuncao.Valor_PagoRetencao,
			&filtroSubFuncao.Anulado_Empenho,
			&filtroSubFuncao.Qtde_Notas_Fiscais,
			&filtroSubFuncao.Qtde_NFe,
			&filtroSubFuncao.Contratos,
			&filtroSubFuncao.Ent_codigo,
			&filtroSubFuncao.Qtde_beneficiarios,
			&filtroSubFuncao.Instrumento_contrato,
			&filtroSubFuncao.Fonte_recurso_fonte,
			&filtroSubFuncao.Tipo_concurso,
			&filtroSubFuncao.Num_Concurso,
			&filtroSubFuncao.Assistencia_social,
			&filtroSubFuncao.Diarias,
			&filtroSubFuncao.Relevante,
			&filtroSubFuncao.Unidade_gestora,
			&filtroSubFuncao.Municipio,
			&filtroSubFuncao.Exercicio,
			&filtroSubFuncao.Codigo_ug,
			&filtroSubFuncao.Projeto_atividade,
			&filtroSubFuncao.Dotacao,
			&filtroSubFuncao.Num_obra,
			&filtroSubFuncao.Tipo_serv_engenharia,
			&filtroSubFuncao.Destinacao_recurso_grupo,
			&filtroSubFuncao.Destinacao_recurso_iduso,
			&filtroSubFuncao.Destinacao_recurso_codiog_iduso,
			&filtroSubFuncao.Codigo_destinacao_recurso,
			&filtroSubFuncao.Codigo_destinacao_recurso_especificacao,
			&filtroSubFuncao.Destinacao_recurso_codigo_grupo,
			&filtroSubFuncao.Fund_compra_direta_codigo,
			&filtroSubFuncao.Fund_compra_direta_descricao,
			&filtroSubFuncao.Destinacao_recurso,
			&filtroSubFuncao.Orgao,
			&filtroSubFuncao.UnidadeOrcamentaria,
			&filtroSubFuncao.Unidade_orcamentaria_codigo,
			&filtroSubFuncao.Funcao_codigo,
			&filtroSubFuncao.Funcao_descricao,
			&filtroSubFuncao.Subfuncao_codigo,
			&filtroSubFuncao.Subfuncao_descricao,
			&filtroSubFuncao.Programa_codigo,
			&filtroSubFuncao.Num_Projeto_Atividade,
			&filtroSubFuncao.Categoria_Economica,
			&filtroSubFuncao.Natureza_Despesa,
			&filtroSubFuncao.Modalidade_aplicacao_codigo,
			&filtroSubFuncao.Elemento_despesa_codigo,
			&filtroSubFuncao.Elemento_despesa_descricao,
			&filtroSubFuncao.Subelemento_despesa_codigo,
			&filtroSubFuncao.Subelemento_despesa,
			&filtroSubFuncao.Descricao,
			&filtroSubFuncao.Num_processo_licitatorio,
			&filtroSubFuncao.Tipo_Contrato,
			&filtroSubFuncao.Num_aditivo_contrato,
			&filtroSubFuncao.Num_convenio,
			&filtroSubFuncao.Num_aditivo_convenio,
			&filtroSubFuncao.Compra_direta,
			&filtroSubFuncao.Tipo,
			&filtroSubFuncao.Mes_referencia,
			&filtroSubFuncao.Identificacao_credor,
			&filtroSubFuncao.Tipo_pesssoaCodigo,
		); err != nil {
			return nil, err
		}
		filtroSubFuncoes = append(filtroSubFuncoes, filtroSubFuncao)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return filtroSubFuncoes, nil
}
