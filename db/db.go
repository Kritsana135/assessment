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

	return db
}

func GetDB() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}
