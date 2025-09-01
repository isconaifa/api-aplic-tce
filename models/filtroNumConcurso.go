package models

type FiltroNumConcurso struct {
	EntCodigo                 *string  `json:"ent_codigo"`                   // Corresponde a: e.ent_codigo
	CodigoUG                  *string  `json:"codigo_ug"`                    // Corresponde a: e.ent_codigo as "Código da UG"
	Exercicio                 *string  `json:"exercicio"`                    // Corresponde a: e.exercicio
	EmpNumero                 *string  `json:"empenho_numero"`               // Corresponde a: e.emp_numero
	EmpData                   *string  `json:"empenho_data"`                 // Corresponde a: e.emp_data
	Tipo                      *string  `json:"tipo"`                         // Corresponde a: decode(e.emp_tipo,...)
	Descricao                 *string  `json:"descricao"`                    // Corresponde a: e.emp_descricao
	MesReferencia             *string  `json:"mes_referencia"`               // Corresponde a: e.mesreferencia
	ValorEmpenhado            *float64 `json:"valor_empenhado"`              // Corresponde a: "Valor Empenhado"
	ValorLiquidado            *float64 `json:"valor_liquidado"`              // Corresponde a: "Valor Liquidado"
	ValorPago                 *float64 `json:"valor_pago"`                   // Corresponde a: "Valor Pago"
	AnuladoEmpenho            *float64 `json:"anulado_empenho"`              // Corresponde a: "Anulado Empenho"
	AnuladoLiquidacao         *float64 `json:"anulado_liquidacao"`           // Corresponde a: "Anulado Liquidação"
	AnuladoPagamento          *float64 `json:"anulado_pagamento"`            // Corresponde a: "Anulado Pagamento"
	ValorRetidoLiquidacao     *float64 `json:"valor_retido_liquidacao"`      // Corresponde a: "Valor Retido(Liquidação)"
	ValorPagoMaisRetencoes    *float64 `json:"valor_pago_mais_retencoes"`    // Corresponde a: "Valor Pago+Retenções"
	EmpenhadoSemAnulacao      *float64 `json:"empenhado_sem_anulacao"`       // Corresponde a: "Empenhado(sem anulação)"
	ValorLiquidadoSemAnulacao *float64 `json:"valor_liquidado_sem_anulacao"` // Corresponde a: "Valor Liquidado(sem anulação)"
	ValorPagoSemAnulacao      *float64 `json:"valor_pago_sem_anulacao"`      // Corresponde a: "Valor Pago(sem anulação)"
	Municipio                 *string  `json:"municipio"`                    // Corresponde a: v.mun_nome
	UnidadeGestora            *string  `json:"unidade_gestora"`              // Corresponde a: v.ent_nome
	Orgao                     *string  `json:"orgao"`                        // Corresponde a: o.org_nome
	UnidadeOrcamentaria       *string  `json:"unidade_orcamentaria"`         // Corresponde a: u.unor_nome
	Credor                    *string  `json:"credor"`                       // Corresponde a: c.cg_nome
	IdentificacaoCredor       *string  `json:"identificacao_credor"`         // Corresponde a: e.cg_identificacao
	TipoPessoaCodigo          *string  `json:"tipo_pessoa_codigo"`           // Corresponde a: c.cg_tipopessoa
	TipoPessoa                *string  `json:"tipo_pessoa"`                  // Corresponde a: decode(c.cg_tipopessoa,...)
	OptanteSimples            *string  `json:"optante_simples"`              // Corresponde a: DECODE(C.CG_OPTANTESIMPLESNACIONAL,...)
	Dotacao                   string   `json:"dotacao"`                      // Corresponde a: e.ctec_codigo || ...
	OrgaoCodigo               *string  `json:"orgao_codigo"`                 // Corresponde a: e.org_codigo
	UnidadeOrcamentariaCodigo *string  `json:"unidade_orcamentaria_codigo"`  // Corresponde a: e.unor_codigo
	FuncaoCodigo              *string  `json:"funcao_codigo"`                // Corresponde a: e.fn_codigo
	FuncaoDescricao           *string  `json:"funcao_descricao"`             // Corresponde a: f.fn_descricao
	SubFuncaoCodigo           *string  `json:"subfuncao_codigo"`             // Corresponde a: e.sfn_codigo
	SubFuncaoDescricao        *string  `json:"subfuncao_descricao"`          // Corresponde a: sf.sfn_descricao
	ProgramaCodigo            *string  `json:"programa_codigo"`              // Corresponde a: e.prg_codigo
	ProjetoAtividade          *string  `json:"projeto_atividade"`            // Corresponde a: prj.prat_descricao
	NumProjetoAtividade       *string  `json:"num_projeto_atividade"`        // Corresponde a: e.prat_numero
	CategoriaEconomica        *string  `json:"categoria_economica"`          // Corresponde a: e.ctec_codigo
	NaturezaDespesa           *string  `json:"natureza_despesa"`             // Corresponde a: e.ndesp_codigo
	ModalidadeAplicacaoCodigo *string  `json:"modalidade_aplicacao_codigo"`  // Corresponde a: e.mdap_codigo
	ElementoDespesaCodigo     *string  `json:"elemento_despesa_codigo"`      // Corresponde a: e.elde_codigo
	ElementoDespesaDescricao  *string  `json:"elemento_despesa_descricao"`   // Corresponde a: el.elde_descricao
	SubelementoDespesaCodigo  *string  `json:"subelemento_despesa_codigo"`   // Corresponde a: e.selde_codigo
	SubelementoDespesa        *string  `json:"subelemento_despesa"`          // Corresponde a: sub.selde_descricao
	FonteRecursoCodigo        *string  `json:"fonte_recurso_codigo"`         // Corresponde a: e.frec_codigo
	FonteRecursoDescricao     *string  `json:"fonte_recurso_descricao"`      // Corresponde a: FRC.FREC_DESCRICAO
	CodDestinacaoRecurso      *string  `json:"cod_destinacao_recurso"`       // Corresponde a: e.destrec_codigo
	DestinacaoRecurso         *string  `json:"destinacao_recurso"`           // Corresponde a: upper(dr.destrec_descricao)
	DestRecCodigoIduso        *string  `json:"dest_rec_codigo_iduso"`        // Corresponde a: e.drids_codigo
	DestRecIduso              *string  `json:"dest_rec_iduso"`               // Corresponde a: upper(dri.drids_descricao)
	DestRecCodigoGrupo        *string  `json:"dest_rec_codigo_grupo"`        // Corresponde a: e.drgrp_codigo
	DestRecGrupo              *string  `json:"dest_rec_grupo"`               // Corresponde a: upper(drg.drgrp_descricao)
	DestRecCodEspecificacao   *string  `json:"dest_rec_cod_especificacao"`   // Corresponde a: e.dresp_codigo
	DestRecEspecificacao      *string  `json:"dest_rec_especificacao"`       // Corresponde a: upper(dre.dresp_descricao)
	NumProcessoLicitatorio    *string  `json:"num_processo_licitatorio"`     // Corresponde a: e.plic_numero
	ModalidadeProcLicitCodigo *string  `json:"modalidade_proc_licit_codigo"` // Corresponde a: e.mlic_codigo
	ModalidadeProcLicitatorio *string  `json:"modalidade_proc_licitatorio"`  // Corresponde a: mlic.mlic_descricao
	CompraDireta              *string  `json:"compra_direta"`                // Corresponde a: decode(e.emp_compradiretaprocesso,...)
	FundCompraDiretaCodigo    *string  `json:"fund_compra_direta_codigo"`    // Corresponde a: E.EMP_FUNDAMENTOCOMPRADIRETA
	FundCompraDiretaDescricao *string  `json:"fund_compra_direta_descricao"` // Corresponde a: FC.FCD_DESCRICAO
	InstrumentoContrato       *string  `json:"instrumento_contrato"`         // Corresponde a: decode(emp_instrumentocontrato,...)
	NumContrato               *string  `json:"num_contrato"`                 // Corresponde a: e.cont_numero
	TipoContrato              *string  `json:"tipo_contrato"`                // Corresponde a: e.cont_tipo
	NumAditivoContrato        *string  `json:"num_aditivo_contrato"`         // Corresponde a: e.cont_numaditivo
	NumConvenio               *string  `json:"num_convenio"`                 // Corresponde a: e.conv_numero
	NumAditivoConvenio        *string  `json:"num_aditivo_convenio"`         // Corresponde a: e.conv_numaditivo
	QtdeNotasFiscais          *int     `json:"qtde_notas_fiscais"`           // Corresponde a: "Qtde.Notas Fiscais"
	QtdeNFe                   *int     `json:"qtde_nfe"`                     // Corresponde a: "Qtde.NF-e"
	Contratos                 *int     `json:"contratos"`                    // Corresponde a: "Contrato(s)"
	Diarias                   *int     `json:"diarias"`                      // Corresponde a: "Diarias"
	QtdeBeneficiarios         *int     `json:"qtde_beneficiarios"`           // Corresponde a: "Qtde Beneficiários"
	Relevante                 *string  `json:"relevante"`                    // Corresponde a: ' ' as "Relevante"
	AssistenciaSocial         *string  `json:"assistencia_social"`           // Corresponde a: DECODE(e.emp_benefassistenciasocial,...)
	NumConcurso               *string  `json:"num_concurso"`                 // Corresponde a: e.conc_numero
	TipoConcurso              *string  `json:"tipo_concurso"`                // Corresponde a: e.conc_tipo
	TipoDespesaRPPS           *string  `json:"tipo_despesa_rpps"`            // Corresponde a: t.trpps_descricao
	NumObra                   *string  `json:"num_obra"`                     // Corresponde a: FN_EMPENHO_OBRA_PROJETO(...)
	TipoServEngenharia        *string  `json:"tipo_serv_engenharia"`         // Corresponde a: TSE.TSENG_DESCRICAO
}
