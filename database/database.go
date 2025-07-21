package database

import (
	"api-aplic-web/configs"
	"database/sql"
	_ "github.com/sijms/go-ora/v2"
	"log"
)

func Connectdb() (*sql.DB, error) {
	db, err := sql.Open("oracle", configs.StringConexion)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error to connect database")
		return nil, err
	}
	return db, nil
}
