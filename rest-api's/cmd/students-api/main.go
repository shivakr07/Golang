package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shivakr07/students-api/internal/config"
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

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Welcome to students api"))
		w.Write([]byte("Welcome to students api"))
	})

	//4: setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server", err)
	}

}
