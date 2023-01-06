package main

import (
	"net/http"
	"time"

	"github.com/Kritsana135/assessment/config"
	"github.com/Kritsana135/assessment/db"
	"github.com/Kritsana135/assessment/expense/delivery/http_"
	"github.com/Kritsana135/assessment/expense/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	config.LoadConfig("./")

	db.ConnectDB()

	expUCase := usecase.NewExpUsecase()

	r := gin.Default()
	r.GET("/health", GetHealthCheck)

	http_.NewExpenseHandler(&r.RouterGroup, expUCase)

	r.Run(":" + viper.GetString("PORT"))
}

type ResponseHealthCheck struct {
	Message string    `json:"message"`
	Uptime  time.Time `json:"uptime"`
}

func GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseHealthCheck{
		Message: "OK",
		Uptime:  time.Now(),
	})
}
