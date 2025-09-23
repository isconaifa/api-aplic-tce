package queries

const FolhaPagamento = "SELECT p.fpgto_anoreferencia AS \"Ano de referência\",\n " +
	"p.fpgto_mesreferencia AS \"Cód. mês de referência\",\n " +
	"m.am_descricao AS \"Mês de referência\",\n " +
	"coalesce(SUM(p.fpgto_valorbase), 0) AS \"Valor bruto\",\n  " +
	"coalesce(SUM(p.fpgto_valorgratificacoes), 0) AS \"Gratificação(ões)\",\n" +
	"coalesce(SUM(p.fpgto_valorbeneficios), 0) AS \"Benefício(s)\",\n " +
	"coalesce(SUM(p.fpgto_valordescontos), 0) AS \"Desconto(s)\",\n  " +
	"coalesce(SUM(p.fpgto_valorliquido), 0) AS \"Valor líquido\",\n" +
	"(SELECT COUNT(DISTINCT pes.pess_cpf)\n " +
	"FROM aplic2008.pessoal_folha_pagamento a\n " +
	"INNER JOIN aplic2008.pessoal pes\n   " +
	"ON a.ent_codigo = pes.ent_codigo\n  " +
	"AND pes.exercicio >= 2015\n   AND a.pess_matricula = pes.pess_matricula\n" +
	"WHERE a.ent_codigo = p.ent_codigo\n  " +
	"AND a.exercicio = p.exercicio\n " +
	"AND a.fpgto_anoreferencia = p.fpgto_anoreferencia\n " +
	"AND a.fpgto_mesreferencia = p.fpgto_mesreferencia) AS \"Qtde. funcionários\"\n" +
	"FROM aplic2008.folha_pagamento p\n INNER JOIN publico.meses m\n" +
	"ON (p.fpgto_mesreferencia = m.am_identificacao)\n" +
	"WHERE p.ent_codigo = :unidadeGestoraCodigo\n " +
	"AND p.exercicio = :ano\n" +
	"GROUP BY p.ent_codigo, p.exercicio,\n     " +
	"p.fpgto_anoreferencia,\n  " +
	"p.fpgto_mesreferencia,\n   " +
	"m.am_descricao\nORDER BY p.fpgto_anoreferencia,\n  " +
	"p.fpgto_mesreferencia"
