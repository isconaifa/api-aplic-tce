package queries

const Pagamento = "SELECT P.PGTO_NUMERO as \"Nº Pagamento\",\n   " +
	"P.PGTO_DATA as \"Data\",\n       P.PGTO_VALOR as \"Valor\",\n    " +
	"PL.ORG_CODIGO as \"Órgão(código)\",\n  " +
	"PL.UNOR_CODIGO as \"Unidade Orçamentária(código)\",\n  " +
	"PL.EMP_NUMERO as \"Nº Empenho\",\n    " +
	"PL.UNOR_CODIGO as \"Und.Orçcamentaria\",\n   " +
	"PL.ORG_CODIGO as \"Órgão\",\n       (SELECT COUNT(1)\n   " +
	"FROM aplic2008.PAGAMENTO_EMPENHO_DOCUMENTO PD\n     " +
	"WHERE PD.ENT_CODIGO = :unidadeGestoraCodigo\n" +
	"AND PD.EXERCICIO = :ano\n   " +
	"AND PD.PGTO_NUMERO = P.PGTO_NUMERO) as \"Qtde Documentos\",\n  " +
	"(SELECT NVL(SUM(A.ANUL_VALOR), 0)\n FROM aplic2008.ANULACAO_PAGAMENTO_EMPENHO A\n  " +
	"WHERE A.ENT_CODIGO = :unidadeGestoraCodigo\n " +
	"AND A.EXERCICIO = :ano\n     " +
	"AND A.PGTO_NUMERO = P.PGTO_NUMERO) as \"Anulação Pagamento\",\n  " +
	"P.CONV_NUMERO as \"Nº Convênio\",\n     " +
	"P.CONV_NUMADITIVO as \"Nº Convênio Adt\",\n   " +
	"PL.PGTO_JUSTIFICATIVAFORAORDEM as \"Jusitificativa\",\n  " +
	"PL.LIQ_NUMERO as \"Nº Liquidação\"\n  FROM aplic2008.PAGAMENTO_EMPENHO P\n" +
	"LEFT JOIN aplic2008.PAGAMENTO_EMPENHO_LIQUIDACAO PL\n " +
	"ON P.ENT_CODIGO = PL.ENT_CODIGO\n  " +
	"AND P.EXERCICIO = PL.EXERCICIO\n  " +
	"AND P.PGTO_NUMERO = PL.PGTO_NUMERO\n" +
	"WHERE PL.ENT_CODIGO = :unidadeGestoraCodigo\n " +
	"AND PL.EXERCICIO = :ano\n " +
	"AND PL.EMP_NUMERO = :numEmpenho\n" +
	"ORDER BY PL.EMP_NUMERO, P.PGTO_DATA"
