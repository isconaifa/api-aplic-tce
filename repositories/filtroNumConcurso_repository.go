package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
	"log"
)

type FiltroNumConcursoRepository struct {
	db *sql.DB
}

func NewFiltroNumConcursoRepository(db *sql.DB) *FiltroNumConcursoRepository {
	return &FiltroNumConcursoRepository{db: db}
}

func (repository *FiltroNumConcursoRepository) GetFiltroNumConcurso(unidadeGestoraCodigo, ano, numConcurso string) ([]models.FiltroNumConcurso, error) {
	rows, err := repository.db.Query(queries.NumConcurso, unidadeGestoraCodigo, ano, numConcurso)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var filtroNumconcursos []models.FiltroNumConcurso

	for rows.Next() {
		var e models.FiltroNumConcurso
		if err := rows.Scan(
			&e.EntCodigo,
			&e.OrgaoCodigo,
			&e.UnidadeOrcamentariaCodigo,
			&e.EmpNumero,
			&e.FuncaoCodigo,
			&e.FuncaoDescricao,
			&e.SubFuncaoCodigo,
			&e.SubFuncaoDescricao,
			&e.ProgramaCodigo,
			&e.NumProjetoAtividade,
			&e.CategoriaEconomica,
			&e.NaturezaDespesa,
			&e.ModalidadeAplicacaoCodigo,
			&e.ElementoDespesaCodigo,
			&e.ElementoDespesaDescricao,
			&e.SubelementoDespesaCodigo,
			&e.SubelementoDespesa,
			&e.EmpData,
			&e.Descricao,
			&e.NumProcessoLicitatorio,
			&e.ModalidadeProcLicitCodigo,
			&e.ModalidadeProcLicitatorio,
			&e.NumContrato,
			&e.TipoContrato,
			&e.NumAditivoContrato,
			&e.NumConvenio,
			&e.NumAditivoConvenio,
			&e.CompraDireta,
			&e.Tipo,
			&e.EmpenhadoSemAnulacao,
			&e.MesReferencia,
			&e.IdentificacaoCredor,
			&e.Credor,
			&e.TipoPessoaCodigo,
			&e.TipoPessoa,
			&e.Orgao,
			&e.UnidadeOrcamentaria,
			&e.FundCompraDiretaCodigo,
			&e.FundCompraDiretaDescricao,
			&e.OptanteSimples,
			&e.QtdeNotasFiscais,
			&e.QtdeNFe,
			&e.Contratos,
			&e.DestRecCodigoIduso,
			&e.DestRecCodigoGrupo,
			&e.DestRecCodEspecificacao,
			&e.CodDestinacaoRecurso,
			&e.DestRecIduso,
			&e.DestRecGrupo,
			&e.DestRecEspecificacao,
			&e.DestinacaoRecurso,
			&e.TipoServEngenharia,
			&e.NumObra,
			&e.ProjetoAtividade,
			&e.Dotacao,
			&e.CodigoUG,
			&e.Exercicio,
			&e.Municipio,
			&e.UnidadeGestora,
			&e.Relevante,
			&e.Diarias,
			&e.FonteRecursoDescricao,
			&e.AssistenciaSocial,
			&e.NumConcurso,
			&e.TipoConcurso,
			&e.FonteRecursoCodigo,
			&e.InstrumentoContrato,
			&e.TipoDespesaRPPS,
			&e.QtdeBeneficiarios,
			&e.ValorEmpenhado,
			&e.AnuladoEmpenho,
			&e.AnuladoLiquidacao,
			&e.AnuladoPagamento,
			&e.ValorRetidoLiquidacao,
			&e.ValorLiquidado,
			&e.ValorLiquidadoSemAnulacao,
			&e.ValorPagoSemAnulacao,
			&e.ValorPago,
			&e.ValorPagoMaisRetencoes,
		); err != nil {
			log.Fatal(err)
		}
		filtroNumconcursos = append(filtroNumconcursos, e)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return filtroNumconcursos, nil
}
