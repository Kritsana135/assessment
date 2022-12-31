package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kritsana135/assessment/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
