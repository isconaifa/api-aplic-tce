package queries

const Pagamento = "select P.PGTO_NUMERO as \"Nº Pagamento\",\n  " +
	"     P.PGTO_DATA as \"Data\",\n   " +
	"    P.PGTO_VALOR as \"Valor\",\n    " +
	"   PL.ORG_CODIGO as \"Órgão(código)\",\n     " +
	"  PL.UNOR_CODIGO as \"Unidade Orçamentária(código)\",\n   " +
	"    PL.EMP_NUMERO as \"Nº Empenho\",\n    " +
	"   PL.UNOR_CODIGO as \"Und.Orçcamentaria\",\n  " +
	"     PL.ORG_CODIGO as \"Órgão\",\n    " +
	"   (select count(1)\n          from aplic2008.PAGAMENTO_EMPENHO_DOCUMENTO PD\n  " +
	"       where PD.ENT_CODIGO = :1\n       " +
	"    and PD.EXERCICIO = :2\n       " +
	"    and PD.PGTO_NUMERO = P.PGTO_NUMERO) as \"Qtde Documentos\",\n   " +
	"    (select NVL(sum(A.ANUL_VALOR), 0)\n     " +
	"     from aplic2008.ANULACAO_PAGAMENTO_EMPENHO A\n    " +
	"     where A.ENT_CODIGO = :1\n         " +
	"  and A.EXERCICIO = :2\n      " +
	"     and A.PGTO_NUMERO = P.PGTO_NUMERO) as \"Anulação Pagamento\",\n    " +
	"   P.CONV_NUMERO as \"Nº Convênio\",\n     " +
	"  P.CONV_NUMADITIVO as \"Nº Convênio Adt\",\n    " +
	"   PL.PGTO_JUSTIFICATIVAFORAORDEM as \"Jusitificativa\",\n    " +
	"   PL.LIQ_NUMERO as \"Nº Liquidação\"\n " +
	" from aplic2008.PAGAMENTO_EMPENHO P\n " +
	" left join aplic2008.PAGAMENTO_EMPENHO_LIQUIDACAO PL on P.ENT_CODIGO = PL.ENT_CODIGO\n                                                     and P.EXERCICIO = PL.EXERCICIO\n                                                     and P.PGTO_NUMERO = PL.PGTO_NUMERO\n" +
	" where PL.ENT_CODIGO = :1\n " +
	"  and PL.EXERCICIO = :2\n " +
	"  and PL.EMP_NUMERO = :3\n" +
	" order by PL.EMP_NUMERO, P.PGTO_DATA"
