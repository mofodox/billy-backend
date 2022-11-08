package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app application) setupRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
	})

	return router
}
