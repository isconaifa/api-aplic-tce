package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
)

type ConsultaEmpenhoRepository struct {
	db *sql.DB
}

func NewConsultaEmpenhoRepository(db *sql.DB) *ConsultaEmpenhoRepository {
	return &ConsultaEmpenhoRepository{db: db}
}

func (repository *ConsultaEmpenhoRepository) GetAllConsultaEmpenhos(unidadeGestoraCodigo string, ano string) ([]models.Empenho, error) {
	rows, err := repository.db.Query(queries.EmpenhoTotal,
		sql.Named("unidadeGestoraCodigo", unidadeGestoraCodigo),
		sql.Named("ano", ano))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var empenhos []models.Empenho

	for rows.Next() {
		var empenho models.Empenho
		err := rows.Scan(
			&empenho.EntCodigo,
			&empenho.OrgaoCodigo,
			&empenho.UnidadeCodigo,
			&empenho.NumeroEmpenho,
			&empenho.FuncaoCodigo,
			&empenho.FuncaoDescricao,
			&empenho.SubfuncaoCodigo,
			&empenho.SubfuncaoDescricao,
			&empenho.ProgramaCodigo,
			&empenho.ProjetoAtividadeNum,
			&empenho.CategoriaEconomica,
			&empenho.NaturezaDespesa,
			&empenho.ModalidadeAplicacao,
			&empenho.ElementoDespesaCod,
			&empenho.ElementoDespesaDesc,
			&empenho.SubElementoDespCod,
			&empenho.SubElementoDespDesc,
			&empenho.Data,
			&empenho.Descricao,
			&empenho.ProcessoLicitatorioNum,
			&empenho.ModalidadeProcLicitCod,
			&empenho.ModalidadeProcLicitDesc,
			&empenho.ContratoNum,
			&empenho.ContratoTipo,
			&empenho.ContratoNumAditivo,
			&empenho.ConvenioNum,
			&empenho.ConvenioNumAditivo,
			&empenho.CompraDireta,
			&empenho.Tipo,
			&empenho.ValorEmpenhadoBruto,
			&empenho.MesReferencia,
			&empenho.CredorIdentificacao,
			&empenho.CredorNome,
			&empenho.CredorTipoPessoaCod,
			&empenho.CredorTipoPessoa,
			&empenho.OrgaoNome,
			&empenho.UnidadeNome,
			&empenho.FundCompraDiretaCod,
			&empenho.FundCompraDiretaDesc,
			&empenho.OptanteSimples,
			&empenho.QtdeNotasFiscais,
			&empenho.QtdeNFe,
			&empenho.QtdeContratos,
			&empenho.DestRecIdusoCod,
			&empenho.DestRecGrupoCod,
			&empenho.DestRecEspecificCod,
			&empenho.DestRecCodigo,
			&empenho.DestRecIdusoDesc,
			&empenho.DestRecGrupoDesc,
			&empenho.DestRecEspecificDesc,
			&empenho.DestRecDesc,
			&empenho.TipoServicoEngenharia,
			&empenho.NumeroObra,
			&empenho.ProjetoAtividadeDesc,
			&empenho.Dotacao,
			&empenho.CodigoUG,
			&empenho.Exercicio,
			&empenho.Municipio,
			&empenho.UnidadeGestora,
			&empenho.Relevante,
			&empenho.QtdeDiarias,
			&empenho.FonteRecursoDesc,
			&empenho.AssistenciaSocial,
			&empenho.ConcursoNumero,
			&empenho.ConcursoTipo,
			&empenho.FonteRecursoCodigo,
			&empenho.InstrumentoContrato,
			&empenho.TipoDespesaRPPS,
			&empenho.QtdeBeneficiarios,
			&empenho.ValorEmpenhado,
			&empenho.AnuladoEmpenho,
			&empenho.AnuladoLiquidacao,
			&empenho.AnuladoPagamento,
			&empenho.ValorRetido,
			&empenho.ValorLiquidado,
			&empenho.ValorLiquidadoBruto,
			&empenho.ValorPagoBruto,
			&empenho.ValorPago,
			&empenho.ValorPagoRetencoes,
		)
		if err != nil {
			return nil, err
		}
		empenhos = append(empenhos, empenho)
	}
	return empenhos, nil
}
