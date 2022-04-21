package db

import (
	"database/sql"
	"fmt"
	"log"
)

var instance *sql.DB

func New(dsn string) (*sql.DB, error) {
	if instance == nil {
		instance = new(sql.DB)

		if dsn == "" {
			return instance, fmt.Errorf("dsn can not be missing")
		}

		inst, err := sql.Open("pgx", dsn)
		if err != nil {
			return instance, err
		}

		instance = inst

		log.Println("Connect to database")

		return instance, nil
	}

	return instance, nil
}