package db

import (
	"database/sql"
	"fmt"
	"testApi/config"
	"testApi/helpers"
)

func SetupDB(cfg config.StorageConfig) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Username, cfg.Password, cfg.Database)
	db, err := sql.Open("postgres", dbinfo)

	helpers.CheckErr(err)

	return db
}
