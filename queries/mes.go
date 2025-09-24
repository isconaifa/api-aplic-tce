package queries

const MesQuery = "select comp.mes_referencia, comp.descricao\n" +
	"from aplic2008.competencia_aplic comp\n" +
	"where mes_referencia between '01' and '12'\n" +
	"order by mes_referencia"
