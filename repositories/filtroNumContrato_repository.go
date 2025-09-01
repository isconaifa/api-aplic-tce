package repositories

import (
	"api-aplic-web/models"
	"api-aplic-web/queries"
	"database/sql"
	"log"
)

type FiltroNumContratoRepository struct {
	db *sql.DB
}

func NewFiltroNumContratoRepository(db *sql.DB) *FiltroNumContratoRepository {
	return &FiltroNumContratoRepository{db: db}
}

func (repository *FiltroNumContratoRepository) GetFiltroNumContrato(unidadeGestoraCodigo, ano, numContrato string) ([]models.FiltroNumContrato, error) {

	rows, err := repository.db.Query(queries.NumContratoQuery, unidadeGestoraCodigo, ano, numContrato)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var filtroNumcontratos []models.FiltroNumContrato

	for rows.Next() {
		var e models.FiltroNumContrato // Variável para armazenar os dados da linha atual

		// O Scan preenche os campos da struct 'e' com os valores da linha.
		// A ordem aqui DEVE ser idêntica à ordem das colunas no seu SELECT.
		if err := rows.Scan(
			// Colunas da query em ordem:
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
			// Trate o erro adequadamente (ex: logar e continuar, ou retornar o erro)
			log.Printf("Erro ao fazer o scan da linha: %v", err)
			continue // ou return err
		}

		// Adiciona o item populado à slice de resultados
		filtroNumcontratos = append(filtroNumcontratos, e)
	}

	// Não se esqueça de verificar erros que podem ter ocorrido durante a iteração
	if err := rows.Err(); err != nil {
		// Trate o erro do loop
		log.Fatalf("Erro após o loop de rows: %v", err)
	}
	return filtroNumcontratos, nil
}
