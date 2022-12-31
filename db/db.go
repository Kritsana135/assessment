package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var db *sqlx.DB

func Connect() *sqlx.DB {
	var err error
	dsn := fmt.Sprintf(
		"user=%v password=%v host=%v dbname=%v sslmode=disable port=%v",
		"postgres", "postgres", "localhost", "postgres", 5432) // TODO: move to .env

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	logrus.Info("Database connected")

	initDb()

	return db
}

func GetDB() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}

func initDb() {
	schema := `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			amount numeric(19,4) NOT NULL,
			note VARCHAR(255),
			tags text[]
		)
	`

	result := db.MustExec(schema)

	_, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	logrus.Info("Database initialized")
}
