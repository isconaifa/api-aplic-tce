package queries

const MesQuery = " SELECT comp.mes_referencia AS MesReferencia,\n" +
	"comp.descricao AS Descricao\n" +
	"FROM aplic2008.competencia_aplic comp\n" +
	"WHERE mes_referencia BETWEEN '01' AND '12'\n" +
	"ORDER BY mes_referencia"
