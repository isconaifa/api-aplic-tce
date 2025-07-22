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

	r := mux.NewRouter()
	r.HandleFunc("/aplic/exercicios", excontroller.GetExercicios).Methods("GET")
	r.HandleFunc("/aplic/competencias", competcontroller.GetCompetencias).Methods("GET")
	r.HandleFunc("/aplic/tipos-de-carga", tiposDeCargasController.GetTiposDeCargas).Methods("GET")
	r.HandleFunc("/aplic/conselheiros", conseelController.GetConselheiros).Methods("GET")
	r.HandleFunc("/aplic/municipios", municipiosController.GetMunicipios).Methods("GET")
	r.HandleFunc("/aplic/unidades-gestoras", unidadesGestorasController.GetUnidadesGestoras).Methods("GET")
	return r, nil
}
