package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/swarajbhagat11/todo-app/app/handler"
	"github.com/swarajbhagat11/todo-app/app/repository"
	"github.com/swarajbhagat11/todo-app/app/service"
	"github.com/swarajbhagat11/todo-app/server"
)

func main() {
	// Configuration phase
	log.SetFormatter(&log.TextFormatter{})

	if err := initConfigs(); err != nil {
		log.Fatal("[Main] error loading config files:", err)
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	hnd := handler.NewHandler(services)

	services.Todo.ReadData()

	// Routing phase
	srv := new(server.Server)
	port := viper.GetString("port")

	go func() {
		if err := srv.Run(port, hnd.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalln("[Main] error running server", err)
		}
	}()

	log.Println("[Main] server running on port", port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()

	log.Println("[Main] server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorln("[Main] error shutting down server:", err)
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
