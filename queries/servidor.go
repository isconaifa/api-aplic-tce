package queries

const ServidorQuery = "select distinct p.pess_matricula, p.pess_nome\n" +
	"from aplic2008.pessoal_folha_pagamento pf,\n" +
	"aplic2008.pessoal p\n" +
	"where pf.ent_codigo = p.ent_codigo\n" +
	"and ((pf.exercicio < '2015' and pf.exercicio = p.exercicio) or (pf.exercicio >= '2015' and p.exercicio >= '2015'))\n" +
	"and pf.pess_matricula = p.pess_matricula\n" +
	"and pf.ent_codigo = :unidadeGestoraCodigo\n" +
	"and pf.exercicio = :ano\n" +
	"order by p.pess_nome"
