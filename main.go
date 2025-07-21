package main

import (
	"api-aplic-web/configs"
	"api-aplic-web/middlewares"
	"api-aplic-web/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configs.Carregar()
	//fmt.Println(database.Connectdb())
	route, err := routes.SetupRoutes()
	if err != nil {
		panic(err)
	}
	route.Use(middlewares.LoggingMiddleware)
	fmt.Printf("Server is running on port %d\n", configs.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Porta), route))

}
