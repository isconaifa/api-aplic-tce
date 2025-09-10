package models

type Empenho struct {
	EntCodigo               string   `db:"ent_codigo"`
	OrgaoCodigo             string   `db:"Órgão(código)"`
	UnidadeCodigo           string   `db:"Unidade Orçamentária(código)"`
	NumeroEmpenho           string   `db:"N° do Empenho"`
	FuncaoCodigo            string   `db:"Função(código)"`
	FuncaoDescricao         string   `db:"Função(descrição)"`
	SubfuncaoCodigo         string   `db:"SubFunção(código)"`
	SubfuncaoDescricao      string   `db:"SubFunção(descrição)"`
	ProgramaCodigo          *string  `db:"Programa(código)"`
	ProjetoAtividadeNum     *string  `db:"N° do Projeto/Atividade"`
	CategoriaEconomica      *string  `db:"Categoria Econômica"`
	NaturezaDespesa         *string  `db:"Natureza da Despesa"`
	ModalidadeAplicacao     *string  `db:"Modalidade aplicação(código)"`
	ElementoDespesaCod      *string  `db:"Elemento de Despesa(código)"`
	ElementoDespesaDesc     *string  `db:"Elemento de Despesa(descrição)"`
	SubElementoDespCod      *string  `db:"Subelemento de Despesa(código)"`
	SubElementoDespDesc     *string  `db:"Subelemento de Despesa"`
	Data                    *string  `db:"Data"`
	Descricao               *string  `db:"Descrição"`
	ProcessoLicitatorioNum  *string  `db:"Nº do processo licitatório"`
	ModalidadeProcLicitCod  *string  `db:"Cod. Modalidade Proc. licit"`
	ModalidadeProcLicitDesc *string  `db:"Modalidade proc. licitatório"`
	ContratoNum             *string  `db:"Nº do contrato"`
	ContratoTipo            *string  `db:"Tipo do contrato"`
	ContratoNumAditivo      *string  `db:"Nº do aditivo do contrato"`
	ConvenioNum             *string  `db:"Nº do convênio"`
	ConvenioNumAditivo      *string  `db:"Nº do aditivo do convênio"`
	CompraDireta            *string  `db:"Compra direta?"`
	Tipo                    *string  `db:"Tipo"`
	ValorEmpenhadoBruto     *float64 `db:"Empenhado(sem anulação)"`
	MesReferencia           *string  `db:"Mês de referência"`
	CredorIdentificacao     *string  `db:"Identificação do credor"`
	CredorNome              string   `db:"Credor"`
	CredorTipoPessoaCod     string   `db:"Tipo pessoa(código)"`
	CredorTipoPessoa        *string  `db:"Tipo pessoa"`
	OrgaoNome               string   `db:"Órgão"`
	UnidadeNome             string   `db:"Unidade Orçamentária"`
	FundCompraDiretaCod     *string  `db:"Fund. Compra Direta Cód."`
	FundCompraDiretaDesc    *string  `db:"Fund. Compra Direta Desc."`
	OptanteSimples          *string  `db:"Optante simples?"`
	QtdeNotasFiscais        int      `db:"Qtde.Notas Fiscais"`
	QtdeNFe                 int      `db:"Qtde.NF-e"`
	QtdeContratos           int      `db:"Contrato(s)"`
	DestRecIdusoCod         string   `db:"Dest. Rec. Código Iduso"`
	DestRecGrupoCod         string   `db:"Dest. Rec. Código Grupo"`
	DestRecEspecificCod     string   `db:"Dest. Rec. Cód. Especificação"`
	DestRecCodigo           string   `db:"Cód. Destinação Recurso"`
	DestRecIdusoDesc        string   `db:"Dest. Rec. Iduso"`
	DestRecGrupoDesc        string   `db:"Dest. Rec. Grupo"`
	DestRecEspecificDesc    string   `db:"Dest. Rec. Especificação"`
	DestRecDesc             string   `db:"Destinação de Recurso"`
	TipoServicoEngenharia   *string  `db:"Tipo serv. engenharia"`
	NumeroObra              *string  `db:"Nº Obra"`
	ProjetoAtividadeDesc    string   `db:"Projeto atividade"`
	Dotacao                 string   `db:"Dotação"`
	CodigoUG                string   `db:"Código da UG"`
	Exercicio               int      `db:"Exercício"`
	Municipio               string   `db:"Município"`
	UnidadeGestora          string   `db:"Unidade Gestora"`
	Relevante               *string  `db:"Relevante"`
	QtdeDiarias             *int     `db:"Diarias"`
	FonteRecursoDesc        *string  `db:"Fonte de recurso - descrição"`
	AssistenciaSocial       *string  `db:"Assistência Social"`
	ConcursoNumero          *string  `db:"N° do Concurso"`
	ConcursoTipo            *string  `db:"Tipo do Concurso"`
	FonteRecursoCodigo      *string  `db:"Fonte de recurso - código"`
	InstrumentoContrato     *string  `db:"Instrumento contrato"`
	TipoDespesaRPPS         *string  `db:"Tipo da despesa(RPPS)"`
	QtdeBeneficiarios       *int     `db:"Qtde Beneficiários"`
	ValorEmpenhado          float64  `db:"Valor Empenhado"`
	AnuladoEmpenho          *float64 `db:"Anulado Empenho"`
	AnuladoLiquidacao       float64  `db:"Anulado Liquidação"`
	AnuladoPagamento        float64  `db:"Anulado Pagamento"`
	ValorRetido             float64  `db:"Valor Retido(Liquidação)"`
	ValorLiquidado          float64  `db:"Valor Liquidado"`
	ValorLiquidadoBruto     float64  `db:"Valor Liquidado(sem anulação)"`
	ValorPagoBruto          float64  `db:"Valor Pago(sem anulação)"`
	ValorPago               float64  `db:"Valor Pago"`
	ValorPagoRetencoes      *float64 `db:"Valor Pago+Retenções"`
}
