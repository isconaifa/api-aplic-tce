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
	adiantamentosRepository := repositories.NewAdiantamentoRepository(db)
	adcontroller := controllers.NewAdiantamentoController(adiantamentosRepository)

	r := mux.NewRouter()
	r.HandleFunc("/adiantamentos", adcontroller.GetAdiantamentos).Methods("GET")
	return r, nil
}
