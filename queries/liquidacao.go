package queries

const Liquidacao = "select L.LIQ_NUMERO as \"Nº da Liquidação\",\n  " +
	"     L.LIQ_DATA as \"Data\",\n       L.LIQ_VALOR as \"Valor\",\n  " +
	"     C.CG_NOME as \"Credor\",\n       L.ORG_CODIGO as \"Órgão(código)\",\n " +
	"      L.UNOR_CODIGO as \"Unidade Orçamentária(código)\",\n     " +
	"  L.EMP_NUMERO as \"Nº Empenho\",\n       (select count(1)\n    " +
	"      from aplic2008.ITEM_LIQUIDACAO_EMPENHO I\n     " +
	"    where I.ENT_CODIGO = L.ENT_CODIGO\n           and I.EXERCICIO = L.EXERCICIO\n    " +
	"       and I.ORG_CODIGO = L.ORG_CODIGO\n           and I.UNOR_CODIGO = L.UNOR_CODIGO\n    " +
	"       and I.EMP_NUMERO = L.EMP_NUMERO\n           and I.LIQ_NUMERO = L.LIQ_NUMERO) as \"Item(ns)\",\n   " +
	"    (select NVL(sum(DLP_VALOR), 0)\n          from aplic2008.DESCONTO_LIQUIDADO_PAGO DLP\n   " +
	"      where DLP.ENT_CODIGO = L.ENT_CODIGO\n           and DLP.EXERCICIO = L.EXERCICIO\n   " +
	"        and DLP.ORG_CODIGO = L.ORG_CODIGO\n           and DLP.UNOR_CODIGO = L.UNOR_CODIGO\n    " +
	"       and DLP.EMP_NUMERO = L.EMP_NUMERO\n           and DLP.LIQ_NUMERO = L.LIQ_NUMERO) as \"Desconto Liq.Pago\",\n  " +
	"     (select NVL(sum(DLIQ_VALOR), 0)\n          from aplic2008.DESCONTO_LIQUIDADO DL\n    " +
	"     where DL.ENT_CODIGO = L.ENT_CODIGO\n           and DL.EXERCICIO = L.EXERCICIO\n      " +
	"     and DL.ORG_CODIGO = L.ORG_CODIGO\n           and DL.UNOR_CODIGO = L.UNOR_CODIGO\n    " +
	"       and DL.EMP_NUMERO = L.EMP_NUMERO\n           and DL.LIQ_NUMERO = L.LIQ_NUMERO) as \"Desconto Liquidado\",\n" +
	"       (select count(1)\n          from aplic2008.NOTA_FISCAL N\n      " +
	"   where N.ENT_CODIGO = L.ENT_CODIGO\n           and N.EXERCICIO = L.EXERCICIO\n    " +
	"       and N.ORG_CODIGO = L.ORG_CODIGO\n           and N.UNOR_CODIGO = L.UNOR_CODIGO\n      " +
	"     and N.EMP_NUMERO = L.EMP_NUMERO\n           and N.LIQ_NUMERO = L.LIQ_NUMERO) as \"Nota Fiscal\",\n    " +
	"   (select count(1)\n          from aplic2008.RECIBO_PAGTO_EVENTUAL R\n    " +
	"     where R.ENT_CODIGO = L.ENT_CODIGO\n           and R.EXERCICIO = L.EXERCICIO\n     " +
	"      and R.ORG_CODIGO = L.ORG_CODIGO\n           and R.UNOR_CODIGO = L.UNOR_CODIGO\n   " +
	"        and R.EMP_NUMERO = L.EMP_NUMERO\n           and R.LIQ_NUMERO = L.LIQ_NUMERO) as \"Recibo(s)\",\n   " +
	"    (select NVL(sum(P.PGTO_VALOR), 0)\n          from aplic2008.PAGAMENTO_EMPENHO P,\n   " +
	"            aplic2008.PAGAMENTO_EMPENHO_LIQUIDACAO PL\n         where P.ENT_CODIGO = PL.ENT_CODIGO\n     " +
	"      and P.EXERCICIO = PL.EXERCICIO\n           and PL.PGTO_NUMERO = P.PGTO_NUMERO\n    " +
	"       and PL.ENT_CODIGO = L.ENT_CODIGO\n           and PL.EXERCICIO = L.EXERCICIO\n      " +
	"     and PL.ORG_CODIGO = L.ORG_CODIGO\n           and PL.UNOR_CODIGO = L.UNOR_CODIGO\n     " +
	"      and PL.EMP_NUMERO = L.EMP_NUMERO\n           and PL.LIQ_NUMERO = L.LIQ_NUMERO) -\n   " +
	"    (select nvl(sum(alp.anul_valor),0)\n         from aplic2008.anulacao_pagamento_empenho alp,\n     " +
	"         aplic2008.pagamento_empenho_liquidacao pel\n         where alp.ent_codigo = pel.ent_codigo\n   " +
	"        and alp.exercicio = pel.exercicio\n           and alp.pgto_numero = pel.pgto_numero\n      " +
	"     and pel.ent_codigo = l.ent_codigo\n           and pel.org_codigo = l.org_codigo\n     " +
	"      and pel.unor_codigo = l.unor_codigo\n           and pel.emp_numero = l.emp_numero\n    " +
	" and pel.liq_numero = l.liq_numero) as \"Valor Pago\",\n   " +
	" (select NVL(sum(A.ANUL_VALOR), 0)\n  " +
	" from aplic2008.ANULACAO_LIQUIDACAO_EMPENHO A\n   " +
	"  where A.ENT_CODIGO = L.ENT_CODIGO\n       " +
	"    and A.EXERCICIO = L.EXERCICIO\n       " +
	"    and A.ORG_CODIGO = L.ORG_CODIGO\n       " +
	"    and A.UNOR_CODIGO = L.UNOR_CODIGO\n       " +
	"    and A.EMP_NUMERO = L.EMP_NUMERO\n      " +
	"     and A.LIQ_NUMERO = L.LIQ_NUMERO) as \"Anulação Liquidação\",\n " +
	"      L.CONV_NUMERO as \"Convênio\",\n    " +
	"   L.CONV_NUMADITIVO as \"Convênio Aditivo\",\n    " +
	"   DECODE(L.LIQ_TIPODOCUMENTOHABIL, '1', 'NOTA FISCAL', '2', 'RECIBO',\n " +
	"      '3', 'COMPROVANTES DIVERSOS(Documentos Não Fiscais)', '4', 'NENHUM COMPROVANTE') as \"Tipo Documento Hábil\",\n  " +
	"     L.LIQ_DATAATESTO as \"Data do atesto\",\n   " +
	"    L.PESS_MATRICULA as \"Matr. do resp. pela liquidação\",\n   " +
	"    L.PESS_MATRICULARESPATESTO as \"Matr. do resp. pelo atesto\",\n  " +
	"     P1.PESS_NOME as \"Responsável pela liquidação\",\n  " +
	"     P2.PESS_NOME as \"Responsável pelo atesto\"\n " +
	" from aplic2008.LIQUIDACAO_EMPENHO L\nleft join aplic2008.EMPENHO E on E.ENT_CODIGO = L.ENT_CODIGO\n   " +
	"and E.EXERCICIO = L.EXERCICIO\n and E.EMP_NUMERO = L.EMP_NUMERO\n " +
	"and E.ORG_CODIGO = L.ORG_CODIGO\n and E.UNOR_CODIGO = L.UNOR_CODIGO\n" +
	"left join aplic2008.CADASTRO_GERAL C on E.ENT_CODIGO = C.ENT_CODIGO\n" +
	" and ((E.EXERCICIO < 2015 AND E.EXERCICIO = C.EXERCICIO)\n" +
	" OR (E.EXERCICIO >= 2015 and C.EXERCICIO >= 2015))\n" +
	" and E.CG_IDENTIFICACAO = C.CG_IDENTIFICACAO\n" +
	"left join aplic2008.PESSOAL P1 on L.ENT_CODIGO = P1.ENT_CODIGO\n " +
	"and ((L.EXERCICIO < 2015 AND L.EXERCICIO = P1.EXERCICIO)\n" +
	" OR (L.EXERCICIO >= 2015 and P1.EXERCICIO >= 2015))\n " +
	"and L.PESS_MATRICULA = P1.PESS_MATRICULA\nleft join aplic2008.PESSOAL P2 on L.ENT_CODIGO = P2.ENT_CODIGO\n" +
	" and ((L.EXERCICIO < 2015 AND L.EXERCICIO = P2.EXERCICIO)\n" +
	"OR (L.EXERCICIO >= 2015 and P2.EXERCICIO >= 2015))\n" +
	" and L.PESS_MATRICULARESPATESTO = P2.PESS_MATRICULA\n" +
	" where L.ENT_CODIGO = :1\n" +
	"   and L.EXERCICIO = :2\n" +
	" and L.ORG_CODIGO = :3\n" +
	" and L.UNOR_CODIGO =:4\n" +
	" and L.EMP_NUMERO = :5\n" +
	"order by L.EMP_NUMERO"
