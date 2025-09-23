package queries

const Adiantamento = "select A.SOLIC_NUMPROCESSO as \"Num. Processo\",\n   " +
	"   A.PESS_MATRICULA as \"Matrícula\",\n       A.ADTO_TIPO as \"Código Tipo\",\n  " +
	"   DECODE(A.ADTO_TIPO, '1', '1 - Material de Cosumo', '2', '2 - Serviço') as \"Tipo\",\n " +
	"   A.ADTO_DATA as \"Data do Doc\",\n       A.ADTO_VALOR as \"Valor Adiantamento\",\n    " +
	"   A.ADTO_OBJETIVO as \"Objetivo\",\n       A.LEI_NUMERO as \"Nº Lei\",\n   " +
	"   A.ORG_CODIGO as \"Órgão\",\n       A.UNOR_CODIGO as \"Código Und.Orçamentaria\",\n " +
	"   A.EMP_NUMERO as \"Empenho\",\n       A.ADTO_DATALIMITE as \"Data Limite\",\n   " +
	"   A.DECR_NUMERO as \"Nº Decreto\",\n       P.PESS_NOME as \"Funcionário\",\n   " +
	"   O.ORG_NOME as \"Desc.Órgão\",\n       U.UNOR_NOME as \"Und.Orçamentaria\",\n  " +
	"   PC.PCAD_DATA as \"Data Prest.Contas\",\n    " +
	"   DECODE(PC.PCAD_APROVADA, 'S', 'Aprovado', 'N', 'Reprovado') as \"Situação Prest. Contas\",\n " +
	"   PC.PCAD_APROVADA,\n       (select count(1)\n          from aplic2008.ADIANTAMENTO AD\n   " +
	"   where AD.ENT_CODIGO = A.ENT_CODIGO\n           and AD.EXERCICIO = A.EXERCICIO\n   " +
	"   and AD.PESS_MATRICULA = A.PESS_MATRICULA\n      " +
	"   and AD.SOLIC_NUMPROCESSO not in (select PA.SOLIC_NUMPROCESSO\n  " +
	" from aplic2008.PRESTACAO_CONTAS_ADIANTAMENTO PA\n  " +
	"  where PA.ENT_CODIGO = AD.ENT_CODIGO\n  " +
	"and PA.EXERCICIO = AD.EXERCICIO\n  " +
	" and PA.SOLIC_NUMPROCESSO = AD.SOLIC_NUMPROCESSO\n " +
	" and PA.PESS_MATRICULA = AD.PESS_MATRICULA\n  " +
	" and PA.ADTO_TIPO = AD.ADTO_TIPO)) as \"Qtd.Ad.sem Prestação\",\n    " +
	" (select count(1)\n    from aplic2008.PRESTACAO_CONTAS_ADIANT_DOC PCA\n    " +
	"  where PCA.ENT_CODIGO = PC.ENT_CODIGO\n   and PCA.EXERCICIO = PC.EXERCICIO\n    " +
	"  and PCA.SOLIC_NUMPROCESSO = PC.SOLIC_NUMPROCESSO\n  " +
	"  and PCA.PESS_MATRICULA = PC.PESS_MATRICULA\n  and PCA.ADTO_TIPO = PC.ADTO_TIPO) as \"Prestação de Contas\",\n  " +
	"  (select count(1)\n     from aplic2008.DEVOLUCAO_ADIANTAMENTO DA\n     " +
	"  where DA.ENT_CODIGO = A.ENT_CODIGO\n  and DA.EXERCICIO = A.EXERCICIO\n    " +
	"  and DA.SOLIC_NUMPROCESSO = A.SOLIC_NUMPROCESSO\n  " +
	"  and DA.PESS_MATRICULA = A.PESS_MATRICULA\n  " +
	" and DA.ADTO_TIPO = A.ADTO_TIPO) as \"Ad. Devolução\"\n  from aplic2008.ADIANTAMENTO A\n " +
	" left join aplic2008.PESSOAL P on A.ENT_CODIGO = P.ENT_CODIGO\n and A.EXERCICIO = P.EXERCICIO\n" +
	" and A.PESS_MATRICULA = P.PESS_MATRICULA\n " +
	" left join aplic2008.PRESTACAO_CONTAS_ADIANTAMENTO PC on PC.ENT_CODIGO = A.ENT_CODIGO\n " +
	" and PC.EXERCICIO = A.EXERCICIO\n  " +
	" and PC.SOLIC_NUMPROCESSO = A.SOLIC_NUMPROCESSO\n  " +
	" and PC.PESS_MATRICULA = A.PESS_MATRICULA\n  " +
	" and PC.ADTO_TIPO = A.ADTO_TIPO\n  left join aplic2008.ORGAO O on A.ENT_CODIGO = O.ENT_CODIGO\n  " +
	" and A.EXERCICIO = O.EXERCICIO\n   and A.ORG_CODIGO = O.ORG_CODIGO\n" +
	"  left join aplic2008.UNIDADE_ORCAMENTARIA U on A.ENT_CODIGO = U.ENT_CODIGO\n " +
	" and A.EXERCICIO = U.EXERCICIO\n   and A.ORG_CODIGO = U.ORG_CODIGO\n  " +
	"and A.UNOR_CODIGO = U.UNOR_CODIGO\n" +
	" where A.ENT_CODIGO = :unidadeGestoraCodigo \n " +
	"and A.EXERCICIO = :ano \n" +
	"order by ADTO_DATA"
