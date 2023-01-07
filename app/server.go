package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kritsana135/assessment/config"
	"github.com/Kritsana135/assessment/db"
	"github.com/Kritsana135/assessment/expense/delivery/http_"
	"github.com/Kritsana135/assessment/expense/repository/postgresql"
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

	connection := db.ConnectDB()

	expRepo := postgresql.NewExpenseRepo(connection)
	expUCase := usecase.NewExpUsecase(expRepo)

	r := gin.Default()
	r.GET("/health", GetHealthCheck)

	http_.NewExpenseHandler(&r.RouterGroup, expUCase)

	port := viper.GetString("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		logrus.Info("Server started on port " + port)

		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("Shutdown Server ...")

	timeOut := viper.GetUint("GF_SHUTDOWN_TIMEOUT") * uint(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut))
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	db.CloseDatabase()
	logrus.Info(fmt.Sprintf("timeout of %v seconds.", viper.GetUint("GF_SHUTDOWN_TIMEOUT")))
	logrus.Info("Server exiting")
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
