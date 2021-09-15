package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	apiServer "github.com/torkolenko/go-notes"
	handler "github.com/torkolenko/go-notes/pkg/handler"
	repository "github.com/torkolenko/go-notes/pkg/repository"
	service "github.com/torkolenko/go-notes/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("failed loading env variables: ", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("failed to initialize db: ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	srv := new(apiServer.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http-server: %s", err.Error())
		}
	}()

	logrus.Print("NotesApp started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("NotesApp shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
