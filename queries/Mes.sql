SELECT comp.mes_referencia AS MesReferencia,
comp.descricao AS Descricao
FROM aplic2008.competencia_aplic comp
WHERE mes_referencia BETWEEN '01' AND '12'
ORDER BY mes_referencia