package config

import (
	pg "github.com/go-pg/pg/v10"
)

const (
	DB_ADDRESS = ":5432"
	DB_HOST = "localhost"
	DB_USER = "muhammadsarvar"
	DB_PASSWORD = "onam2575"
	DB_PORT = 5432
	DB_NAME = "folders"
)

func DBConnection () (*pg.DB) {

	db := pg.Connect(&pg.Options{
			Addr:	DB_ADDRESS,
			User:	DB_USER,
			Password:	DB_PASSWORD,
			Database: DB_NAME,
	})

	return db
}