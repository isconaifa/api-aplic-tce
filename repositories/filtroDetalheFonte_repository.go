package repositories

import (
	"api-aplic-web/models"
	"database/sql"
)

type FiiltroDetalheFonteRepository struct {
	db *sql.DB
}

func NewFiiltroDetalheFonteRepository(db *sql.DB) *FiiltroDetalheFonteRepository {
	return &FiiltroDetalheFonteRepository{db: db}
}
func (repository *FiiltroDetalheFonteRepository) GetAllFiiltroDetalheFonte(unidadeGestoraCodigo, ano string, codigoDestinacaoRecurso int) ([]models.FiltroDetalheFonte, error) {
	query := "SELECT DISTINCT\n    e.emp_data AS \"Emp_Data\",\n    e.emp_numero AS \"Emp_Numero\",\n    c.cg_nome AS \"Credor\",\n    e.emp_valor - (select nvl(sum(a.anul_valor),0) from aplic2008.anulacao_empenho a where a.ent_codigo = e.ent_codigo and a.exercicio = e.exercicio and a.org_codigo = e.org_codigo and a.unor_codigo = e.unor_codigo and a.emp_numero = e.emp_numero) AS \"Valor_Empenhado\",\n    (select nvl(sum(l.liq_valor),0) from aplic2008.liquidacao_empenho l where l.ENT_CODIGO = E.ENT_CODIGO and l.EXERCICIO = E.EXERCICIO and l.ORG_CODIGO = E.ORG_CODIGO and l.UNOR_CODIGO = E.UNOR_CODIGO and l.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(al.anul_valor),0) from aplic2008.anulacao_liquidacao_empenho al where al.ent_codigo = e.ent_codigo and al.exercicio = e.exercicio and al.org_codigo = e.org_codigo and al.unor_codigo = e.unor_codigo and al.emp_numero = e.emp_numero) AS \"Valor_Liquidado\",\n    (select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo and pel.emp_numero = e.emp_numero) AS \"Valor_Pago\",\n    (select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p, aplic2008.pagamento_empenho_liquidacao pl where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp, aplic2008.pagamento_empenho_liquidacao pel where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo and pel.emp_numero = e.emp_numero) + (select coalesce(sum(DLIQ_VALOR),0) from aplic2008.desconto_liquidado d where d.ent_codigo = e.ent_codigo and d.exercicio = e.exercicio and d.org_codigo = e.org_codigo and d.unor_codigo = e.unor_codigo and d.emp_numero = e.emp_numero) - (select nvl(sum(aedl.aedliq_valor),0) from aplic2008.anulacao_estorno_desc_liquidad aedl where aedl.ent_codigo = e.ent_codigo and aedl.exercicio = e.exercicio and aedl.org_codigo = e.org_codigo and aedl.unor_codigo = e.unor_codigo and aedl.emp_numero = e.emp_numero) AS \"Valor_PagoRetencao\",\n    (select nvl(sum(a.anul_valor),0) from aplic2008.anulacao_empenho a where a.ent_codigo = e.ent_codigo and a.exercicio = e.exercicio and a.org_codigo = e.org_codigo and a.unor_codigo = e.unor_codigo and a.emp_numero = e.emp_numero) AS \"Anulado_Empenho\",\n    (select count(1) from aplic2008.NOTA_FISCAL N where N.ENT_CODIGO = E.ENT_CODIGO and N.EXERCICIO = E.EXERCICIO and N.ORG_CODIGO = E.ORG_CODIGO and N.UNOR_CODIGO = E.UNOR_CODIGO and N.EMP_NUMERO = E.EMP_NUMERO) AS \"Qtde_Notas_Fiscais\",\n    (select count(1) from aplic2008.nota_fiscal n where n.ent_codigo = e.ent_codigo and n.exercicio = e.exercicio and n.org_codigo = e.org_codigo and n.unor_codigo = e.unor_codigo and n.emp_numero = e.emp_numero and n.ntfsc_numeronfe is not null) AS \"Qtde_NFe\",\n    (select count(1) from aplic2008.contrato_empenho n where n.ent_codigo = e.ent_codigo and n.exercicio = e.exercicio and n.org_codigo = e.org_codigo and n.unor_codigo = e.unor_codigo and n.emp_numero = e.emp_numero) AS \"Contratos\",\n    e.ent_codigo AS \"Ent_codigo\",\n    (select Count(1) from aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas where bas.ENT_CODIGO = E.ENT_CODIGO and bas.EXERCICIO = E.EXERCICIO and bas.ORG_CODIGO = E.ORG_CODIGO and bas.UNOR_CODIGO = E.UNOR_CODIGO and bas.EMP_NUMERO = E.EMP_NUMERO) AS \"Qtde_beneficiarios\",\n    decode(emp_instrumentocontrato,'1','Contrato','2','Carta-Contrato','3','Nota de Empenho da Despesa','4','Autorização de Compra','5','Ordem de Execução de Serviço','6','Outros Instrumentos Hábeis') AS \"Instrumento_contrato\",\n    FRC.FREC_DESCRICAO AS \"Fonte_recurso_fonte\",\n    e.conc_tipo AS \"Tipo_concurso\",\n    e.conc_numero AS \"Num_Concurso\",\n    decode(e.emp_benefassistenciasocial, '1', 'Não existem beneficiários','2', 'Existem beneficiários sem cadastro informatizado','3', 'Existem beneficiários com cadastro informatizado') AS \"Assistencia_social\",\n    (select Count(1) from aplic2008.diaria dir where dir.ent_codigo = e.ent_codigo and dir.exercicio = e.exercicio and dir.org_codigo = e.org_codigo and dir.unor_codigo = e.unor_codigo and dir.emp_numero = e.emp_numero) AS \"Diarias\",\n    ' ' AS \"Relevante\",\n    v.ent_nome AS \"Unidade_gestora\",\n    v.mun_nome AS \"Municipio\",\n    e.exercicio AS \"Exercicio\",\n    e.ent_codigo AS \"Codigo_ug\",\n    prj.prat_descricao AS \"Projeto_atividade\",\n    e.ctec_codigo || '.' || e.ndesp_codigo || '.' || e.mdap_codigo || '.' || e.elde_codigo || '.' || e.selde_codigo AS \"Dotacao\",\n    aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) AS \"Num_obra\",\n    TSE.TSENG_DESCRICAO AS \"Tipo_serv_engenharia\",\n    UPPER(drg.drgrp_descricao) AS \"Destinacao_recurso_grupo\",\n    UPPER(dri.drids_descricao) AS \"Destinacao_recurso_iduso\",\n    e.drids_codigo AS \"Destinacao_recurso_codiog_iduso\",\n    e.destrec_codigo AS \"Codigo_destinacao_recurso\",\n    e.dresp_codigo AS \"Codigo_destinacao_recurso_especificacao\",\n    e.drgrp_codigo AS \"Destinacao_recurso_codigo_grupo\",\n    E.EMP_FUNDAMENTOCOMPRADIRETA AS \"Fund_compra_direta_codigo\",\n    FC.FCD_DESCRICAO AS \"Fund_compra_direta_descricao\",\n    UPPER(dr.destrec_descricao) AS \"Destinacao_recurso\",\n    o.org_nome AS \"Orgao\",\n    u.unor_nome AS \"UnidadeOrcamentaria\",\n    e.unor_codigo AS \"Unidade_orcamentaria_codigo\",\n    e.fn_codigo AS \"Funcao_codigo\",\n    f.fn_descricao AS \"Funcao_descricao\",\n    e.sfn_codigo AS \"Subfuncao_codigo\",\n    sf.sfn_descricao AS \"Subfuncao_descricao\",\n    e.prg_codigo AS \"Programa_codigo\",\n    e.prat_numero AS \"Num_Projeto_Atividade\",\n    e.ctec_codigo AS \"Categoria_Economica\",\n    e.ndesp_codigo AS \"Natureza_Despesa\",\n    e.mdap_codigo AS \"Modalidade_aplicacao_codigo\",\n    e.elde_codigo AS \"Elemento_despesa_codigo\",\n    el.elde_descricao AS \"Elemento_despesa_descricao\",\n    e.selde_codigo AS \"Subelemento_despesa_codigo\",\n    sub.selde_descricao AS \"Subelemento_despesa\",\n    e.emp_descricao AS \"Descricao\",\n    e.plic_numero AS \"Num_processo_licitatorio\",\n    e.cont_tipo AS \"Tipo_Contrato\",\n    e.cont_numaditivo AS \"Num_aditivo_contrato\",\n    e.conv_numero AS \"Num_convenio\",\n    e.conv_numaditivo AS \"Num_aditivo_convenio\",\n    decode(e.emp_compradiretaprocesso,'1','Não','2','Sim','N','Não','S','Sim') AS \"Compra_direta\",\n    decode(e.emp_tipo,'1','Estimativo','2','Global','3','Ordinário') AS \"Tipo\",\n    e.mesreferencia AS \"Mes_referencia\",\n    e.cg_identificacao AS \"Identificacao_credor\",\n    c.cg_tipopessoa AS \"Tipo_pesssoaCodigo\"\nFROM\n    aplic2008.empenho e\nLEFT JOIN aplic2008.cadastro_geral c ON (e.ent_codigo = c.ent_codigo AND c.exercicio >= 2015 AND e.cg_identificacao = c.cg_identificacao)\nLEFT JOIN aplic2008.empenho_obra eo ON (e.ent_codigo = eo.ent_codigo AND e.exercicio = eo.exercicio AND e.org_codigo = eo.org_codigo AND e.unor_codigo = eo.unor_codigo AND e.emp_numero = eo.emp_numero)\nLEFT JOIN aplic2008.orgao o ON (e.ent_codigo = o.ent_codigo AND e.exercicio = o.exercicio AND e.org_codigo = o.org_codigo)\nLEFT JOIN aplic2008.unidade_orcamentaria u ON (e.ent_codigo = u.ent_codigo AND e.exercicio = u.exercicio AND e.org_codigo = u.org_codigo AND e.unor_codigo = u.unor_codigo)\nLEFT JOIN aplic2008.projeto_atividade prj ON (e.ent_codigo = prj.ent_codigo AND e.exercicio = prj.exercicio AND e.prat_numero = prj.prat_numero AND e.prg_codigo = prj.prg_codigo)\nLEFT JOIN aplic2008.FONTE_RECURSO FRC ON (E.FREC_CODIGO = FRC.FREC_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO_IDUSO DRI ON (E.EXERCICIO = DRI.EXERCICIO AND E.DRIDS_CODIGO = DRI.DRIDS_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO_GRUPO DRG ON (E.EXERCICIO = DRG.EXERCICIO AND E.DRGRP_CODIGO = DRG.DRGRP_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO_ESPECIFIC DRE ON (E.EXERCICIO = DRE.EXERCICIO AND E.DRESP_CODIGO = DRE.DRESP_CODIGO)\nLEFT JOIN aplic2008.DESTINACAO_RECURSO DR ON (E.EXERCICIO = DR.EXERCICIO AND E.DESTREC_CODIGO = DR.DESTREC_CODIGO)\nLEFT JOIN aplic2008.elemento_despesa el ON (e.elde_codigo = el.elde_codigo AND e.exercicio = el.exercicio)\nLEFT JOIN aplic2008.subelemento_despesa sub ON (e.exercicio = sub.selde_exercicio AND e.elde_codigo = sub.elde_codigo AND e.selde_codigo = sub.selde_codigo)\nLEFT JOIN aplic2008.funcao f ON (e.fn_codigo = f.fn_codigo)\nLEFT JOIN aplic2008.subfuncao sf ON (e.sfn_codigo = sf.sfn_codigo)\nINNER JOIN vw_entidade_aplic v ON (e.ent_codigo = v.ent_codigo)\nLEFT JOIN aplic2008.tipo_despesa_rpps t ON (e.emp_tipodespesarpps = t.trpps_codigo)\nLEFT JOIN aplic2008.contrato_empenho ce ON (ce.ent_codigo = e.ent_codigo AND ce.exercicio = e.exercicio AND ce.org_codigo = e.org_codigo AND ce.unor_codigo = e.unor_codigo AND ce.emp_numero = e.emp_numero)\nLEFT JOIN aplic2008.TIPO_SERVICO_ENGENHARIA TSE ON (E.EMP_TIPOSERVICOENGENHARIA = TSE.TSENG_CODIGO)\nLEFT JOIN aplic2008.MODALIDADE_LICITACAO MLIC ON (E.MLIC_CODIGO = MLIC.MLIC_CODIGO)\n" +
		"LEFT JOIN aplic2008.FUNDAMENTO_COMPRA_DIRETA FC ON FC.FCD_CODIGO = E.EMP_FUNDAMENTOCOMPRADIRETA\n" +
		"WHERE\n   " +
		" e.ent_codigo = :1\n " +
		"   AND e.exercicio = :2\n  " +
		"  AND e.destrec_codigo = :3\n" +
		"ORDER BY\n    v.mun_nome,\n    v.ent_nome,\n    e.emp_numero,\n    e.emp_data"

	rows, err := repository.db.Query(query, unidadeGestoraCodigo, ano, codigoDestinacaoRecurso)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var filtroDetalheFontes []models.FiltroDetalheFonte
	for rows.Next() {
		var filtroDetalheFonte models.FiltroDetalheFonte
		if err := rows.Scan(
			&filtroDetalheFonte.Emp_Data,
			&filtroDetalheFonte.Emp_Numero,
			&filtroDetalheFonte.Credor,
			&filtroDetalheFonte.Valor_Empenhado,
			&filtroDetalheFonte.Valor_Liquidado,
			&filtroDetalheFonte.Valor_Pago,
			&filtroDetalheFonte.Valor_PagoRetencao,
			&filtroDetalheFonte.Anulado_Empenho,
			&filtroDetalheFonte.Qtde_Notas_Fiscais,
			&filtroDetalheFonte.Qtde_NFe,
			&filtroDetalheFonte.Contratos,
			&filtroDetalheFonte.Ent_codigo,
			&filtroDetalheFonte.Qtde_beneficiarios,
			&filtroDetalheFonte.Instrumento_contrato,
			&filtroDetalheFonte.Fonte_recurso_fonte,
			&filtroDetalheFonte.Tipo_concurso,
			&filtroDetalheFonte.Num_Concurso,
			&filtroDetalheFonte.Assistencia_social,
			&filtroDetalheFonte.Diarias,
			&filtroDetalheFonte.Relevante,
			&filtroDetalheFonte.Unidade_gestora,
			&filtroDetalheFonte.Municipio,
			&filtroDetalheFonte.Exercicio,
			&filtroDetalheFonte.Codigo_ug,
			&filtroDetalheFonte.Projeto_atividade,
			&filtroDetalheFonte.Dotacao,
			&filtroDetalheFonte.Num_obra,
			&filtroDetalheFonte.Tipo_serv_engenharia,
			&filtroDetalheFonte.Destinacao_recurso_grupo,
			&filtroDetalheFonte.Destinacao_recurso_iduso,
			&filtroDetalheFonte.Destinacao_recurso_codiog_iduso,
			&filtroDetalheFonte.Codigo_destinacao_recurso,
			&filtroDetalheFonte.Codigo_destinacao_recurso_especificacao,
			&filtroDetalheFonte.Destinacao_recurso_codigo_grupo,
			&filtroDetalheFonte.Fund_compra_direta_codigo,
			&filtroDetalheFonte.Fund_compra_direta_descricao,
			&filtroDetalheFonte.Destinacao_recurso,
			&filtroDetalheFonte.Orgao,
			&filtroDetalheFonte.UnidadeOrcamentaria,
			&filtroDetalheFonte.Unidade_orcamentaria_codigo,
			&filtroDetalheFonte.Funcao_codigo,
			&filtroDetalheFonte.Funcao_descricao,
			&filtroDetalheFonte.Subfuncao_codigo,
			&filtroDetalheFonte.Subfuncao_descricao,
			&filtroDetalheFonte.Programa_codigo,
			&filtroDetalheFonte.Num_Projeto_Atividade,
			&filtroDetalheFonte.Categoria_Economica,
			&filtroDetalheFonte.Natureza_Despesa,
			&filtroDetalheFonte.Modalidade_aplicacao_codigo,
			&filtroDetalheFonte.Elemento_despesa_codigo,
			&filtroDetalheFonte.Elemento_despesa_descricao,
			&filtroDetalheFonte.Subelemento_despesa_codigo,
			&filtroDetalheFonte.Subelemento_despesa,
			&filtroDetalheFonte.Descricao,
			&filtroDetalheFonte.Num_processo_licitatorio,
			&filtroDetalheFonte.Tipo_Contrato,
			&filtroDetalheFonte.Num_aditivo_contrato,
			&filtroDetalheFonte.Num_convenio,
			&filtroDetalheFonte.Num_aditivo_convenio,
			&filtroDetalheFonte.Compra_direta,
			&filtroDetalheFonte.Tipo,
			&filtroDetalheFonte.Mes_referencia,
			&filtroDetalheFonte.Identificacao_credor,
			&filtroDetalheFonte.Tipo_pesssoaCodigo,
		); err != nil {
			return nil, err
		}
		filtroDetalheFontes = append(filtroDetalheFontes, filtroDetalheFonte)
	}
	return filtroDetalheFontes, nil
}
