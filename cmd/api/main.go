package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const VERSION = "0.0.1"

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "env (dev | staging | prod)")
	flag.Parse()

	app := application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.setupRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("%s server is starting on port %s\n", cfg.env, srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
