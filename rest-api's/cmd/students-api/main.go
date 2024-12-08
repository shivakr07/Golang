package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shivakr07/students-api/internal/config"
	"github.com/shivakr07/students-api/internal/http/handlers/student"
)

func main() {

	//1: load config
	//: load logger if you have custom[here builtin]
	//2: db setup
	//3: setup router
	//4: setup server

	//1:load config
	cfg := config.MustLoad()

	//3: setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	//4: setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal("failed to start server", err)
	// }

	//we will run listen and serve in a go routine
	//buffered [size 1]
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//as out main function get finish we will defer it
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
	//so till now whatever we did was it will complete ongoing req but will not take incomning req
}
