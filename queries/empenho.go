package queries

const EmpenhoTotal = "select distinct e.ent_codigo, e.org_codigo as \"Órgão(código)\",\n" +
	"e.unor_codigo as \"Unidade Orçamentária(código)\",    \n " +
	"e.emp_numero as \"N° do Empenho\",  \n" +
	"e.fn_codigo as \"Função(código)\",  \n " +
	"f.fn_descricao as \"Função(descrição)\",  \n  " +
	"e.sfn_codigo as \"SubFunção(código)\", \n " +
	"sf.sfn_descricao as \"SubFunção(descrição)\", \n " +
	" e.prg_codigo as \"Programa(código)\", \n  " +
	"e.prat_numero as \"N° do Projeto/Atividade\",  \n " +
	"e.ctec_codigo as \"Categoria Econômica\", \n " +
	"e.ndesp_codigo as \"Natureza da Despesa\", \n " +
	"e.mdap_codigo as \"Modalidade aplicação(código)\", \n" +
	"e.elde_codigo as \"Elemento de Despesa(código)\",\n" +
	"el.elde_descricao as \"Elemento de Despesa(descrição)\",\n" +
	"e.selde_codigo as \"Subelemento de Despesa(código)\", \n" +
	"sub.selde_descricao as \"Subelemento de Despesa\", \n" +
	"e.emp_data as \"Data\", \n e.emp_descricao as \"Descrição\", \n " +
	"e.plic_numero as \"Nº do processo licitatório\",\n " +
	"e.mlic_codigo as \"Cod. Modalidade Proc. licit\",\n " +
	"mlic.mlic_descricao as \"Modalidade proc. licitatório\"," +
	"\n e.cont_numero as \"Nº do contrato\", " +
	"\n e.cont_tipo as \"Tipo do contrato\",   \n       e.cont_numaditivo as \"Nº do aditivo do contrato\", e.conv_numero as \"Nº do convênio\", \n " +
	"e.conv_numaditivo as \"Nº do aditivo do convênio\", \n " +
	"decode(e.emp_compradiretaprocesso,'1','Não','2','Sim','N','Não','S','Sim')\n " +
	"as \"Compra direta?\",\n " +
	"decode(e.emp_tipo,'1','Estimativo','2','Global','3','Ordinário') as \"Tipo\",\n " +
	"e.emp_valor as \"Empenhado(sem anulação)\", \n" +
	"e.mesreferencia as \"Mês de referência\", \n " +
	"e.cg_identificacao as \"Identificação do credor\",\n" +
	"c.cg_nome as \"Credor\", \n" +
	"c.cg_tipopessoa as \"Tipo pessoa(código)\",\n" +
	"decode(c.cg_tipopessoa,'1','Física','2','Jurídica') as \"Tipo pessoa\",\n " +
	"o.org_nome as \"Órgão\", \n u.unor_nome as \"Unidade Orçamentária\", \n" +
	"E.EMP_FUNDAMENTOCOMPRADIRETA as \"Fund. Compra Direta Cód.\",\n" +
	"FC.FCD_DESCRICAO as \"Fund. Compra Direta Desc.\",\n" +
	"DECODE(C.CG_OPTANTESIMPLESNACIONAL, 'S', 'SIM', 'N', 'NÃO') as \"Optante simples?\",\n" +
	"(select count(1) from aplic2008.NOTA_FISCAL N  where N.ENT_CODIGO = E.ENT_CODIGO   and N.EXERCICIO = E.EXERCICIO and N.ORG_CODIGO = E.ORG_CODIGO  and N.UNOR_CODIGO = E.UNOR_CODIGO  and N.EMP_NUMERO = E.EMP_NUMERO) \"Qtde.Notas Fiscais\",  " +
	"(select count(1)  FROM aplic2008.nota_fiscal n WHERE n.ent_codigo = e.ent_codigo \n  " +
	"AND n.exercicio = e.exercicio  AND n.org_codigo = e.org_codigo  AND n.unor_codigo = e.unor_codigo  \n " +
	"AND n.emp_numero = e.emp_numero  AND n.ntfsc_numeronfe is not null)  \"Qtde.NF-e\", (select count(1) FROM aplic2008.contrato_empenho n \n  " +
	"WHERE n.ent_codigo = e.ent_codigo AND n.exercicio = e.exercicio AND n.org_codigo = e.org_codigo  \n  " +
	"AND n.unor_codigo = e.unor_codigo  AND n.emp_numero = e.emp_numero)  \"Contrato(s)\", \n" +
	" e.drids_codigo as \"Dest. Rec. Código Iduso\",\n e.drgrp_codigo as \"Dest. Rec. Código Grupo\",\n" +
	" e.dresp_codigo as \"Dest. Rec. Cód. Especificação\",\n e.destrec_codigo as \"Cód. Destinação Recurso\",\n" +
	" upper(dri.drids_descricao) as \"Dest. Rec. Iduso\",\n upper(drg.drgrp_descricao) as \"Dest. Rec. Grupo\",\n" +
	" upper(dre.dresp_descricao) as \"Dest. Rec. Especificação\",\n upper(dr.destrec_descricao) as \"Destinação de Recurso\",\n" +
	"TSE.TSENG_DESCRICAO as \"Tipo serv. engenharia\",\n  aplic2008.FN_EMPENHO_OBRA_PROJETO(E.ENT_CODIGO, E.EXERCICIO, E.ORG_CODIGO, E.UNOR_CODIGO, E.EMP_NUMERO) as \"Nº Obra\",\n " +
	"prj.prat_descricao as \"Projeto atividade\", e.ctec_codigo || '.' ||  e.ndesp_codigo || '.' || e.mdap_codigo || '.' ||       e.elde_codigo || '.' ||  e.selde_codigo as \"Dotação\",\n " +
	"e.ent_codigo as \"Código da UG\",  e.exercicio as \"Exercício\",  v.mun_nome as \"Município\",  v.ent_nome as \"Unidade Gestora\",   '   ' as \"Relevante\", \n " +
	"(select Count(1) from aplic2008.diaria dir  where dir.ent_codigo = e.ent_codigo  and dir.exercicio = e.exercicio and dir.org_codigo = e.org_codigo  and dir.unor_codigo = e.unor_codigo and dir.emp_numero = e.emp_numero) as Diarias ,FRC.FREC_DESCRICAO as \"Fonte de recurso - descrição\",\n " +
	"DECODE(e.emp_benefassistenciasocial, '1', 'Não existem beneficiários', \n" +
	"'2', 'Existem beneficiários sem cadastro informatizado', '3', 'Existem beneficiários com cadastro informatizado') as \"Assistência Social\",  \n" +
	"e.conc_numero as \"N° do Concurso\",  e.conc_tipo as \"Tipo do Concurso\", e.frec_codigo as \"Fonte de recurso - código\", decode(emp_instrumentocontrato,'1','Contrato','2','Carta-Contrato','3','Nota de Empenho da Despesa','4',\n " +
	"'Autorização de Compra','5','Ordem de Execução de Serviço','6', \n" +
	"'Outros Instrumentos Hábeis') as \"Instrumento contrato\", t.trpps_descricao as \"Tipo da despesa(RPPS)\", (select Count(1) from aplic2008.BENEF_ASSISTENCIA_SOCIAL_EMP bas \n " +
	"where bas.ENT_CODIGO = E.ENT_CODIGO and bas.EXERCICIO = E.EXERCICIO and bas.ORG_CODIGO = E.ORG_CODIGO and bas.UNOR_CODIGO = E.UNOR_CODIGO \n" +
	"and bas.EMP_NUMERO = E.EMP_NUMERO) as \"Qtde Beneficiários\"  ,e.emp_valor - (select nvl(sum(a.anul_valor),0)\n" +
	"from aplic2008.anulacao_empenho a where a.ent_codigo = e.ent_codigo  and a.exercicio = e.exercicio and a.org_codigo = e.org_codigo  \n" +
	"and a.unor_codigo = e.unor_codigo and a.emp_numero = e.emp_numero) as \"Valor Empenhado\", (select nvl(sum(a.anul_valor),0) from aplic2008.anulacao_empenho a\n" +
	"where a.ent_codigo = e.ent_codigo  and a.exercicio = e.exercicio  and a.org_codigo = e.org_codigo  and a.unor_codigo = e.unor_codigo \n" +
	"and a.emp_numero = e.emp_numero ) as \"Anulado Empenho\",(select nvl(sum(al.anul_valor),0) from aplic2008.anulacao_liquidacao_empenho al\n" +
	" where al.ent_codigo = e.ent_codigo and al.exercicio = e.exercicio and al.org_codigo = e.org_codigo and al.unor_codigo = e.unor_codigo \n " +
	"and al.emp_numero = e.emp_numero) as \"Anulado Liquidação\", (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp,  \n" +
	"aplic2008.pagamento_empenho_liquidacao pel where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio  and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo  and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo\n" +
	"and pel.emp_numero = e.emp_numero) as \"Anulado Pagamento\",(select coalesce(sum(DLIQ_VALOR),0)  from aplic2008.desconto_liquidado d \n " +
	"where d.ent_codigo = e.ent_codigo  and d.exercicio = e.exercicio and d.org_codigo = e.org_codigo  and d.unor_codigo = e.unor_codigo\n" +
	"and d.emp_numero = e.emp_numero) as \"Valor Retido(Liquidação)\" ,(select nvl(sum(l.liq_valor),0) from aplic2008.liquidacao_empenho l \n" +
	"where l.ENT_CODIGO = E.ENT_CODIGO and l.EXERCICIO = E.EXERCICIO and l.ORG_CODIGO = E.ORG_CODIGO  and l.UNOR_CODIGO = E.UNOR_CODIGO and l.EMP_NUMERO = E.EMP_NUMERO ) - (select nvl(sum(al.anul_valor),0) from aplic2008.anulacao_liquidacao_empenho al\n" +
	"where al.ent_codigo = e.ent_codigo and al.exercicio = e.exercicio and al.org_codigo = e.org_codigo and al.unor_codigo = e.unor_codigo\n" +
	"and al.emp_numero = e.emp_numero ) as \"Valor Liquidado\" ,(select nvl(sum(l.liq_valor),0) from aplic2008.liquidacao_empenho l where l.ENT_CODIGO = E.ENT_CODIGO and l.EXERCICIO = E.EXERCICIO and l.ORG_CODIGO = E.ORG_CODIGO\n" +
	"and l.UNOR_CODIGO = E.UNOR_CODIGO and l.EMP_NUMERO = E.EMP_NUMERO )  as \"Valor Liquidado(sem anulação)\"  ,(select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p, \n" +
	"aplic2008.pagamento_empenho_liquidacao pl  where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO\n" +
	"and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO ) as \"Valor Pago(sem anulação)\" \n ,(select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p,\n" +
	"aplic2008.pagamento_empenho_liquidacao pl  where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO\n" +
	"and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO  and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO) - (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp,  \n" +
	" aplic2008.pagamento_empenho_liquidacao pel where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio  and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo  and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo \n" +
	"and pel.emp_numero = e.emp_numero ) \"Valor Pago\"\n ,(select nvl(sum(p.pgto_valor),0) from aplic2008.pagamento_empenho p,  aplic2008.pagamento_empenho_liquidacao pl  where P.ENT_CODIGO = PL.ENT_CODIGO and P.EXERCICIO = PL.EXERCICIO and P.PGTO_NUMERO = PL.PGTO_NUMERO\n      and PL.ENT_CODIGO = E.ENT_CODIGO and PL.EXERCICIO = E.EXERCICIO and PL.ORG_CODIGO = E.ORG_CODIGO  \n" +
	"and PL.UNOR_CODIGO = E.UNOR_CODIGO and PL.EMP_NUMERO = E.EMP_NUMERO ) - (select nvl(sum(alp.anul_valor),0) from aplic2008.anulacao_pagamento_empenho alp,  aplic2008.pagamento_empenho_liquidacao pel \n" +
	"where alp.ent_codigo = pel.ent_codigo and alp.exercicio = pel.exercicio  and alp.pgto_numero = pel.pgto_numero and pel.ent_codigo = e.ent_codigo and pel.org_codigo = e.org_codigo and pel.unor_codigo = e.unor_codigo \n" +
	"and pel.emp_numero = e.emp_numero ) + \n (select coalesce(sum(DLIQ_VALOR),0)  from aplic2008.desconto_liquidado d  where d.ent_codigo = e.ent_codigo and d.exercicio = e.exercicio  and d.org_codigo = e.org_codigo \n" +
	"and d.unor_codigo = e.unor_codigo and d.emp_numero = e.emp_numero) - \n" +
	" (select nvl(sum(aedl.aedliq_valor),0) from aplic2008.anulacao_estorno_desc_liquidad aedl  where aedl.ent_codigo = e.ent_codigo and aedl.exercicio = e.exercicio \n " +
	"and aedl.org_codigo = e.org_codigo and aedl.unor_codigo = e.unor_codigo  and aedl.emp_numero = e.emp_numero ) as \"Valor Pago+Retenções\" from aplic2008.empenho e \n" +
	"LEFT JOIN aplic2008.cadastro_geral c On (e.ent_codigo = c.ent_codigo and  c.exercicio >= 2015 and  e.cg_identificacao = c.cg_identificacao) \n" +
	"LEFT JOIN aplic2008.empenho_obra eo ON  ((e.ent_codigo = eo.ent_codigo) and  (e.exercicio = eo.exercicio) and (e.org_codigo = eo.org_codigo) and \n" +
	"(e.unor_codigo = eo.unor_codigo) and (e.emp_numero = eo.emp_numero)) LEFT JOIN aplic2008.orgao o ON (e.ent_codigo = o.ent_codigo)  \n" +
	" and (e.exercicio = o.exercicio) and (e.org_codigo = o.org_codigo) LEFT JOIN  aplic2008.unidade_orcamentaria u ON (e.ent_codigo = u.ent_codigo) \n " +
	"and (e.exercicio = u.exercicio)  and (e.org_codigo = u.org_codigo)  \n " +
	"and (e.unor_codigo = u.unor_codigo) LEFT JOIN aplic2008.projeto_atividade prj on (e.ent_codigo = prj.ent_codigo)  and (e.exercicio = prj.exercicio) \n" +
	"and (e.prat_numero = prj.prat_numero)  and (e.prg_codigo = prj.prg_codigo) left join aplic2008.FONTE_RECURSO FRC ON (E.FREC_CODIGO = FRC.FREC_CODIGO) \n" +
	"left join aplic2008.DESTINACAO_RECURSO_IDUSO DRI on (E.EXERCICIO = DRI.EXERCICIO and E.DRIDS_CODIGO = DRI.DRIDS_CODIGO)\n " +
	"left join aplic2008.DESTINACAO_RECURSO_GRUPO DRG on (E.EXERCICIO = DRG.EXERCICIO and E.DRGRP_CODIGO = DRG.DRGRP_CODIGO)\n " +
	"left join aplic2008.DESTINACAO_RECURSO_ESPECIFIC DRE on (E.EXERCICIO = DRE.EXERCICIO and E.DRESP_CODIGO = DRE.DRESP_CODIGO)\n " +
	"left join aplic2008.DESTINACAO_RECURSO DR on (E.EXERCICIO = DR.EXERCICIO and E.DESTREC_CODIGO = DR.DESTREC_CODIGO) LEFT JOIN aplic2008.elemento_despesa el on \n" +
	"(e.elde_codigo = el.elde_codigo) and (e.exercicio = el.exercicio) LEFT JOIN aplic2008.subelemento_despesa sub on  (e.exercicio = sub.selde_exercicio) \n" +
	"and (e.elde_codigo = sub.elde_codigo) and (e.selde_codigo = sub.selde_codigo) LEFT JOIN aplic2008.funcao f on  (e.fn_codigo = f.fn_codigo) \n" +
	"LEFT JOIN aplic2008.subfuncao sf on (e.sfn_codigo = sf.sfn_codigo)  INNER JOIN vw_entidade_aplic v on (e.ent_codigo = v.ent_codigo)  \n" +
	"LEFT JOIN aplic2008.tipo_despesa_rpps t on (e.emp_tipodespesarpps = t.trpps_codigo) LEFT JOIN aplic2008.contrato_empenho ce \n" +
	"on (ce.ent_codigo = e.ent_codigo)  and (ce.exercicio = e.exercicio) and (ce.org_codigo = e.org_codigo) and (ce.unor_codigo = e.unor_codigo) \n" +
	"and (ce.emp_numero = e.emp_numero) left join aplic2008.TIPO_SERVICO_ENGENHARIA TSE ON (E.EMP_TIPOSERVICOENGENHARIA = TSE.TSENG_CODIGO) \n" +
	"left join aplic2008.MODALIDADE_LICITACAO MLIC ON (E.MLIC_CODIGO = MLIC.MLIC_CODIGO) left join aplic2008.FUNDAMENTO_COMPRA_DIRETA FC on FC.FCD_CODIGO = E.EMP_FUNDAMENTOCOMPRADIRETA \n" +
	"where 1 = 1  and e.ent_codigo = :unidadeGestoraCodigo\n" +
	"and e.exercicio = :ano\n" +
	" Order by v.mun_nome, v.ent_nome, e.emp_numero, e.emp_data"
