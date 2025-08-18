package routes

import (
	"api-aplic-web/controllers"
	"api-aplic-web/database"
	"api-aplic-web/repositories"
	"github.com/gorilla/mux"
)

func SetupRoutes() (*mux.Router, error) {

	db, err := database.Connectdb()
	if err != nil {
		return nil, err
	}
	// request de exercicios
	exerciciosRepository := repositories.NewExercicioRepository(db)
	excontroller := controllers.NewExercicioController(exerciciosRepository)

	// request de competencias
	competRepository := repositories.NewCompetenciaRepository(db)
	competcontroller := controllers.NewCompetenciaController(competRepository)

	// request de tipos de carga
	tiposDeCargasRepository := repositories.NewTipoDeCargaRepository(db)
	tiposDeCargasController := controllers.NewTipoDeCargaController(tiposDeCargasRepository)

	// request de conselheiros
	conselheirosRepository := repositories.NewConselheiroRepository(db)
	conseelController := controllers.NewConselheiroController(conselheirosRepository)

	// request de municipios
	municipiosRepository := repositories.NewMunicipioRepository(db)
	municipiosController := controllers.NewMunicipioController(municipiosRepository)

	// request de unidades gestoras
	unidadesGestorasRepository := repositories.NewUnidadeGestoraRepository(db)
	unidadesGestorasController := controllers.NewUnidadeGestoraController(unidadesGestorasRepository)

	// request de  consultas de empenhos
	consultaEmpenhoRepository := repositories.NewConsultaEmpenhoRepository(db)
	consultaEmpenhoController := controllers.NewConsultaEmpenhoController(consultaEmpenhoRepository)

	// request de cadastro geral
	cadastroGeralRepository := repositories.NewCadastroGeralRepository(db)
	cadGeralController := controllers.NewCadastroGeralController(cadastroGeralRepository)

	// request de orgaos
	orgaoRepository := repositories.NewOrgaoRepository(db)
	orgaoController := controllers.NewOrgaoController(orgaoRepository)

	// request de funcoes
	funcoesRepository := repositories.NewFuncaoRepository(db)
	funcoesController := controllers.NewFuncaoController(funcoesRepository)

	// request de fontes de destinacao de recursos
	fonteDestinacaoRecursoRepository := repositories.NewFonteDestinacaoRecursoRepository(db)
	fonteDestinacaoRecursoController := controllers.NewFonteDestinacaoRecursoController(fonteDestinacaoRecursoRepository)

	// request de açoes
	acaoRepository := repositories.NewAcaoRepository(db)
	acaoController := controllers.NewAcaoController(acaoRepository)

	// request de subfuncoes
	subFuncaoRepository := repositories.NewSubFuncaoRepository(db)
	subFuncaoController := controllers.NewSubFuncaoController(subFuncaoRepository)

	// request de detalhe fonte
	detalheFonteRepository := repositories.NewDetalhefonteRepository(db)
	detalheFonteController := controllers.NewDetalheFonteController(detalheFonteRepository)

	// request de grupos de fonte
	grupoFonteRepository := repositories.NewGrupoFonteRepository(db)
	grupoFonteController := controllers.NewGrupoFonteController(grupoFonteRepository)

	// request de unidades orcamentarias
	unidadeOrcamentariaRepository := repositories.NewUnidadeOrcamentariaRepository(db)
	unidadeOrcamentariaController := controllers.NewUnidadeOrcamentariaController(unidadeOrcamentariaRepository)

	// request de programa
	programaRepository := repositories.NewProgramaRepository(db)
	programaController := controllers.NewProgramaController(programaRepository)

	//request de filtroOrgao
	filtroOrgaoRepository := repositories.NewFiltroOrgaoRepository(db)
	filtroOrgaoController := controllers.NewFiltroOrgaoController(filtroOrgaoRepository)

	// request de filtroFuncao
	filtroFuncaoRepository := repositories.NewFiltroFuncaoRepository(db)
	filtroFuncaoController := controllers.NewFiltroFuncaoController(filtroFuncaoRepository)

	// request de  filtrosubFuncao
	filtroSubFuncaoRepository := repositories.NewFiltroSubFuncaoRepository(db)
	filtroSubFuncaoController := controllers.NewFiltroSubFuncaoController(filtroSubFuncaoRepository)

	// request de diltrodestinacaorecursos
	filtroDestinacaoRecursoRepository := repositories.NewFiltroDestinacaoRecursoRepository(db)
	filtroDestinacaoRecursoController := controllers.NewFiltroDestinacaoRecursoController(filtroDestinacaoRecursoRepository)

	// request de  filtroGrupoFonte
	filtroGrupoFonteRepository := repositories.NewFiltroGrupoFonteRepository(db)
	filtroGrupoFonteController := controllers.NewFiltroGrupoFonteController(filtroGrupoFonteRepository)

	// request de  filtroPrograma
	filtroProgramaRepository := repositories.NewFiltroProgramaRepository(db)
	filtroProgramaController := controllers.NewFiltroProgramaController(filtroProgramaRepository)

	// request de  filtroAcao
	filtroAcaoRepository := repositories.NewFiltroAcaoRepository(db)
	filtroAcaoController := controllers.NewFiltroAcaoController(filtroAcaoRepository)

	// request de  filtroDetalheFonte
	filtroDetalheFonteRepository := repositories.NewFiiltroDetalheFonteRepository(db)
	filtroDetalheFonteController := controllers.NewFiltroDetalheFonteController(filtroDetalheFonteRepository)

	// request de  filtroUnidadeOrcamentaria
	filtroUnidadeOrcamentariaRepository := repositories.NewFiltroUnidadeOrcamentariaRepository(db)
	filtroUnidadeOrcamentariaController := controllers.NewFiltroUnidadeOrcamentariaController(filtroUnidadeOrcamentariaRepository)

	r := mux.NewRouter()
	r.HandleFunc("/aplic/exercicios", excontroller.GetExercicios).Methods("GET")
	r.HandleFunc("/aplic/competencias", competcontroller.GetCompetencias).Methods("GET")
	r.HandleFunc("/aplic/tipos-de-carga", tiposDeCargasController.GetTiposDeCargas).Methods("GET")
	r.HandleFunc("/aplic/conselheiros", conseelController.GetConselheiros).Methods("GET")
	r.HandleFunc("/aplic/municipios", municipiosController.GetMunicipios).Methods("GET")
	r.HandleFunc("/aplic/unidades-gestoras", unidadesGestorasController.GetUnidadesGestoras).Methods("GET")
	r.HandleFunc("/aplic/consultas-empenhos", consultaEmpenhoController.GetAllConsultaEmpenhos).Methods("GET")
	r.HandleFunc("/aplic/cadastro-geral", cadGeralController.GetAllCadastroGeral).Methods("GET")
	r.HandleFunc("/aplic/orgaos", orgaoController.GetAllOrgaos).Methods("GET")
	r.HandleFunc("/aplic/funcoes", funcoesController.GetFuncoes).Methods("GET")
	r.HandleFunc("/aplic/fontes-destinacao-recurso", fonteDestinacaoRecursoController.GetFontesDestinacaoRecurso).Methods("GET")
	r.HandleFunc("/aplic/acoes", acaoController.GetAllAcoes).Methods("GET")
	r.HandleFunc("/aplic/subfuncoes", subFuncaoController.GetAllSubFuncoes).Methods("GET")
	r.HandleFunc("/aplic/detalhe-fonte", detalheFonteController.GetAllDetalheFonte).Methods("GET")
	r.HandleFunc("/aplic/grupo-fonte", grupoFonteController.GetAllGruposFonte).Methods("GET")
	r.HandleFunc("/aplic/unidades-orcamentarias", unidadeOrcamentariaController.GetAllUnidadeOrcamentaria).Methods("GET")
	r.HandleFunc("/aplic/programas", programaController.GetAllProgramas).Methods("GET")
	r.HandleFunc("/aplic/filtroOrgao", filtroOrgaoController.GetAllFiltroOrgao).Methods("GET")
	r.HandleFunc("/aplic/filtroFuncao", filtroFuncaoController.GetAllFiltroFuncao).Methods("GET")
	r.HandleFunc("/aplic/filtroSubFuncao", filtroSubFuncaoController.GetAllFiltroSubFuncao).Methods("GET")
	r.HandleFunc("/aplic/filtroDestinacaoRecurso", filtroDestinacaoRecursoController.GetAllFiltroDestinacaoRecurso).Methods("GET")
	r.HandleFunc("/aplic/filtroGrupoFonte", filtroGrupoFonteController.GetAllFiltroGrupoFonte).Methods("GET")
	r.HandleFunc("/aplic/filtroPrograma", filtroProgramaController.GetAllProgramas).Methods("GET")
	r.HandleFunc("/aplic/filtroAcao", filtroAcaoController.GetAllFiltroAcao).Methods("GET")
	r.HandleFunc("/aplic/filtroDetalheFonte", filtroDetalheFonteController.GetAllFiiltroDetalheFonte).Methods("GET")
	r.HandleFunc("/aplic/filtroUnidadeOrcamentaria", filtroUnidadeOrcamentariaController.GetAllFiltroUnidadeOrcamentaria).Methods("GET")
	return r, nil
}
